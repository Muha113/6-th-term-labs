package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
	"os"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/Muha113/6-th-term-labs/isob/lab2/pkg/des"
	"github.com/Muha113/6-th-term-labs/isob/lab2/pkg/kerberos/common"
)

type Client struct {
	ID                         string
	PasswordKey                string
	SessionKey                 string
	GeneratedKey               string
	TGTEncrypted               string
	TimeStamp                  string
	ExpTime                    string
	TGSServerResponseEncrypted string
	ASConn                     net.Conn
	TGSConn                    net.Conn
	ServerConn                 net.Conn
	ASAddr                     string
	TGSAddr                    string
	ServerAddr                 string
}

func main() {
	client := &Client{}
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Name: ")
	name, err := reader.ReadString('\n')
	common.HandleError(err, "Client-38: ")
	fmt.Print("Password: ")
	passwd, err := reader.ReadString('\n')
	common.HandleError(err, "Client-41: ")
	client.ID = name[:len(name)-1]
	client.PasswordKey = common.Hash(passwd[:len(passwd)-1])
	client.ASConn, err = net.Dial("tcp", "127.0.0.1:8000")
	common.HandleError(err, "Client-45: ")
	client.RequestAS()
	client.ResponseAS()
	client.ASConn.Close()
	client.TGSConn, err = net.Dial("tcp", client.TGSAddr)
	common.HandleError(err, "Client-50: ")
	client.RequestTGS()
	client.ResponseTGS()
	client.TGSConn.Close()
	client.ServerConn, err = net.Dial("tcp", "127.0.0.1:8002")
	common.HandleError(err, "Client-55: ")
	client.RequestServer()
	client.ResponseServer()
	client.ServerConn.Close()
}

func (c *Client) RequestAS() {
	fmt.Println("+++++ Requesting AS... +++++")
	req := common.ASClientRequest{
		TS:   time.Now().Format("2006-01-02 15:04:05"),
		ID:   c.ID,
		Req:  "TestServer",
		Addr: c.ASConn.LocalAddr().String(),
	}
	//logrus.Error("KeyByte:", []byte(c.PasswordKey), " KeyStr:", c.PasswordKey, " Len:", len(c.PasswordKey))
	req.TS = des.Encode(req.TS, c.PasswordKey)
	msg, err := json.Marshal(req)
	common.HandleError(err, "Client-70: ")
	fmt.Println("----- Body request encrypted:", string(common.PrettyPrint(msg)))
	msg = append(msg, '\n')
	c.ASConn.Write(msg)
}

func (c *Client) ResponseAS() {
	fmt.Println("+++++ Response AS... +++++")
	buff := make([]byte, 1024)
	bytes, err := c.ASConn.Read(buff)
	common.HandleError(err, "Client-79: ")
	fmt.Println("----- Body response encrypted:", string(common.PrettyPrint(buff[:bytes-1])))
	var resp common.ASResponseEncrypted
	err = json.Unmarshal(buff[:bytes-1], &resp)
	common.HandleError(err, "Client-82: ")
	fmt.Println("----- Body response ASCliResp:", string(common.PrettyPrint([]byte(des.Decode(resp.ASClientResponse, c.PasswordKey)))))
	var cliResp common.ASClientResponse
	err = json.Unmarshal([]byte(des.Decode(resp.ASClientResponse, c.PasswordKey)), &cliResp)
	common.HandleError(err, "Client-85: ")
	c.TGTEncrypted = resp.TGT
	c.SessionKey = cliResp.SessionKey
	c.TGSAddr = cliResp.TGSAddr
	c.TimeStamp = cliResp.TS
	c.ExpTime = cliResp.Exp
}

