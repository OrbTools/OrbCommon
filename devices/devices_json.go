// +build !xdr

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

//go:embed json/*
var data embed.FS

func init() {
	DeviceTypes = make(map[string]*DeviceDef)
	files, _ := fs.ReadDir(data, "json")
	for _, file := range files {
		dev := new(DeviceDef)
		data, _ := data.ReadFile("json/" + file.Name())
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
