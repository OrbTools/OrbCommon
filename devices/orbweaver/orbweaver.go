package orbweaver

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
