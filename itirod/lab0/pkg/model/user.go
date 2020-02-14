package model

type User struct {
	ID        uint
	Name      string
	Dialogues []uint
	Groups    []uint
}
