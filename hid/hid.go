package hid

import (
	"bytes"
	_ "embed"
	"github.com/Minizbot2012/minxdr"
)

type KeyMaps struct {
	Usb   map[uint16]Key
	Evdev map[uint16]Key
	Xkb   map[uint16]Key
	Win   map[uint16]Key
	Mac   map[uint16]Key
	Code  map[string]Key
	Arr   []Key
}

type Key struct {
	Usb   uint16
	Evdev uint16
	Xkb   uint16
	Win   uint16
	Mac   uint16
	Code  string
}

//go:embed generated.bin
var file []byte

func init() {
	minxdr.Unmarshal(bytes.NewReader(file), &Mappings)
}

var Mappings KeyMaps = KeyMaps{}

func GetMappingFromHID(uv uint16) Key {
	return Mappings.Usb[uv]
}

func GetMappingFromWindows(uv uint16) Key {
	return Mappings.Win[uv]
}

func GetMappingFromLinux(uv uint16) Key {
	return Mappings.Evdev[uv]
}

func GetMappingFromName(name string) Key {
	return Mappings.Code[name]
}

func GetMappingFromX(code uint16) Key {
	return Mappings.Xkb[code]
}
