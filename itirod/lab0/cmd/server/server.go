package main

import (
	"encoding/json"
	"errors"
	"net"

	"github.com/sirupsen/logrus"

	"github.com/Muha113/6-th-term-labs/itirod/lab0/cmd/common"
)

type Server struct {
	conn       *net.UDPConn
	messages   chan Message
	clients    map[uint]*Client
	messagesDB []Message
	dialogues  map[uint][]*Conf
	groups     map[uint][]*Conf
}

type Conf struct {
	clients  []*Client
	messages []Message
}

type Client struct {
	id   uint
	name string
	addr *net.UDPAddr
}

type Message struct {
	MessageHeader MessageHeader `json:"header"`
	UserID        uint          `json:"id"`
	UserName      string        `json:"name"`
	Content       string        `json:"content"`
	Time          string        `json:"time"`
	RecieverAddr  *net.UDPAddr  `json:"addr"`
	ErrorSending  error         `json:"error"`
}

type MessageHeader struct {
	MessageType common.MessageType `json:"type"`
	ID          uint               `json:"id"`
}

//complete
func handleError(errType common.ErrorType, err error) {
	if err != nil {
		switch errType {
		case common.FATAl:
			logrus.Fatal(err)
		case common.ERROR:
			logrus.Error(err)
		default:
			logrus.Error(err)
		}
	}
}

//not complete
func (s *Server) handleMessage() {
	buffer := make([]byte, 1024)
	bytes, remoteAddr, err := s.conn.ReadFromUDP(buffer)
	handleError(common.ERROR, err)
	var message Message
	err = json.Unmarshal(buffer[:bytes-1], &message)
	handleError(common.ERROR, err)
	switch message.MessageHeader.MessageType {
	case common.FUNC:
	case common.LOGIN:
		if s.isLoggedIn(message) {
			message.Content = ""
			message.RecieverAddr = remoteAddr
			message.ErrorSending = errors.New("You have already logged in")
			s.messages <- message
		} else {
			s.clients[uint(len(s.clients)+1)] = &Client{
				id:   uint(len(s.clients)),
				name: message.UserName,
				addr: remoteAddr,
			}
			message.Content = "Login success!"
			message.RecieverAddr = remoteAddr
			message.ErrorSending = nil
			s.messages <- message
		}
	case common.DIALOGUE:
	case common.GROUP:
	case common.GENERAL:
		for _, v := range s.clients {
			message.RecieverAddr = v.addr
			s.messages <- message
		}
	default:
	}
}

//complete
func (s *Server) isLoggedIn(msg Message) bool {
	for _, v := range s.clients {
		if v.name == msg.UserName {
			return true
		}
	}
	return false
}

//complete
func (s *Server) sendMessage() {
	message := <-s.messages
	text, err := json.Marshal(message)
	handleError(common.ERROR, err)
	s.conn.WriteToUDP(text, message.RecieverAddr)
}

func main() {
	port := ":8000"
	udpAddress, err := net.ResolveUDPAddr("udp4", port)
	handleError(common.ERROR, err)
	var s Server
	s.messages = make(chan Message, 40)
	s.clients = make(map[uint]*Client)
	s.conn, err = net.ListenUDP("udp", udpAddress)
	handleError(common.ERROR, err)
	defer s.conn.Close()

	go s.sendMessage()

	for {
		s.handleMessage()
	}
}
