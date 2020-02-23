package main

import (
	"encoding/json"
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
	conn, err := s.Listener.Accept()
	common.HandleError(err, "DomainServer-19: ")
	defer conn.Close()
	buff := make([]byte, 1024)
	bytes, err := conn.Read(buff)
	common.HandleError(err, "DomainServer-23: ")
	var req common.ServerRequestEncrypted
	err = json.Unmarshal(buff[:bytes-1], &req)
	common.HandleError(err, "DomainServer-26: ")
	var servReq common.TGSServerDataResponse
	err = json.Unmarshal([]byte(des.Decode(req.TGSServerResponse, s.ServerKey)), &servReq)
	var cliReq common.ServerClientRequest
	err = json.Unmarshal([]byte(des.Decode(req.ServerClientRequest, servReq.GeneratedKey)), &cliReq)
	common.HandleError(err, "DomainServer-31: ")
	if cliReq.ID != servReq.ID || cliReq.TS != servReq.TS {
		logrus.Fatal("Unequal TGSResponse and ClientRequest to domain")
	}
	resp := common.ServerResponse{
		Name: "TestServer",
		TS:   cliReq.TS,
	}
	respJSON, err := json.Marshal(resp)
	common.HandleError(err, "DomainServer-40: ")
	respEncrypted := des.Encode(string(respJSON), servReq.GeneratedKey)
	msg := common.ServerResponseEncrypted{
		ServerResponse: respEncrypted,
	}
	msgJSON, err := json.Marshal(msg)
	common.HandleError(err, "DomainServer-46: ")
	msgJSON = append(msgJSON, '\n')
	conn.Write(msgJSON)
}

func main() {
	server := Server{
		ServerKey: "serverMasterKey",
		Listener:  nil,
	}
	var err error
	server.Listener, err = net.Listen("tcp", "127.0.0.1:8002")
	common.HandleError(err, "DomainServer-58: ")
	server.HandleClientRequest()
	server.Listener.Close()
}
