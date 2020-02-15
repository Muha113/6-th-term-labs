package common

type MessageType int

const (
	FUNC MessageType = iota + 1
	LOGIN
	GROUP
	DIALOGUE
	GENERAL
)

type ErrorType int

const (
	FATAl ErrorType = iota + 1
	ERROR
)
