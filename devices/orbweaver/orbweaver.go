package orbweaver

import (
	"encoding/binary"
	"io"
	"os"

	"github.com/OrbTools/OrbCommon/common"
)

//BINDING Map of Default Keys
var BINDING = [...]byte{41, 2, 3, 4, 5, 15, 16, 17, 18, 19, 58, 30, 31, 32, 33, 42, 44, 45, 46, 47, 56, 103, 106, 108, 105, 57}

//GUI gui definition for an orbweaver
var GUI = &common.GUI{
	Pages: []common.Page{
		{
			Hive: "MIP",
			Name: "Grid",
			Type: common.PGrid,
			Keys: []common.Key{
				{
					KeyID:   0,
					KeyName: "01",
				}, {
					KeyID:   1,
					KeyName: "02",
				}, {
					KeyID:   3,
					KeyName: "04",
				}, {
					KeyID:   4,
					KeyName: "05",
				}, {
					KeyID:   5,
					KeyName: "06",
				}, {
					KeyID:   6,
					KeyName: "07",
				}, {
					KeyID:   7,
					KeyName: "08",
				}, {
					KeyID:   8,
					KeyName: "09",
				}, {
					KeyID:   9,
					KeyName: "10",
				}, {
					KeyID:   10,
					KeyName: "11",
				}, {
					KeyID:   11,
					KeyName: "12",
				}, {
					KeyID:   12,
					KeyName: "13",
				}, {
					KeyID:   13,
					KeyName: "11",
				}, {
					KeyID:   14,
					KeyName: "11",
				}, {
					KeyID:   15,
					KeyName: "11",
				}, {
					KeyID:   16,
					KeyName: "11",
				}, {
					KeyID:   17,
					KeyName: "11",
				}, {
					KeyID:   18,
					KeyName: "11",
				}, {
					KeyID:   19,
					KeyName: "20",
				},
			},
		}, {
			Hive: "SIP",
			Name: "Side Keys",
			Type: common.PList,
			Keys: []common.Key{
				{
					KeyID:   0,
					KeyName: "Upper Button",
				}, {
					KeyID:   1,
					KeyName: "DPad Up",
				}, {
					KeyID:   2,
					KeyName: "DPad Right",
				}, {
					KeyID:   3,
					KeyName: "DPad Left",
				}, {
					KeyID:   4,
					KeyName: "DPad Down",
				}, {
					KeyID:   5,
					KeyName: "Lower Button",
				},
			},
		},
	},
}

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
func SavePKMKeymap(mapped *PKM, file io.WriteCloser) {
	binary.Write(file, binary.LittleEndian, mapped)
	file.Close()
}

//LoadPKMKeymap loads an orb for editing
func LoadPKMKeymap(file io.ReadCloser) *PKM {
	mapped := new(PKM)
	binary.Read(file, binary.LittleEndian, mapped)
	file.Close()
	return mapped
}
