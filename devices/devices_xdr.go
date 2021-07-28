// +build xdr

package devices

import (
	"bytes"
	"embed"
	xdr "github.com/davecgh/go-xdr/xdr2"
	"io"
	"io/fs"
	"os"
	"strings"
)

//go:embed xdr/*
var df embed.FS

func init() {
	DeviceTypes = make(map[string]*DeviceDef)
	files, _ := fs.ReadDir(df, "xdr")
	for _, file := range files {
		dev := new(DeviceDef)
		data, _ := df.ReadFile("xdr/" + file.Name())
		xdr.Unmarshal(bytes.NewReader(data), dev)
		DeviceTypes[strings.Split(file.Name(), ".")[0]] = dev
	}
}

//LoadKeymap Load Orbmap KM structure
func LoadKeymap(file string, dev *DeviceDef) *KeyMap {
	mapped := new(KeyMap)
	of, _ := os.Open(file)
	xdr.Unmarshal(of, mapped)
	return mapped
}

//SavePKMKeymap saves an orb after edit
func SaveKeymap(mapped interface{}, file io.WriteCloser) {
	xdr.Marshal(file, mapped)
	file.Close()
}
