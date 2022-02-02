package main

import (
	"github.com/Minizbot2012/minxdr"
	"github.com/OrbTools/OrbCommon/hid/structs"
	"io"
	"io/fs"
	"os"
	"regexp"
	"strconv"
)

func main() {
	rege, _ := regexp.Compile("DOM_CODE\\(0x07([0-9a-f]*), 0x([0-9a-f]*), 0x([0-9a-f]*), 0x([0-9a-f]*), 0x([0-9a-f]*), \"?[A-Za-z0-9]*\"?, ([A-Za-z_0-9]*)")
	//DOM_CODE(USB, evdev, XKB, Win, Mac, _, Code)
	fil, _ := os.OpenFile("hid/data/keycode_data.inc", 0, fs.FileMode(os.O_RDONLY))
	byts, _ := io.ReadAll(fil)
	fil.Close()
	matches := rege.FindAllSubmatch(byts, -1)
	KeyMaps := structs.KeyMaps{
		Usb:   make(map[uint16]structs.Key),
		Evdev: make(map[uint16]structs.Key),
		Xkb:   make(map[uint16]structs.Key),
		Win:   make(map[uint16]structs.Key),
		Mac:   make(map[uint16]structs.Key),
		Code:  make(map[string]structs.Key),
	}
	Arr := make([]structs.Key, 0)
	for _, bar := range matches {
		U, _ := strconv.ParseUint(string(bar[1]), 16, 16)
		E, _ := strconv.ParseUint(string(bar[2]), 16, 16)
		X, _ := strconv.ParseUint(string(bar[3]), 16, 16)
		W, _ := strconv.ParseUint(string(bar[4]), 16, 16)
		M, _ := strconv.ParseUint(string(bar[5]), 16, 16)
		Keys := structs.Key{
			Usb:   uint16(U),
			Evdev: uint16(E),
			Xkb:   uint16(X),
			Win:   uint16(W),
			Mac:   uint16(M),
			Code:  string(bar[6]),
		}
		if _, ok := KeyMaps.Usb[uint16(U)]; !ok {
			KeyMaps.Usb[uint16(U)] = Keys
		}
		if _, ok := KeyMaps.Evdev[uint16(E)]; !ok {
			KeyMaps.Evdev[uint16(E)] = Keys
		}
		if _, ok := KeyMaps.Xkb[uint16(X)]; !ok {
			KeyMaps.Xkb[uint16(X)] = Keys
		}
		if _, ok := KeyMaps.Win[uint16(W)]; !ok {
			KeyMaps.Win[uint16(W)] = Keys
		}
		if _, ok := KeyMaps.Mac[uint16(M)]; !ok {
			KeyMaps.Mac[uint16(M)] = Keys
		}
		KeyMaps.Code[string(bar[6])] = Keys
		Arr = append(Arr, Keys)
	}
	KeyMaps.Arr = Arr
	xdo, _ := os.Create("hid/generated.bin")
	defer xdo.Close()
	minxdr.Marshal(xdo, KeyMaps)
}
