package configparser

import (
	"fmt"

	"github.com/dh1tw/remoteSwitch/switch/w1mx"
	"github.com/spf13/viper"
)

// GetW1MXSwitchConfig tries to parse the config file via viper
// and returns on success a dummySwitch.SwitchConfig object.
func GetW1MXSwitchConfig(switchName string) (w1mx.SwitchConfig, error) {
	sc := w1mx.SwitchConfig{}

	n := func(p string) string { return fmt.Sprintf("%s.%s", switchName, p) }
	// let's check first if all necessary keys exist in the config file
	for _, p := range []string{
		"name",
		"index",
		"portname",
		"radios",
		"terminals",
		"controls",
	} {
		if !viper.IsSet(n(p)) {
			return sc, fmt.Errorf("missing %s parameter for switch %s", p, switchName)
		}
	}

	// get the values
	name := viper.GetString(n("name"))
	if len(name) == 0 {
		return sc, fmt.Errorf("name parameter of switch %s must not be empty", switchName)
	}

	index := viper.GetInt(n("index"))
	portname := viper.GetString(n("portname"))
	radios := viper.GetStringSlice(n("radios"))
	if len(radios) == 0 {
		return sc, fmt.Errorf("no radios found for switch %s", switchName)
	}
	terminals := viper.GetStringSlice(n("terminals"))
	if len(terminals) == 0 {
		return sc, fmt.Errorf("no terminals found for switch %s", switchName)
	}
	controls := viper.GetStringSlice(n("controls"))

	sc.Name = name
	sc.Index = index
	sc.PortName = portname
	sc.TerminalNames = terminals

	for _, control := range controls {
		c, err := getW1MXControl(control)
		if err != nil {
			return sc, err
		}
		sc.Controls = append(sc.Controls, c)
	}

	for _, radio := range radios {
		p, err := getW1MXRadio(radio)
		if err != nil {
			return sc, err
		}
		sc.Radios = append(sc.Radios, p)
	}

	return sc, nil
}

func getW1MXControl(controlName string) (w1mx.Control, error) {
	rc := w1mx.Control{}

	// let's check first if all necessary keys exist in the config file
	if !viper.IsSet(controlName) {
		return rc, fmt.Errorf("no configuration found for control %s", controlName)
	}

	n := func(p string) string { return fmt.Sprintf("%s.%s", controlName, p) }

	if !viper.IsSet(n("name")) {
		return rc, fmt.Errorf("missing name parameter for port %s", controlName)
	}
	if !viper.IsSet(n("channel")) {
		return rc, fmt.Errorf("missing channel parameter for port %s")
	}

	// get the values
	name := viper.GetString(n("name"))
	if len(name) == 0 {
		return rc, fmt.Errorf("name parameter of control %s must not be empty", controlName)
	}

	rc.Name = name
	rc.Channel = viper.GetInt(n("channel"))
	return rc, nil
}

func getW1MXRadio(radioName string) (w1mx.Radio, error) {
	rc := w1mx.Radio{}

	// let's check first if all necessary keys exist in the config file
	if !viper.IsSet(radioName) {
		return rc, fmt.Errorf("no configuration found for radio %s", radioName)
	}

	n := func(p string) string { return fmt.Sprintf("%s.%s", radioName, p) }

	if !viper.IsSet(n("name")) {
		return rc, fmt.Errorf("missing name parameter for port %s", radioName)
	}

	// get the values
	name := viper.GetString(n("name"))
	if len(name) == 0 {
		return rc, fmt.Errorf("name parameter of radio %s must not be empty", radioName)
	}

	rc.Name = name
	if viper.IsSet(n("enable_control")) {
		rc.EnableControl = viper.GetString(n("enable_control"))
	}
	return rc, nil
}
