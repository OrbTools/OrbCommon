package gui

//PageType PGrid, PCircle or PList
type PageType int

const (
	//PGrid Grid type (keypads)
	PGrid PageType = iota
	//PCircle type (joysticks)
	PCircle
	//PList List Type
	PList
)

//GUI a GUI tree
type GUI struct {
	Pages []*Page
}

//Page A page of a binding GUI
type Page struct {
	Hive string
	Name string
	Keys []*Key
	Type PageType
}

//Key a keybind
type Key struct {
	KeyID   int
	KeyName string
	Default uint16
}
