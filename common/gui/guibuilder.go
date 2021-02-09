package gui

type PageType int

const (
	//PGrid Grid type (keypads)
	PGrid PageType = iota
	//PCircle type (joysticks)
	PCircle
	//PList List Type
	PList
)

type GUI struct {
	Pages []Page
}

type Page struct {
	Hive string
	Name string
	Keys []Key
	Type PageType
}

type Key struct {
	KeyID   int
	KeyName string
}
