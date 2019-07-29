package outputs

import (
	tpdevices "github.com/kastenpotential/hello-go/tplink/devices"
	"errors"
)


// Output asdfasdf
type Output interface {
	Write(tpdevices.TPDevices) error
}

type bad struct {}
func (b *bad) Write(devices tpdevices.TPDevices) error {
	return nil
}

// Outputer asdfasdf
type Outputer func() (error, Output)

// Outputs adsfasdf
var Outputs = map[string]Outputer{}

// AddOutput adsfasdf
func AddOutput(name string, output Outputer) {
	Outputs[name] = output
}

// GetOutput asdfasdf
func GetOutput(name string) (error, Output) {
	output, ok := Outputs[name]
	if !ok {
		return errors.New("output not found"), &bad{}
	}
	return output()
}