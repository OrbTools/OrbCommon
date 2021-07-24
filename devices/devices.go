package devices

import (
	"embed"
	"encoding/binary"
	"encoding/json"
	"io"
	"io/fs"
	"os"
	"strings"
)

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

//go:embed json/*
var jsons embed.FS

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
		Hive string
		Name string
		Type string
		Keys []struct {
			KeyID   int
			KeyName string
		}
	}
}

var DeviceTypes map[string]*DeviceDef

func init() {
	DeviceTypes = make(map[string]*DeviceDef)
	files, _ := fs.ReadDir(jsons, "json")
	for _, file := range files {
		dev := new(DeviceDef)
		data, _ := jsons.ReadFile("json/" + file.Name())
		json.Unmarshal(data, dev)
		DeviceTypes[strings.Split(file.Name(), ".")[0]] = dev
	}
}

//LoadKM Load Orbmap KM structure
func LoadKeymap(file string, dev *DeviceDef) *KeyMap {
	mapped := new(KeyMap)
	of, _ := os.Open(file)
	defer of.Close()
	mapped.Keymap = make([]uint16, dev.NumKeys)
	binary.Read(of, binary.LittleEndian, mapped.Keymap)
	binary.Read(of, binary.LittleEndian, mapped.Color)
	return mapped
}

//SavePKMKeymap saves an orb after edit
func SaveKeymap(mapped interface{}, file io.WriteCloser) {
	binary.Write(file, binary.LittleEndian, mapped)
	file.Close()
}
