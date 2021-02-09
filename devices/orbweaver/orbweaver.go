package orbweaver

import (
	"encoding/binary"
	"os"

	"fyne.io/fyne"
)

//BINDING Map of Default Keys
var BINDING = [...]byte{41, 2, 3, 4, 5, 15, 16, 17, 18, 19, 58, 30, 31, 32, 33, 42, 44, 45, 46, 47, 56, 103, 106, 108, 105, 57}

//KeyMap singular keymap
type KeyMap struct {
	Keymap [26]uint16
	Color  [3]byte
}

//KeyMaps a set of keymaps
type KeyMaps struct {
	Maps       [7]*KeyMap
	Currentmap int
	MCount     int
}

//PKM format for altering the keymap
type PKM struct {
	MIP [20]uint16
	SIP [6]uint16
	COL [3]byte
}

//LoadKM Load Orbmap KM structure
func LoadKM(file string) *KeyMap {
	mapped := new(KeyMap)
	of, _ := os.Open(file)
	binary.Read(of, binary.LittleEndian, mapped)
	return mapped
}

//SavePKMKeymap saves an orb after edit
func SavePKMKeymap(mapped *PKM, file fyne.URIWriteCloser) {
	binary.Write(file, binary.LittleEndian, mapped)
	file.Close()
}

//LoadPKMKeymap loads an orb for editing
func LoadPKMKeymap(file fyne.URIReadCloser) *PKM {
	mapped := new(PKM)
	binary.Read(file, binary.LittleEndian, mapped)
	file.Close()
	return mapped
}
