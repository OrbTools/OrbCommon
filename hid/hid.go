package hid

import (
	"bytes"
	_ "embed"
	"github.com/Minizbot2012/minxdr"
	"github.com/OrbTools/OrbCommon/hid/structs"
)

//go:embed generated.bin
var file []byte

func init() {
	minxdr.Unmarshal(bytes.NewReader(file), &Mappings)
}

var Mappings structs.KeyMaps

func GetMappingFromHID(uv uint16) structs.Key {
	return Mappings.Usb[uv]
}

func GetMappingFromWindows(uv uint16) structs.Key {
	return Mappings.Win[uv]
}

func GetMappingFromLinux(uv uint16) structs.Key {
	return Mappings.Evdev[uv]
}

func GetMappingFromName(name string) structs.Key {
	return Mappings.Code[name]
}

func GetMappingFromX(code uint16) structs.Key {
	return Mappings.Xkb[code]
}
