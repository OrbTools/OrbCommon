// +build !xdr

package hid

import (
	_ "embed"
	"encoding/json"
)

//go:embed generated.json
var file []byte

func init() {
	json.Unmarshal(file, &Mappings)
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

func GetMappingFromX(code uint16) Key {
	return Mappings.Xkb[code]
}
