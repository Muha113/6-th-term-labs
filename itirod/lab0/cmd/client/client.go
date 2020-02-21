package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/Muha113/6-th-term-labs/itirod/lab0/cmd/common"
	"github.com/sirupsen/logrus"
)

type ClientState int

const (
	INLOGIN ClientState = iota + 1
	INCHOOSINGDIALOGUE
	INCHOOSINGGROUP
	INCREATINGDIALOGUE
	INCREATINGGROUP
	INDIALOGUE
	INGROUP
	INGENERAL
	INMENU
)

type ConfType uint

const (
	TYPEDIALOGUE ConfType = iota + 1
	TYPEGROUP
)

type State struct {
	state ClientState
	id    uint
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

type Client struct {
	conn                  *net.UDPConn
	userID                uint
	userName              string
	state                 *State
	dialogues             map[uint][]Message
	groups                map[uint][]Message
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
		//fmt.Println("<Msg read success>")
		err = json.Unmarshal(buff[:bytes-1], &msg)
		handleError(common.FATAl, err)
		c.recievedMessagesQueue <- msg
	}
}

//complete
func (c *Client) sendMessage() {
	for {
		msg := <-c.sendingMessageQueue
		//fmt.Println("<Sending msg>")
		sendable, err := json.Marshal(msg)
		handleError(common.FATAl, err)
		sendable = append(sendable, byte('\n'))
		c.conn.Write(sendable)
	}
}

//complete
func (c *Client) printMessage() {
	for {
		msg := <-c.printMessagesQueue
		//fmt.Println("<Printing msg>")
		switch msg.MessageHeader.MessageType {
		case common.FUNC:
			switch common.FromStrToCmd[msg.Content] {
			case common.COMMANDCREATEDIALOGUE:
				if c.state.state == INCREATINGDIALOGUE {
					if msg.ErrorSending == nil {
						fmt.Println("Dialogue created successfull with id:", msg.MessageHeader.ID)
						fmt.Println()
					} else {
						fmt.Println("Some errors occured during creating dialogue")
						fmt.Println()
					}
				}
			case common.COMMANDCREATEGROUP:
				if c.state.state == INCREATINGGROUP {
					if msg.ErrorSending == nil {
						fmt.Println("Group created successfull with id:", msg.MessageHeader.ID)
						fmt.Println()
					} else {
						fmt.Println("Some errors occured during creating group")
						fmt.Println()
					}
				}
			default:
			}
		case common.LOGIN:
			if c.state.state == INLOGIN {
				fmt.Printf("%s\n", msg.Content)
			}
		case common.DIALOGUE:
			if c.state.state == INDIALOGUE && c.state.id == msg.MessageHeader.ID {
				if msg.MessageHeader.MessageRecievedStatus == 0 {
					fmt.Printf("#%s# %s: %s\n", msg.UserName, msg.Time, msg.Content)
				} else {
					switch msg.MessageHeader.MessageRecievedStatus {
					case common.RECIEVED:
						fmt.Println("<Message recieved>")
					case common.FAILED:
						fmt.Println("<Client failed to recieve message>")
					case common.TIMEOUT:
						fmt.Println("<Time out recieving message>")
					default:
					}
				}
			}
		case common.GROUP:
			if c.state.state == INGROUP && c.state.id == msg.MessageHeader.ID {
				fmt.Printf("#%s# %s: %s\n", msg.UserName, msg.Time, msg.Content)
			}
		case common.GENERAL:
			if c.state.state == INGENERAL {
				fmt.Printf("#%s# %s: %s\n", msg.UserName, msg.Time, msg.Content)
			}
		default:
		}
	}
}

//complete
func (c *Client) readInput() {
	reader := bufio.NewReader(os.Stdin)
	for {
		text, err := reader.ReadString('\n')
		handleError(common.ERROR, err)
		text = text[:len(text)-1]
		c.handleInput(text)
	}
}

