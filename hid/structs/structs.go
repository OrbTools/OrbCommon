package structs

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
