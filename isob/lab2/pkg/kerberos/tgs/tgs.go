package tgs

import (
	"encoding/json"
	"fmt"
	"net"

	"github.com/sirupsen/logrus"

	"github.com/Muha113/6-th-term-labs/isob/lab2/pkg/des"
	"github.com/Muha113/6-th-term-labs/isob/lab2/pkg/kerberos/common"
)

type TGS struct {
	TGSKey     string
	SessionKey string
	Listener   net.Listener
}

func (t *TGS) HandleClientRequest() {
	fmt.Println("+++++ Handling client request... +++++")
	conn, err := t.Listener.Accept()
	common.HandleError(err, "TGS-20: ")
	defer conn.Close()
	buff := make([]byte, 1024)
	bytes, err := conn.Read(buff)
	common.HandleError(err, "TGS-24: ")
	fmt.Println("----- Body client request encrypted:", string(common.PrettyPrint(buff[:bytes-1])))
	var req common.TGSRequestEncrypted
	err = json.Unmarshal(buff[:bytes-1], &req)
	common.HandleError(err, "TGS-27: ")
	fmt.Println("----- Ticket granting ticket:", string(common.PrettyPrint([]byte(des.Decode(req.TGT, t.TGSKey)))))
	var tgt common.TicketGrantingTicket
	err = json.Unmarshal([]byte(des.Decode(req.TGT, t.TGSKey)), &tgt)
	common.HandleError(err, "TGS-30: ")
	fmt.Println("----- Body TGS client request:", string(common.PrettyPrint([]byte(des.Decode(req.TGSClientRequest, tgt.SessionKey)))))
	var cliReq common.TGSClientRequest
	err = json.Unmarshal([]byte(des.Decode(req.TGSClientRequest, tgt.SessionKey)), &cliReq)
	common.HandleError(err, "TGS-33: ")
	if cliReq.ID != tgt.ID {
		logrus.Fatal("Unequal tgt and client requests")
	}
	t.SessionKey = tgt.SessionKey

	dialer, err := net.Dial("tcp", "127.0.0.1:8000")
	common.HandleError(err, "TGS-40: ")
	dialer.Write([]byte(cliReq.Req + "\n"))
	bytes, err = dialer.Read(buff)
	common.HandleError(err, "TGS-43: ")
	dialer.Close()
	fmt.Println("+++++ Build client response... +++++")
	serverKey := string(buff[:bytes-1])
	generatedKey := common.GenKey(8)
	cliResp := common.TGSClientResponse{
		GeneratedKey: generatedKey,
		Dest:         cliReq.Req,
		TS:           cliReq.TS,
		Exp:          tgt.Exp,
	}
	servData := common.TGSServerDataResponse{
		ID:           cliReq.ID,
		GeneratedKey: generatedKey,
		Addr:         tgt.Addr,
		TS:           tgt.TS,
		Exp:          tgt.Exp,
	}
	servDataJSON, err := json.Marshal(servData)
	common.HandleError(err, "TGS-61: ")
	fmt.Println("----- Body servData client response:", string(common.PrettyPrint(servDataJSON)))
	cliRespJSON, err := json.Marshal(cliResp)
	common.HandleError(err, "TGS-63: ")
	fmt.Println("----- Body cliResp client response:", string(common.PrettyPrint(cliRespJSON)))
	servDataEncrypted := des.Encode(string(servDataJSON), serverKey)
	cliRespEncrypted := des.Encode(string(cliRespJSON), tgt.SessionKey)
	msg := common.TGSResponseEncrypted{
		TGSClientResponse:     cliRespEncrypted,
		TGSServerDataResponse: servDataEncrypted,
	}
	msgJSON, err := json.Marshal(msg)
	common.HandleError(err, "TGS-71: ")
	fmt.Println("----- Body client response encrypted:", string(common.PrettyPrint(msgJSON)))
	msgJSON = append(msgJSON, '\n')
	conn.Write(msgJSON)
}
