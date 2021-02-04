package w1mx

type SwitchConfig struct {
	Name string
	Index int
	PortName string
	Radios []Radio
	TerminalNames []string
	Controls []Control
}
type Radio struct {
	Name string
	EnableControl string
}
type Control struct {
	Name string
	Channel int
}
