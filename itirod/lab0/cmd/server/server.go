package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"strings"

	"github.com/sirupsen/logrus"

	"github.com/Muha113/6-th-term-labs/itirod/lab0/cmd/common"
)

type Server struct {
	conn      *net.UDPConn
	messages  chan Message
	clients   map[uint]*Client
	dialogues map[uint]*Conf
	groups    map[uint]*Conf
}

type Conf struct {
	clients  map[uint]*Client
	messages []Message
}

type Client struct {
	id          uint
	name        string
	dialoguesID []uint
	groupsID    []uint
	addr        *net.UDPAddr
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
	MessageType           common.MessageType       `json:"type"`
	MessageRecievedStatus common.MessageStatusCode `json:"status"`
	CreateConf            common.CreateConfRequest `json:"createConf"`
	ID                    uint                     `json:"id"`
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

//seems to be complete
//add checking isLoggedIn
func (s *Server) handleMessage() {
	buffer := make([]byte, 1024)
	bytes, remoteAddr, err := s.conn.ReadFromUDP(buffer)
	handleError(common.ERROR, err)
	fmt.Println("Handling msg from", remoteAddr.String())
	var message Message
	err = json.Unmarshal(buffer[:bytes-1], &message)
	handleError(common.ERROR, err)
	switch message.MessageHeader.MessageType {
	case common.FUNC:
		switch common.FromStrToCmd[message.Content] {
		case common.COMMANDCREATEDIALOGUE:
			id := uint(len(s.dialogues) + 1)
			s.dialogues[id] = &Conf{
				clients:  make(map[uint]*Client),
				messages: make([]Message, 100),
			}
			for _, v := range message.MessageHeader.CreateConf.IDs {
				val, ok := s.clients[v]
				if ok {
					s.dialogues[id].clients[val.id] = val
					message.RecieverAddr = val.addr
					message.ErrorSending = nil
					message.MessageHeader.ID = id
					s.messages <- message
				}
			}
			//TODO: make response msg. UPD: done
			// message.RecieverAddr = remoteAddr
			// message.ErrorSending = nil
			// message.MessageHeader.ID = id
			// s.messages <- message
		case common.COMMANDCREATEGROUP:
			id := uint(len(s.groups) + 1)
			s.groups[id] = &Conf{
				clients:  make(map[uint]*Client),
				messages: make([]Message, 100),
			}
			for _, v := range message.MessageHeader.CreateConf.IDs {
				val, ok := s.clients[v]
				if ok {
					s.groups[id].clients[val.id] = val
					message.RecieverAddr = val.addr
					message.ErrorSending = nil
					message.MessageHeader.ID = id
					s.messages <- message
				}
			}
			//TODO: make response msg. UPD: done
			// message.RecieverAddr = remoteAddr
			// message.ErrorSending = nil
			// message.MessageHeader.ID = id
			// s.messages <- message
		default:
		}
	case common.LOGIN:
		if s.isLoggedIn(message) {
			message.Content = ""
			message.RecieverAddr = remoteAddr
			message.ErrorSending = errors.New("You have already logged in")
			s.messages <- message
		} else {
			tmp := uint(len(s.clients) + 1)
			s.clients[tmp] = &Client{
				id:   tmp,
				name: message.UserName,
				addr: remoteAddr,
			}
			message.UserID = tmp
			message.Content = "Login success!"
			message.RecieverAddr = remoteAddr
			message.ErrorSending = nil
			s.messages <- message
		}
	case common.DIALOGUE:
		tmp := message.MessageHeader.ID
		s.dialogues[tmp].messages = append(s.dialogues[tmp].messages, message)
		for _, v := range s.dialogues[tmp].clients {
			message.RecieverAddr = v.addr
			s.messages <- message
		}
	case common.GROUP:
		tmp := message.MessageHeader.ID
		s.groups[tmp].messages = append(s.groups[tmp].messages, message)
		for _, v := range s.groups[tmp].clients {
			message.RecieverAddr = v.addr
			s.messages <- message
		}
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
		if strings.Compare(v.name, msg.UserName) == 0 {
			return true
		}
	}
	return false
}

//complete
func (s *Server) sendMessage() {
	for {
		message := <-s.messages
		fmt.Println("Sending msg to", message.RecieverAddr.String())
		text, err := json.Marshal(message)
		handleError(common.ERROR, err)
		text = append(text, byte('\n'))
		s.conn.WriteToUDP(text, message.RecieverAddr)
		fmt.Println("Message sent to", message.RecieverAddr.String())
	}
}

func main() {
	var s Server
	var err error
	s.messages = make(chan Message, 40)
	s.clients = make(map[uint]*Client)
	s.dialogues = make(map[uint]*Conf)
	s.groups = make(map[uint]*Conf)
	fmt.Println("Starting server...")
	udpAddr := &net.UDPAddr{
		Port: 8000,
		IP:   net.ParseIP("127.0.0.1"),
	}
	s.conn, err = net.ListenUDP("udp", udpAddr)
	handleError(common.FATAl, err)
	fmt.Println("Status OK. Listening on", s.conn.LocalAddr())
	defer s.conn.Close()

	go s.sendMessage()

	for {
		s.handleMessage()
	}
}
