package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
	"os"

	"github.com/Muha113/6-th-term-labs/itirod/lab0/cmd/common"
	"github.com/sirupsen/logrus"
)

type ClientState int

const (
	INLOGIN ClientState = iota + 1
	INDIALOGUE
	INGROUP
	INGENERAL
	INMENU
)

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

type Client struct {
	conn                  *net.UDPConn
	userID                uint
	userName              string
	state                 ClientState
	dialogues             map[uint]Message
	groups                map[uint]Message
	sendingMessageQueue   chan Message
	recievedMessagesQueue chan Message
	printMessagesQueue    chan Message
}

//complete
func (c *Client) recieveMessage() {
	var msg Message
	buff := make([]byte, 2048)
	for {
		bytes, _, err := c.conn.ReadFromUDP(buff)
		handleError(common.ERROR, err)
		err = json.Unmarshal(buff[:bytes-1], &msg)
		c.recievedMessagesQueue <- msg
	}
}

//complete
func (c *Client) sendMessage() {
	for {
		msg := <-c.sendingMessageQueue
		sendable, err := json.Marshal(msg)
		handleError(common.FATAl, err)
		sendable = append(sendable, byte('\n'))
		c.conn.Write(sendable)
	}
}

//seems to be complete
func (c *Client) printMessage() {
	msg := <-c.printMessagesQueue
	switch msg.MessageHeader.MessageType {
	case common.LOGIN:
		if c.state == INLOGIN {
			fmt.Printf("%s\n", msg.Content)
		}
	case common.DIALOGUE:
		if c.state == INDIALOGUE {
			fmt.Printf("#%s# %s: %s", msg.UserName, msg.Time, msg.Content)
		}
	case common.GROUP:
		if c.state == INGROUP {
			fmt.Printf("#%s# %s: %s", msg.UserName, msg.Time, msg.Content)
		}
	case common.GENERAL:
		if c.state == INGENERAL {
			fmt.Printf("#%s# %s: %s", msg.UserName, msg.Time, msg.Content)
		}
	default:
	}
}

//complete
func (c *Client) readInput() {
	reader := bufio.NewReader(os.Stdin)
	for {
		text, err := reader.ReadString('\n')
		handleError(common.ERROR, err)
		c.handleInput(text)
	}
}

//not complete
//it will handle menu commands and switching client state
func (c *Client) handleInput(input string) {
}

//not complete
//it will handle errors and message logic (ex: saving)
func (c *Client) handleRecieved() {
	msg := <-c.recievedMessagesQueue
	c.printMessagesQueue <- msg
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

func main() {
	udpAddr, err := net.ResolveUDPAddr("udp4", "127.0.0.1:8000")
	handleError(common.FATAl, err)
	var name string
	fmt.Println("Enter username:")
	_, err = fmt.Scan(&name)
	handleError(common.FATAl, err)
	client := &Client{
		userName:              name,
		state:                 INMENU,
		sendingMessageQueue:   make(chan Message, 20),
		recievedMessagesQueue: make(chan Message, 20),
		printMessagesQueue:    make(chan Message, 20),
	}
	client.conn, err = net.DialUDP("udp", nil, udpAddr)
	handleError(common.FATAl, err)
	defer client.conn.Close()

	go client.sendMessage()

	go client.recieveMessage()

	go client.printMessage()

	client.readInput()
}
