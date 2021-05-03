package hid

import (
	_ "embed"
	"encoding/json"
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

var Mappings KeyMaps = KeyMaps{}

//go:embed generated.json
var file []byte

func init() {
	json.Unmarshal(file, &Mappings)
}
func GetWindowsFromHid(uv uint16) uint16 {
	return Mappings.Usb[uv].Win
}
func GetHidFromWindows(uv uint16) uint16 {
	return Mappings.Win[uv].Usb
}

func GetLinuxFromHid(uv uint16) uint16 {
	return Mappings.Usb[uv].Evdev
}
func GetHidFromLinux(uv uint16) uint16 {
	return Mappings.Evdev[uv].Usb
}

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
