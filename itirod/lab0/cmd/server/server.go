package main

import (
	"fmt"
	"net"
	"strconv"
	"strings"

	"github.com/goapt/logger"
)

const (
	TypeFatal    = "fatal"
	TypeNotFatal = "nofatal"
)

const (
	CommandRegister = iota + 1
	CommandLogin
	CommandGetAllDialogues
	CommandGetAllUsersOnline
	CommandCreateDialogue
	CommandCreateGroup
	CommandExit
)

var commandsMapper = map[string]int{
	"@register":       CommandRegister,
	"@login":          CommandLogin,
	"@dialogues":      CommandGetAllDialogues,
	"@users":          CommandGetAllUsersOnline,
	"@createdialogue": CommandCreateDialogue,
	"@creategroup":    CommandCreateGroup,
}

func main() {
	fmt.Println("Starting server")
	listener, err := net.ListenUDP("udp", &net.UDPAddr{
		Port: 8000,
		IP:   net.ParseIP("127.0.0.1"),
	})
	errorHandler(TypeNotFatal, err)
	defer listener.Close()
	fmt.Println("Listening on", listener.LocalAddr())
	buffer := make([]byte, 2048)
	for {
		bytes, _, err := listener.ReadFromUDP(buffer)
		errorHandler(TypeNotFatal, err)
		command, err := strconv.ParseInt(string(buffer[:bytes-1]), 10, 64)
		errorHandler(TypeNotFatal, err)
		switch command { // TODO: goroutines
		case CommandRegister:
			register()
		case CommandLogin:
			login()
		case CommandGetAllDialogues:
			getAllDialogues()
		case CommandGetAllUsersOnline:
			getAllUsersOnline()
		case CommandCreateDialogue:
			createDialogue()
		case CommandCreateGroup:
			createGroup()
		default:
			fmt.Println("No such command") // TODO: write to udp socket
		}
	}
}

func errorHandler(errType string, err error) {
	if err != nil {
		if strings.Compare(errType, TypeFatal) == 0 {
			logger.Fatal(err)
		} else {
			logger.Error(err)
		}
	}
}
