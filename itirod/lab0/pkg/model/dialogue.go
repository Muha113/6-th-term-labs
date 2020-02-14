package model

type Dialogue struct {
	ID       uint
	Partner  *User
	Messages []*Message
}
