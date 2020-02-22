package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
	"os"
	"time"

	"github.com/Muha113/6-th-term-labs/isob/lab2/pkg/des"
	"github.com/Muha113/6-th-term-labs/isob/lab2/pkg/kerberos/common"
)

type Client struct {
	ID                string
	PasswordKey       string
	SessionKey        string
	GeneratedKey      string
	TGT               common.TicketGrantingTicket
	TGSClientResponse common.TGSClientResponse
	ASConn            net.Conn
	TGSConn           net.Conn
	ServerConn        net.Conn
}

func main() {
	client := &Client{}
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Name: ")
	name, err := reader.ReadString('\n')
	common.HandleError(err)
	fmt.Print("Password: ")
	passwd, err := reader.ReadString('\n')
	common.HandleError(err)
	client.ID = name
	client.PasswordKey = common.Hash(passwd)
	client.ASConn, err = net.Dial("tcp", "127.0.0.1:8000")
	common.HandleError(err)
	client.RequestAS()
}

func (c *Client) RequestAS() {
	req := common.ASClientRequest{
		TS:   time.Now().Format("2006-01-02 15:04:05"),
		ID:   c.ID,
		Req:  "TestServer",
		Addr: c.ASConn.LocalAddr(),
	}
	req.TS = des.Encode(req.TS, c.PasswordKey)
	msg, err := json.Marshal(req)
	common.HandleError(err)
	msg = append(msg, '\n')
	fmt.Println(string(msg))
	c.ASConn.Write(msg)
}
