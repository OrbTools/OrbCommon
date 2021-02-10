package devices

import "github.com/OrbTools/OrbCommon/gui"

//Device defines a JSON device
//NYI
type Device struct {
	EB      int
	BINDLEN int
	BINDING []byte
	GUI     gui.GUI
}

//DeviceList List of supported devices
var DeviceList = []string{
	"Orbweaver",
}
