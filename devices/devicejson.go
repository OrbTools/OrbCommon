package devices

import (
	"github.com/OrbTools/OrbCommon/gui"
)

//ExtraBytes defines extra bytes to a type
type ExtraBytes struct {
	Name string
	Size int
}

//Device defines a JSON device
type Device struct {
	BINDING []byte
	EB      []*ExtraBytes
	GUI     gui.GUI
}

//DeviceList List of supported devices
var DeviceList = []string{
	"Orbweaver",
}
