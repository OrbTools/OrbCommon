package devices

import (
	_ "embed"
	"github.com/Minizbot2012/minxdr"
)

var DeviceTypes map[string]*DeviceDef

//go:embed xdr/*
var df embed.FS

//KeyMap singular keymap
type KeyMap struct {
	Device string
	Keymap []uint16
	Color  []byte
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
	NumColor    int
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

func init() {
	DeviceTypes = make(map[string]*DeviceDef)
	files, _ := fs.ReadDir(df, "xdr")
	for _, file := range files {
		dev := new(DeviceDef)
		data, _ := df.ReadFile("xdr/" + file.Name())
		_, err := mindxdr.Unmarshal(bytes.NewReader(data), dev)
		if err != nil {
			panic(err.Error())
		}
		DeviceTypes[strings.Split(file.Name(), ".")[0]] = dev
	}
}

//LoadKeymap Load Orbmap KM structure
func LoadKeymap(file io.ReadCloser) *KeyMap {
	mapped := new(KeyMap)
	_, err := minxdr.Unmarshal(file, mapped)
	if err != nil {
		panic(err.Error())
	}
	file.Close()
	return mapped
}

//SaveKeymap Save Orbmap KM struction
func SaveKeymap(file io.WriteCloser, mapped interface{}) {
	minxdr.Marshal(file, mapped)
	file.Close()
}
