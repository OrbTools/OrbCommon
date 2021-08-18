//go:build !xdr
// +build !xdr

package devices

import (
	"embed"
	"encoding/binary"
	"encoding/json"
	"io"
	"io/fs"
	"strings"
)

//go:embed json/*
var df embed.FS

func init() {
	DeviceTypes = make(map[string]*DeviceDef)
	files, _ := fs.ReadDir(df, "json")
	for _, file := range files {
		dev := new(DeviceDef)
		data, _ := df.ReadFile("json/" + file.Name())
		json.Unmarshal(data, dev)
		DeviceTypes[strings.Split(file.Name(), ".")[0]] = dev
	}
}

//LoadKM Load Orbmap KM structure
func LoadKeymap(file io.ReadCloser, dev *DeviceDef) *KeyMap {
	mapped := new(KeyMap)
	mapped.Keymap = make([]uint16, dev.NumKeys)
	binary.Read(file, binary.LittleEndian, mapped.Keymap)
	binary.Read(file, binary.LittleEndian, mapped.Color)
	file.Close()
	return mapped
}

//SaveKeymap Saves Orbmap KM structure
func SaveKeymap(file io.WriteCloser, mapped interface{}) {
	binary.Write(file, binary.LittleEndian, mapped)
	file.Close()
}
