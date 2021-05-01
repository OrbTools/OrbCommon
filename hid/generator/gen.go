package main

import (
	"encoding/json"
	"io"
	"io/fs"
	"os"
	"regexp"
	"strconv"
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

func main() {
	rege, _ := regexp.Compile("DOM_CODE\\(0x[0-9a-f]{2}([0-9a-f]*), 0x([0-9a-f]*), 0x([0-9a-f]*), 0x([0-9a-f]*), 0x([0-9a-f]*), \"?[A-Za-z]*\"?, ([A-Za-z_]*)")
	//DOM_CODE(USB, evdev, XKB, Win, Mac, _, Code)
	fil, _ := os.OpenFile("hid/data/keycode_data.inc", 0, fs.FileMode(os.O_RDONLY))
	byts, _ := io.ReadAll(fil)
	fil.Close()
	matches := rege.FindAllSubmatch(byts, -1)
	KeyMaps := KeyMaps{
		Usb:   make(map[uint16]Key),
		Evdev: make(map[uint16]Key),
		Xkb:   make(map[uint16]Key),
		Win:   make(map[uint16]Key),
		Mac:   make(map[uint16]Key),
		Code:  make(map[string]Key),
	}
	Arr := make([]Key, 0)
	for _, bar := range matches {
		U, _ := strconv.ParseUint(string(bar[1]), 16, 16)
		E, _ := strconv.ParseUint(string(bar[2]), 16, 16)
		X, _ := strconv.ParseUint(string(bar[3]), 16, 16)
		W, _ := strconv.ParseUint(string(bar[4]), 16, 16)
		M, _ := strconv.ParseUint(string(bar[5]), 16, 16)
		Keys := Key{
			Usb:   uint16(U),
			Evdev: uint16(E),
			Xkb:   uint16(X),
			Win:   uint16(W),
			Mac:   uint16(M),
			Code:  string(bar[6]),
		}
		KeyMaps.Usb[uint16(U)] = Keys
		KeyMaps.Evdev[uint16(E)] = Keys
		KeyMaps.Xkb[uint16(X)] = Keys
		KeyMaps.Win[uint16(W)] = Keys
		KeyMaps.Mac[uint16(M)] = Keys
		KeyMaps.Code[string(bar[6])] = Keys
		Arr = append(Arr, Keys)
	}
	KeyMaps.Arr = Arr
	out, _ := os.Create("hid/generated.json")
	defer out.Close()
	jso, _ := json.Marshal(KeyMaps)
	out.Write(jso)
}
