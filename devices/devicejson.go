package devices

import (
	"encoding/binary"
	"io"

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

//SavePKMKeymap saves an orb after edit
func SavePKMKeymap(mapped interface{}, file io.WriteCloser) {
	binary.Write(file, binary.LittleEndian, mapped)
	file.Close()
}
