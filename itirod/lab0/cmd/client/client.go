package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
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

var commandsMapper = map[int]string{
	CommandRegister:          "@register",
	CommandLogin:             "@login",
	CommandGetAllDialogues:   "@dialogues",
	CommandGetAllUsersOnline: "@users",
	CommandCreateDialogue:    "@createdialogue",
	CommandCreateGroup:       "@creategroup",
}

func main() {
	addr := &net.UDPAddr{
		Port: 8000,
		IP:   net.ParseIP("127.0.0.1"),
	}
	dialer, err := net.DialUDP("udp", nil, addr)
	defer dialer.Close()
	errorHandler(TypeNotFatal, err)
	//buffer := make([]byte, 2048)
	for {
		reader := bufio.NewReader(os.Stdin)
		printMainMenu()
		text, err := reader.ReadString('\n')
		errorHandler(TypeNotFatal, err)
		command, err := strconv.ParseInt(text[:len(text)-1], 10, 64)
		errorHandler(TypeNotFatal, err)

		switch command {
		case CommandRegister:
			dialer.Write([]byte(commandsMapper[CommandRegister] + "\n"))
		case CommandLogin:
			login(dialer)
		case CommandGetAllDialogues:
			dialer.Write([]byte(commandsMapper[CommandGetAllDialogues] + "\n"))
		case CommandGetAllUsersOnline:
			dialer.Write([]byte(commandsMapper[CommandGetAllUsersOnline] + "\n"))
		case CommandCreateDialogue:
			dialer.Write([]byte(commandsMapper[CommandCreateDialogue] + "\n"))
		case CommandCreateGroup:
			dialer.Write([]byte(commandsMapper[CommandCreateGroup] + "\n"))
		case CommandExit:
			os.Exit(1)
		default:
			fmt.Println("There is no such command")
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

func printMainMenu() {
	fmt.Println("\nMenu:")
	fmt.Println("\n" + "+++++++++++++++++++++")
	fmt.Println(CommandRegister, ": Register")
	fmt.Println(CommandLogin, ": Login")
	fmt.Println(CommandGetAllDialogues, ": Get all dialogues")
	fmt.Println(CommandGetAllUsersOnline, ": Get all users online")
	fmt.Println(CommandCreateDialogue, ": Start chating with smbd")
	fmt.Println(CommandCreateGroup, ": Create group")
	fmt.Println(CommandExit, ": Exit chat app")
	fmt.Println("+++++++++++++++++++++" + "\n")
}

func register() {}

func login(conn *net.UDPConn) {
	conn.Write([]byte(commandsMapper[CommandLogin] + "\n"))
	buffer := make([]byte, 2048)
	bytes, _, err := conn.ReadFromUDP(buffer)
	errorHandler(TypeNotFatal, err)
	msg := buffer[:bytes-1]
	fmt.Println(string(msg))
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	errorHandler(TypeNotFatal, err)
	conn.Write([]byte(text))
	bytes, _, err = conn.ReadFromUDP(buffer)
	errorHandler(TypeNotFatal, err)
	fmt.Println(string(buffer[:bytes-1]))
}
