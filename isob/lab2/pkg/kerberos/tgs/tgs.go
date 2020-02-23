package tgs

import (
	"encoding/json"
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
	conn, err := t.Listener.Accept()
	common.HandleError(err, "TGS-20: ")
	defer conn.Close()
	buff := make([]byte, 1024)
	bytes, err := conn.Read(buff)
	common.HandleError(err, "TGS-24: ")
	var req common.TGSRequestEncrypted
	err = json.Unmarshal(buff[:bytes-1], &req)
	common.HandleError(err, "TGS-27: ")
	var tgt common.TicketGrantingTicket
	err = json.Unmarshal([]byte(des.Decode(req.TGT, t.TGSKey)), &tgt)
	common.HandleError(err, "TGS-30: ")
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
	serverKey := string(buff[:bytes-1])
	generatedKey := common.GenKey(32)
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
	cliRespJSON, err := json.Marshal(cliResp)
	common.HandleError(err, "TGS-63: ")
	servDataEncrypted := des.Encode(string(servDataJSON), serverKey)
	cliRespEncrypted := des.Encode(string(cliRespJSON), tgt.SessionKey)
	msg := common.TGSResponseEncrypted{
		TGSClientResponse:     cliRespEncrypted,
		TGSServerDataResponse: servDataEncrypted,
	}
	msgJSON, err := json.Marshal(msg)
	common.HandleError(err, "TGS-71: ")
	msgJSON = append(msgJSON, '\n')
	conn.Write(msgJSON)
}
