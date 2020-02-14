package model

type Group struct {
	ID       uint
	Users    []*User
	Messages []*Message
}
