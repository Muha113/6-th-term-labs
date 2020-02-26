package main

import (
	"encoding/json"
	"fmt"
	"net"

	"github.com/sirupsen/logrus"

	"github.com/Muha113/6-th-term-labs/isob/lab2/pkg/des"
	"github.com/Muha113/6-th-term-labs/isob/lab2/pkg/kerberos/common"
)

type Server struct {
	ServerKey string
	Listener  net.Listener
}

func (s *Server) HandleClientRequest() {
	fmt.Println("+++++ Handling client request... +++++")
	conn, err := s.Listener.Accept()
	common.HandleError(err, "DomainServer-19: ")
	defer conn.Close()
	buff := make([]byte, 1024)
	bytes, err := conn.Read(buff)
	common.HandleError(err, "DomainServer-23: ")
	fmt.Println("----- Body client request encrypted:", string(common.PrettyPrint(buff[:bytes-1])))
	var req common.ServerRequestEncrypted
	err = json.Unmarshal(buff[:bytes-1], &req)
	common.HandleError(err, "DomainServer-26: ")
	fmt.Println("----- Body TGS serverData client request:", string(common.PrettyPrint([]byte(des.Decode(req.TGSServerResponse, s.ServerKey)))))
	var servReq common.TGSServerDataResponse
	err = json.Unmarshal([]byte(des.Decode(req.TGSServerResponse, s.ServerKey)), &servReq)
	common.HandleError(err, "DomainServer-34: ")
	fmt.Println("----- Body servCli client request:", string(common.PrettyPrint([]byte(des.Decode(req.ServerClientRequest, servReq.GeneratedKey)))))
	var cliReq common.ServerClientRequest
	err = json.Unmarshal([]byte(des.Decode(req.ServerClientRequest, servReq.GeneratedKey)), &cliReq)
	common.HandleError(err, "DomainServer-31: ")
	if cliReq.ID != servReq.ID || cliReq.TS != servReq.TS {
		logrus.Fatal("Unequal TGSResponse and ClientRequest to domain")
	} else {
		logrus.Println("User authenticated domain server")
	}
	fmt.Println("+++++ Build client response... +++++")
	resp := common.ServerResponse{
		Name: "TestServer",
		TS:   cliReq.TS,
	}
	respJSON, err := json.Marshal(resp)
	common.HandleError(err, "DomainServer-40: ")
	fmt.Println("----- Body client response:", string(common.PrettyPrint(respJSON)))
	respEncrypted := des.Encode(string(respJSON), servReq.GeneratedKey)
	msg := common.ServerResponseEncrypted{
		ServerResponse: respEncrypted,
	}
	msgJSON, err := json.Marshal(msg)
	common.HandleError(err, "DomainServer-46: ")
	fmt.Println("----- Body client response encrypted:", string(common.PrettyPrint(msgJSON)))
	msgJSON = append(msgJSON, '\n')
	conn.Write(msgJSON)
}

func main() {
	server := Server{
		ServerKey: "ServMKey",
		Listener:  nil,
	}
	var err error
	logrus.Println("Starting domain server...")
	server.Listener, err = net.Listen("tcp", "127.0.0.1:8002")
	common.HandleError(err, "DomainServer-58: ")
	logrus.Printf("Listening on -> %s ...\n", server.Listener.Addr().String())
	server.HandleClientRequest()
	server.Listener.Close()
}
