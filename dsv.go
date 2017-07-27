package dsv

// Config represents a configuration to generate a dsv string
type Config struct {
	Separator string
}

// Tabbed represents a basic structure to generate a dsv string
type Dsv struct {
	Struct        interface{}
	Configuration Config
}

// ToDsv returns a dsv string from the struct
func (t *Dsv) ToDsv() (string, error) {
	return "", nil
}
