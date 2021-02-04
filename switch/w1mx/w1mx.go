package w1mx

import (
	"fmt"
	"log"
	"sync"

	sw "github.com/dh1tw/remoteSwitch/switch"
	"github.com/goburrow/modbus"
)

type Switch struct {
	sync.RWMutex
	name  string
	index int
	// ports maps name to active terminal
	ports map[string]int
	// terminals maps name to terminal index
	terminals map[string]int
	// controls maps name to state
	controls     map[string]bool
	switchConfig SwitchConfig
	eventHandler func(sw.Switcher, sw.Device)
	client       modbus.Client
}

func NewSwitch(sc SwitchConfig, eventHandler func(sw.Switcher, sw.Device)) *Switch {
	s := &Switch{
		name:         "my switch",
		index:        0,
		switchConfig: sc,
		eventHandler: eventHandler,
	}
	return s
}

func (s *Switch) Init() error {
	s.name = s.switchConfig.Name
	s.index = s.switchConfig.Index
	s.client = modbus.RTUClient(s.switchConfig.PortName)
	s.ports = make(map[string]int)
	for _, r := range s.switchConfig.Radios {
		s.ports[r.Name] = len(s.switchConfig.TerminalNames) - 1
	}
	s.terminals = make(map[string]int)
	for i, t := range s.switchConfig.TerminalNames {
		s.terminals[t] = i
	}
	s.controls = make(map[string]bool)
	for _, c := range s.switchConfig.Controls {
		s.controls[c.Name] = false
	}

	return nil
}

// Name returns the Name of this switch
func (s *Switch) Name() string {
	s.RLock()
	defer s.RUnlock()
	return s.name
}

const signalsName = "Signals"

// SetPort sets the Terminals of a particular Port. The portRequest
// can contain n termials.
func (s *Switch) SetPort(portRequest sw.Port) error {
	s.Lock()
	defer s.Unlock()

	// ensure that the requested port exists
	if _, ok := s.ports[portRequest.Name]; !ok && portRequest.Name != signalsName {
		return fmt.Errorf("%s is an invalid port", portRequest.Name)
	}

	if portRequest.Name == signalsName {
		for _, t := range portRequest.Terminals {
			if _, ok := s.controls[t.Name]; !ok {
				return fmt.Errorf("%s is an invalid terminal", t.Name)
			}
		}
		// set state of the terminal
		for _, t := range portRequest.Terminals {
			s.controls[t.Name] = t.State
		}
	} else {
		// ensure that the requested terminal exists
		for _, t := range portRequest.Terminals {
			if _, ok := s.terminals[t.Name]; !ok {
				return fmt.Errorf("%s is an invalid terminal", t.Name)
			}
		}

		for prtName, prtTerminal := range s.ports {
			// only check the remaining ports
			if prtName == portRequest.Name {
				continue
			}

			for _, t := range portRequest.Terminals {
				if s.switchConfig.TerminalNames[prtTerminal] == t.Name {
					return fmt.Errorf("terminal %s in use by port %s",
						t.Name, prtName)
				}
			}
		}

		// set state of the terminal
		for _, t := range portRequest.Terminals {
			index := s.terminals[t.Name]

			if t.State {
				s.ports[portRequest.Name] = index
			} else {
				// TODO: Default to local
				s.ports[portRequest.Name] = -1
			}
		}
	}

	// TODO: Set TX REQ based on whether radio has a terminal

	if s.eventHandler != nil {
		device := s.serialize()
		go s.eventHandler(s, device)
	}

	return s.executeState()
}

func (s *Switch) executeState() error {
	var signals byte
	for _, c := range s.switchConfig.Controls {
		if s.controls[c.Name] {
			signals |= 1 << (c.Channel - 1)
		}
	}
	out := []byte{signals, 0}
	for _, r := range s.switchConfig.Radios {
		if s.ports[r.Name] >= 0 {
			out = append(out, 1<<s.ports[r.Name])
		} else {
			out = append(out, 0)
		}
	}
	log.Printf("executing state %02x", out)
	// Slot 1 is relays, slot 2 is SixPak
	_, err := s.client.WriteMultipleRegisters(0, uint16(len(out)/2), out)
	return err
}

// GetPort returns switch.Port struct containing the current state of
// the requested port.
func (s *Switch) GetPort(portName string) (sw.Port, error) {
	s.RLock()
	defer s.RUnlock()

	if _, ok := s.ports[portName]; ok {
		return s.serializeRadio(portName), nil
	} else if portName == signalsName {
		return s.serializeSignals(), nil
	}
	return sw.Port{}, fmt.Errorf("%s in an invalid port", portName)
}

// Serialize returns a switch.Device struct containing the current
// state and configuration of this Dummy switch.
func (s *Switch) Serialize() sw.Device {
	s.RLock()
	defer s.RUnlock()

	return s.serialize()
}

// serialize returns a switch.Port struct containing the current
// state and configuration of this Port. This method
// is not threadsafe.
func (s *Switch) serializeRadio(n string) sw.Port {
	index := 0
	var r Radio
	for index, r = range s.switchConfig.Radios {
		if r.Name == n {
			break
		}
	}
	swPort := sw.Port{
		Name:      r.Name,
		Index:     index,
		Terminals: []sw.Terminal{},
	}

	for i, tn := range s.switchConfig.TerminalNames {
		t := sw.Terminal{
			Name:  tn,
			Index: i,
			State: s.ports[r.Name] == i,
		}
		swPort.Terminals = append(swPort.Terminals, t)
	}

	return swPort
}

func (s *Switch) serializeSignals() sw.Port {
	swPort := sw.Port{
		Name:      signalsName,
		Index:     100,
		Terminals: []sw.Terminal{},
	}

	for i, c := range s.switchConfig.Controls {
		t := sw.Terminal{
			Name:  c.Name,
			Index: i,
			State: s.controls[c.Name],
		}
		swPort.Terminals = append(swPort.Terminals, t)
	}

	return swPort
}

// serialize returns a switch.Device struct containing the current
// state and configuration of this Dummy switch. This method
// is not threadsafe.
func (s *Switch) serialize() sw.Device {

	dev := sw.Device{
		Name:  s.name,
		Index: s.index,
	}

	// serialize all ports
	for _, r := range s.switchConfig.Radios {
		swPort := s.serializeRadio(r.Name)
		dev.Ports = append(dev.Ports, swPort)
	}

	dev.Ports = append(dev.Ports, s.serializeSignals())

	return dev
}

// Close shuts down the switch.
func (s *Switch) Close() {
	// nothing to do
}
