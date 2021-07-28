package devices

//KeyMap singular keymap
type KeyMap struct {
	Keymap []uint16
	Color  [3]byte
}

//KeyMaps a set of keymaps
type KeyMaps struct {
	Maps       []*KeyMap
	Currentmap int
	MCount     int
}

type DeviceDef struct {
	Backend     string
	IsColor     bool
	MaxMappings int
	NumKeys     int
	Binding     []byte
	Device      struct {
		SystemFile string
		VendorID   int
		ProdID     int
	}
	GuiPages []struct {
		Name string
		Type string
		Keys []struct {
			KeyID   int
			KeyName string
		}
	}
}

var DeviceTypes map[string]*DeviceDef
