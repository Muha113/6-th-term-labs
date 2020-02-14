package main

import (
	"fmt"
	"net"
	"strings"

	"github.com/Muha113/6-th-term-labs/itirod/lab0/pkg/model"

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

type Client struct {
	User   *model.User
	Server *Server
	Addr   *net.UDPAddr
}

type Server struct {
	Listener  *net.UDPConn
	Clients   map[uint]*Client
	Dialogues []*model.Dialogue
	Groups    []*model.Group
}

func main() {
	fmt.Println("Starting server")
	listener, err := net.ListenUDP("udp", &net.UDPAddr{
		Port: 8000,
		IP:   net.ParseIP("127.0.0.1"),
	})
	errorHandler(TypeNotFatal, err)
	defer listener.Close()
	serv := &Server{
		Listener:  listener,
		Clients:   make(map[uint]*Client),
		Dialogues: make([]*model.Dialogue, 1),
		Groups:    make([]*model.Group, 1),
	}
	fmt.Println("Listening on", serv.Listener.LocalAddr())
	buffer := make([]byte, 2048)
	for {
		bytes, addr, err := serv.Listener.ReadFromUDP(buffer)
		errorHandler(TypeNotFatal, err)
		command, _ := commandsMapper[string(buffer[:bytes-1])]
		switch command { // TODO: goroutines
		case CommandRegister:
			//register(serv, addr, uint(len(serv.Clients)+1))
			//fmt.Println(len(serv.Clients))
		case CommandLogin:
			login(serv, addr, uint(len(serv.Clients)+1))
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

func register(server *Server, addr *net.UDPAddr, id uint) {
	buffer := make([]byte, 4096)
	client := &Client{
		Addr:   addr,
		Server: server,
		User: &model.User{
			ID:        id,
			Dialogues: make([]uint, 1),
			Groups:    make([]uint, 1),
		},
	}
	server.Listener.WriteToUDP([]byte("Login:"+"\n"), addr)
	bytes, _, err := server.Listener.ReadFromUDP(buffer)
	errorHandler(TypeNotFatal, err)
	name := string(buffer[:bytes-1])
	client.User.Name = name
	server.Clients[id] = client
	server.Listener.WriteToUDP([]byte("Register success!"+"\n"), addr)
}

func login(server *Server, addr *net.UDPAddr, id uint) {
	buffer := make([]byte, 4096)
	client := &Client{
		Addr:   addr,
		Server: server,
		User: &model.User{
			ID:        id,
			Dialogues: make([]uint, 1),
			Groups:    make([]uint, 1),
		},
	}
	server.Listener.WriteToUDP([]byte("Login:"+"\n"), addr)
	bytes, _, err := server.Listener.ReadFromUDP(buffer)
	errorHandler(TypeNotFatal, err)
	name := string(buffer[:bytes-1])
	client.User.Name = name
	server.Clients[id] = client
	server.Listener.WriteToUDP([]byte("Login success!"+"\n"), addr)
}

func getAllDialogues() {
	fmt.Println("getAllDialogues")
}

func getAllUsersOnline() {
	fmt.Println("getAllUsersOnline")
}

func createDialogue() {
	fmt.Println("createDialogue")
}

func createGroup() {
	fmt.Println("createGroup")
}
