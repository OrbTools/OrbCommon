package devices

import (
	"bytes"
	"embed"
	"github.com/Minizbot2012/minxdr"
	"github.com/OrbTools/OrbCommon/devices/structs"
	"io"
	"io/fs"
	"strings"
)

var DeviceTypes map[string]*structs.DeviceDef

//go:embed xdr/*
var df embed.FS

func init() {
	DeviceTypes = make(map[string]*structs.DeviceDef)
	files, _ := fs.ReadDir(df, "xdr")
	for _, file := range files {
		dev := new(structs.DeviceDef)
		data, _ := df.ReadFile("xdr/" + file.Name())
		_, err := minxdr.Unmarshal(bytes.NewReader(data), dev)
		if err != nil {
			panic(err.Error())
		}
		DeviceTypes[strings.Split(file.Name(), ".")[0]] = dev
	}
}

//LoadKeymap Load Orbmap KM structure
func LoadKeymap(file io.ReadCloser) *structs.KeyMap {
	mapped := new(structs.KeyMap)
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