//complete
func (c *Client) handleInput(input string) {
	switch c.state.state {
	case INMENU:
		choice, err := strconv.ParseInt(input, 10, 64)
		handleError(common.FATAl, err)
		switch common.Command(choice) {
		case common.COMMANDLOGIN:
			c.state.state = INLOGIN
			fmt.Println("Enter login")
		case common.COMMANDCLIENTINFO:
			c.printClientInfo()
			printMenu()
		case common.COMMANDCREATEDIALOGUE:
			c.state.state = INCREATINGDIALOGUE
			fmt.Println("Enter friend's id")
		case common.COMMANDCREATEGROUP:
			c.state.state = INCREATINGGROUP
			fmt.Println("Enter friends's ids")
		case common.COMMANDCHOOSEDIALOGUE:
			c.state.state = INCHOOSINGDIALOGUE
			fmt.Println("Enter dialogue id")
			for k := range c.dialogues {
				fmt.Println("Dialogue:", k)
			}
		case common.COMMANDCHOOSEGROUP:
			c.state.state = INCHOOSINGGROUP
			fmt.Println("Enter group id")
			for k := range c.groups {
				fmt.Println("Group:", k)
			}
		case common.COMMANDGENERAL:
			c.state.state = INGENERAL
			fmt.Println("Your are in general chat room")
		case common.COMMANDEXIT:
			//add exit handling
			os.Exit(1)
		default:
		}
	case INLOGIN:
		if isTransitPrevState(input) {
			c.state.state = INMENU
			printMenu()
		} else {
			m := Message{
				MessageHeader: MessageHeader{
					MessageType: common.LOGIN,
				},
				UserName: input,
			}
			c.sendingMessageQueue <- m
		}
	case INCREATINGDIALOGUE:
		if isTransitPrevState(input) {
			c.state.state = INMENU
			printMenu()
		} else {
			strSliceID := strings.Split(input, ",")
			uintSliceID := make([]uint, len(strSliceID))
			for _, v := range strSliceID {
				trimed := strings.TrimSpace(v)
				converted, err := strconv.ParseUint(trimed, 10, 64)
				handleError(common.FATAl, err)
				uintSliceID = append(uintSliceID, uint(converted))
			}
			uintSliceID = append(uintSliceID, c.userID)
			m := Message{
				MessageHeader: MessageHeader{
					MessageType: common.FUNC,
					CreateConf: common.CreateConfRequest{
						IDs: uintSliceID,
					},
				},
				UserID:   c.userID,
				UserName: c.userName,
				Content:  common.FromCmdToStr[common.COMMANDCREATEDIALOGUE],
			}
			c.sendingMessageQueue <- m
		}
	case INCREATINGGROUP:
		if isTransitPrevState(input) {
			c.state.state = INMENU
			printMenu()
		} else {
			strSliceID := strings.Split(input, ",")
			uintSliceID := make([]uint, len(strSliceID))
			for _, v := range strSliceID {
				trimed := strings.TrimSpace(v)
				converted, err := strconv.ParseUint(trimed, 10, 64)
				handleError(common.FATAl, err)
				uintSliceID = append(uintSliceID, uint(converted))
			}
			uintSliceID = append(uintSliceID, c.userID)
			m := Message{
				MessageHeader: MessageHeader{
					MessageType: common.FUNC,
					CreateConf: common.CreateConfRequest{
						IDs: uintSliceID,
					},
				},
				UserID:   c.userID,
				UserName: c.userName,
				Content:  common.FromCmdToStr[common.COMMANDCREATEGROUP],
			}
			c.sendingMessageQueue <- m
		}
	case INCHOOSINGDIALOGUE:
		if isTransitPrevState(input) {
			c.state.state = INMENU
			printMenu()
		} else {
			trimed := strings.TrimSpace(input)
			converted, err := strconv.ParseUint(trimed, 10, 64)
			handleError(common.FATAl, err)
			c.state.id = uint(converted)
			c.state.state = INDIALOGUE
			c.printConfMessages(TYPEDIALOGUE, c.state.id)
		}
	case INCHOOSINGGROUP:
		if isTransitPrevState(input) {
			c.state.state = INMENU
			printMenu()
		} else {
			trimed := strings.TrimSpace(input)
			converted, err := strconv.ParseUint(trimed, 10, 64)
			handleError(common.FATAl, err)
			c.state.id = uint(converted)
			c.state.state = INGROUP
			c.printConfMessages(TYPEGROUP, c.state.id)
		}
	case INDIALOGUE:
		if isTransitPrevState(input) {
			c.state.state = INMENU
			printMenu()
		} else {
			m := Message{
				MessageHeader: MessageHeader{
					MessageType: common.DIALOGUE,
					ID:          c.state.id,
				},
				UserID:   c.userID,
				UserName: c.userName,
				Time:     time.Now().Format("2006-01-02 15:04:05"),
				Content:  input,
			}
			c.sendingMessageQueue <- m
		}
	case INGROUP:
		if isTransitPrevState(input) {
			c.state.state = INMENU
			printMenu()
		} else {
			m := Message{
				MessageHeader: MessageHeader{
					MessageType: common.GROUP,
					ID:          c.state.id,
				},
				UserID:   c.userID,
				UserName: c.userName,
				Time:     time.Now().Format("2006-01-02 15:04:05"),
				Content:  input,
			}
			c.sendingMessageQueue <- m
		}
	case INGENERAL:
		if isTransitPrevState(input) {
			c.state.state = INMENU
			printMenu()
		} else {
			m := Message{
				MessageHeader: MessageHeader{
					MessageType: common.GENERAL,
				},
				UserID:   c.userID,
				UserName: c.userName,
				Time:     time.Now().Format("2006-01-02 15:04:05"),
				Content:  input,
			}
			c.sendingMessageQueue <- m
		}
	default:
	}
}

