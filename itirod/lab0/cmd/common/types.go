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

type Command int

const (
	COMMANDLOGIN Command = iota + 1
	COMMANDCLIENTINFO
	COMMANDCREATEDIALOGUE
	COMMANDCREATEGROUP
	COMMANDCHOOSEDIALOGUE
	COMMANDCHOOSEGROUP
	COMMANDGENERAL
	COMMANDEXIT
)

var FromCmdToStr = map[Command]string{
	COMMANDCREATEDIALOGUE: "@create_dialogue",
	COMMANDCREATEGROUP:    "@create_group",
}

var FromStrToCmd = map[string]Command{
	"@create_dialogue": COMMANDCREATEDIALOGUE,
	"@create_group":    COMMANDCREATEGROUP,
}

type CreateConfRequest struct {
	IDs []uint
}