func (c *Client) RequestTGS() {
	fmt.Println("+++++ Request TGS... +++++")
	req := common.TGSClientRequest{
		ID:  c.ID,
		TS:  c.TimeStamp,
		Req: "TestServer",
	}
	reqJSON, err := json.Marshal(req)
	common.HandleError(err, "Client-100: ")
	fmt.Println("----- Body request:", string(common.PrettyPrint(reqJSON)))
	//logrus.Error([]byte(c.SessionKey), " str:", c.SessionKey)
	reqEncrypted := des.Encode(string(reqJSON), c.SessionKey)
	msg := common.TGSRequestEncrypted{
		TGSClientRequest: reqEncrypted,
		TGT:              c.TGTEncrypted,
	}
	msgJSON, err := json.Marshal(msg)
	common.HandleError(err, "Client-107: ")
	fmt.Println("----- Body request encrypted:", string(common.PrettyPrint(msgJSON)))
	msgJSON = append(msgJSON, '\n')
	c.TGSConn.Write(msgJSON)
}

func (c *Client) ResponseTGS() {
	fmt.Println("+++++ Response TGS... +++++")
	buff := make([]byte, 1024)
	bytes, err := c.TGSConn.Read(buff)
	common.HandleError(err, "Client-115: ")
	fmt.Println("----- Body response encrypted", string(common.PrettyPrint(buff[:bytes-1])))
	var resp common.TGSResponseEncrypted
	err = json.Unmarshal(buff[:bytes-1], &resp)
	common.HandleError(err, "Client-118: ")
	fmt.Println("----- Body reponse TGSCliResp:", string(common.PrettyPrint([]byte(des.Decode(resp.TGSClientResponse, c.SessionKey)))))
	var cliResp common.TGSClientResponse
	err = json.Unmarshal([]byte(des.Decode(resp.TGSClientResponse, c.SessionKey)), &cliResp)
	common.HandleError(err, "Client-121: ")
	c.GeneratedKey = cliResp.GeneratedKey
	c.TGSServerResponseEncrypted = resp.TGSServerDataResponse
}

func (c *Client) RequestServer() {
	fmt.Println("+++++ Request domain server... +++++")
	cliReq := common.ServerClientRequest{
		ID: c.ID,
		TS: c.TimeStamp,
	}
	cliReqJSON, err := json.Marshal(cliReq)
	common.HandleError(err, "Client-132: ")
	fmt.Println("----- Body request cliReq:", string(common.PrettyPrint(cliReqJSON)))
	cliReqEncrypted := des.Encode(string(cliReqJSON), c.GeneratedKey)
	msg := common.ServerRequestEncrypted{
		ServerClientRequest: cliReqEncrypted,
		TGSServerResponse:   c.TGSServerResponseEncrypted,
	}
	msgJSON, err := json.Marshal(msg)
	common.HandleError(err, "Client-139: ")
	fmt.Println("----- Body request encrypted:", string(common.PrettyPrint(msgJSON)))
	msgJSON = append(msgJSON, '\n')
	c.ServerConn.Write(msgJSON)
}

func (c *Client) ResponseServer() {
	fmt.Println("+++++ Response domain server... +++++")
	buff := make([]byte, 1024)
	bytes, err := c.ServerConn.Read(buff)
	common.HandleError(err, "Client-147: ")
	fmt.Println("----- Body response encrypted:", string(common.PrettyPrint(buff[:bytes-1])))
	var resp common.ServerResponseEncrypted
	err = json.Unmarshal(buff[:bytes-1], &resp)
	common.HandleError(err, "Client-150: ")
	fmt.Println("----- Body response cliResp:", string(common.PrettyPrint([]byte(des.Decode(resp.ServerResponse, c.GeneratedKey)))))
	var cliResp common.ServerResponse
	err = json.Unmarshal([]byte(des.Decode(resp.ServerResponse, c.GeneratedKey)), &cliResp)
	common.HandleError(err, "Client-153: ")
	if cliResp.TS != c.TimeStamp {
		logrus.Fatal("Unequal ts of cliResp and client")
	} else {
		logrus.Println("User authenticated domain server")
	}
}
