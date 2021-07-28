// +build xdr

package devices

import (
	"bytes"
	"embed"
	xdr "github.com/davecgh/go-xdr/xdr2"
	"io"
	"io/fs"
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
func LoadKeymap(file io.ReadCloser, dev *DeviceDef) *KeyMap {
	mapped := new(KeyMap)
	xdr.Unmarshal(file, mapped)
	file.Close()
	return mapped
}

//SaveKeymap Save Orbmap KM struction
func SaveKeymap(file io.WriteCloser, mapped interface{}) {
	xdr.Marshal(file, mapped)
	file.Close()
}
