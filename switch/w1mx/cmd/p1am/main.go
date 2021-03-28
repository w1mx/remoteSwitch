package main

import (
	"fmt"
	"io"
	"machine"
	"runtime"
	"time"

	"github.com/simpleiot/simpleiot/modbus"
	"tinygo.org/x/drivers/p1am"
)

func main() {
	for i := 10; i > 0; i-- {
		fmt.Printf("Waiting %d\n", i)
		time.Sleep(time.Second)
	}
	for {
		if err := loop(); err != nil {
			fmt.Printf("loop failed, retrying: %v\n", err)
			time.Sleep(500 * time.Millisecond)
		}
	}
}

type P1amRegs struct {
	led machine.Pin
	sw  machine.Pin
}

func (p *P1amRegs) ReadReg(address int) (uint16, error) {
	// No holding registers
	return 0, modbus.ExcIllegalFunction
}

func (p *P1amRegs) WriteReg(address int, value uint16) error {
	if address < 0 || address >= p1am.Controller.Slots {
		return modbus.ExcIllegalAddress
	}
	slot := p1am.Controller.Slot(address + 1)
	if slot.Props.DO > 0 {
		return slot.WriteDiscrete(uint32(value))
	}
	return modbus.ExcIllegalAddress
}

func (p *P1amRegs) ReadInputReg(address int) (uint16, error) {
	if address >= p1am.Controller.Slots {
		return 0, modbus.ExcIllegalAddress
	}
	slot := p1am.Controller.Slot(address + 1)
	if slot.Props.DI > 0 {
		value, err := slot.ReadDiscrete()
		return uint16(value), err
	}
	return 0, modbus.ExcIllegalAddress
}

func (p *P1amRegs) ReadDiscreteInput(num int) (bool, error) {
	s := num / 16
	switch {
	case num == 0:
		return p.sw.Get(), nil
	case s >= 1 && s <= p1am.Controller.Slots:
		slot := p1am.Controller.Slot(s)
		return slot.Channel(1 + (num % 16)).ReadDiscrete()
	}
	return false, modbus.ExcIllegalAddress
}

func (p *P1amRegs) ReadCoil(num int) (bool, error) {
	// TODO: Support ReadCoil
	return false, modbus.ExcIllegalFunction
}
func (p *P1amRegs) WriteCoil(num int, value bool) error {
	s := num / 16
	switch {
	case num == 0:
		p.led.Set(value)
		return nil
	case s >= 1 && s <= p1am.Controller.Slots:
		slot := p1am.Controller.Slot(s)
		return slot.Channel(1 + (num % 16)).WriteDiscrete(value)
	}
	return modbus.ExcIllegalAddress
}

const interCharDelay = 750 * time.Microsecond

func loop() error {
	led := machine.LED
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})
	sw := machine.SWITCH
	sw.Configure(machine.PinConfig{Mode: machine.PinInput})

	if err := p1am.Controller.Initialize(); err != nil {
		return fmt.Errorf("initializing controller: %w", err)
	}

	version, err := p1am.Controller.Version()
	if err != nil {
		return fmt.Errorf("fetching base controller version: %w", err)
	}
	fmt.Printf("Base controller version: %d.%d.%d\n", version[0], version[1], version[2])

	if err := p1am.Controller.ConfigureWatchdog(5*time.Second, true); err != nil {
		return err
	}
	if err := p1am.Controller.StartWatchdog(); err != nil {
		return err
	}

	regs := &P1amRegs{led: led, sw: sw}

	var buf [255]byte
	offset := 0

	var lastTime time.Time
	for {
		cnt, err := machine.UART0.Read(buf[offset:])
		// TODO: Read until a 3.5 char gap has occurred
		// Spec says inter-char timeout should be 750us, end of frame timeout should be 1.75ms
		if err != nil && err != io.EOF {
			return err
		}
		if cnt > 0 {
			lastTime = time.Now()
		}
		offset += cnt
		if offset == 0 || time.Since(lastTime) < interCharDelay {
			if err := p1am.Controller.PetWatchdog(); err != nil {
				return err
			}
			runtime.Gosched()
			continue
		}

		packet := buf[:offset]
		offset = 0

		req, err := modbus.RtuDecode(packet)
		if err != nil {
			// Standard says to ignore bad packets
			fmt.Println("failed decode")
			continue
		}
		_, resp, err := req.ProcessRequest(regs)
		if err != nil {
			fmt.Println("failed process")
			continue
		}
		respRtu, err := modbus.RtuEncode(packet[0], resp)
		if err != nil {
			continue
		}
		machine.UART0.Write(respRtu)
	}
	return nil
}