func (c *Client) printClientInfo() {
	if c.userID == 0 {
		fmt.Println("You are not logged in")
	} else {
		fmt.Println("\nInfo:")
		fmt.Println("Name:", c.userName)
		fmt.Println("Id:", c.userID)
		fmt.Println("Dialogues amount:", len(c.dialogues))
		fmt.Println("Groups amount:", len(c.groups))
		fmt.Println()
	}
}

//complete
func (c *Client) printConfMessages(ctype ConfType, id uint) {
	switch ctype {
	case TYPEDIALOGUE:
		for _, v := range c.dialogues[id] {
			c.printMessagesQueue <- v
		}
	case TYPEGROUP:
		for _, v := range c.groups[id] {
			c.printMessagesQueue <- v
		}
	default:
	}
}

//complete
func isTransitPrevState(input string) bool {
	if strings.Compare(input, ":back") == 0 {
		return true
	}
	return false
}

//complete
func (c *Client) handleRecieved() {
	for {
		msg := <-c.recievedMessagesQueue
		//fmt.Println("<Handling msg>")
		switch msg.MessageHeader.MessageType {
		case common.FUNC:
			switch common.FromStrToCmd[msg.Content] {
			case common.COMMANDCREATEDIALOGUE:
				if msg.ErrorSending == nil {
					responseID := msg.MessageHeader.ID
					c.dialogues[responseID] = make([]Message, 100)
				}
				c.printMessagesQueue <- msg
			case common.COMMANDCREATEGROUP:
				if msg.ErrorSending == nil {
					responseID := msg.MessageHeader.ID
					c.groups[responseID] = make([]Message, 100)
				}
				c.printMessagesQueue <- msg
			default:
			}
		case common.LOGIN:
			if msg.ErrorSending != nil {
				msg.Content = msg.ErrorSending.Error()
			} else {
				c.userID = msg.UserID
				c.userName = msg.UserName
			}
			c.printMessagesQueue <- msg
		case common.DIALOGUE:
			if msg.MessageHeader.MessageRecievedStatus == 0 {
				c.dialogues[msg.MessageHeader.ID] = append(c.dialogues[msg.MessageHeader.ID], msg)
				c.printMessagesQueue <- msg
				if msg.UserName != c.userName {
					msg.MessageHeader.MessageRecievedStatus = common.RECIEVED
					c.sendingMessageQueue <- msg
				}
			} else {
				if msg.UserName == c.userName {
					c.printMessagesQueue <- msg
				}
			}
		case common.GROUP:
			c.groups[msg.MessageHeader.ID] = append(c.groups[msg.MessageHeader.ID], msg)
			c.printMessagesQueue <- msg
		case common.GENERAL:
			c.printMessagesQueue <- msg
		default:
		}
	}
}

func printMenu() {
	fmt.Println("\nMenu:")
	fmt.Println("\n" + "++++++++++++")
	fmt.Println(common.COMMANDLOGIN, ": Login")
	fmt.Println(common.COMMANDCLIENTINFO, ": User info")
	fmt.Println(common.COMMANDCREATEDIALOGUE, ": Create dialogue")
	fmt.Println(common.COMMANDCREATEGROUP, ": Create group")
	fmt.Println(common.COMMANDCHOOSEDIALOGUE, ": Choose dialogue")
	fmt.Println(common.COMMANDCHOOSEGROUP, ": Choose group")
	fmt.Println(common.COMMANDGENERAL, ": Enter general chat room")
	fmt.Println(common.COMMANDEXIT, ": Exit")
	fmt.Println("++++++++++++" + "\n")
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
	client := &Client{
		state: &State{
			state: INMENU,
		},
		sendingMessageQueue:   make(chan Message, 20),
		recievedMessagesQueue: make(chan Message, 20),
		printMessagesQueue:    make(chan Message, 20),
		dialogues:             make(map[uint][]Message),
		groups:                make(map[uint][]Message),
	}
	client.conn, err = net.DialUDP("udp", nil, udpAddr)
	handleError(common.FATAl, err)
	defer client.conn.Close()
	printMenu()
	go client.sendMessage()

	go client.recieveMessage()

	go client.handleRecieved()

	go client.printMessage()

	client.readInput()
}
