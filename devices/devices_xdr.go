// +build xdr

package devices

import (
	"embed"
	"github.com/davecgh/go-xdr/xdr2"
	"io"
	"io/fs"
	"os"
	"strings"
)

//go:embed xdr/*
var data embed.FS

func init() {
	DeviceTypes = make(map[string]*DeviceDef)
	files, _ := fs.ReadDir(data, "xdr")
	for _, file := range files {
		dev := new(DeviceDef)
		data, _ := jsons.ReadFile("xdr/" + file.Name())
		xdr2.Unmarshal(data, dev)
		DeviceTypes[strings.Split(file.Name(), ".")[0]] = dev
	}
}

//LoadKM Load Orbmap KM structure
func LoadKeymap(file string, dev *DeviceDef) *KeyMap {
	mapped := new(KeyMap)
	of, _ := os.Open(file)
	xdr.Unmarshal(of, KeyMap)
	return mapped
}

//SavePKMKeymap saves an orb after edit
func SaveKeymap(mapped interface{}, file io.WriteCloser) {
	xdr.Marshal(file, mapped)
	file.Close()
}
