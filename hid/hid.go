package hid

import "encoding/json"

type KeyMaps struct {
	USB   map[uint16]Key
	EVDEV map[uint16]Key
	XKB   map[uint16]Key
	Win   map[uint16]Key
	Mac   map[uint16]Key
	Code  map[string]Key
	Arr   []Key
}

type Key struct {
	USB   uint16
	Evdev uint16
	Xkb   uint16
	Win   uint16
	Mac   uint16
	Code  string
}

var Mappings KeyMaps = KeyMaps{}

//go:embed hid/generated.json
var file []byte

func init() {
	json.Unmarshal(file, &Mappings)
}
func GetWindowsFromHid(uv uint16) uint16 {
	return Mappings.USB[uv].Win
}
func GetHidFromWindows(uv uint16) uint16 {
	return Mappings.Win[uv].USB
}
func GetLinuxFromHid(uv uint16) uint16 {
	return Mappings.USB[uv].Evdev
}
func GetHidFromLinux(uv uint16) uint16 {
	return Mappings.EVDEV[uv].USB
}
