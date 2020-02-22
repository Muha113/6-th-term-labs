package as

import (
	"encoding/json"
	"net"
	"time"

	"github.com/Muha113/6-th-term-labs/isob/lab2/pkg/des"
	"github.com/Muha113/6-th-term-labs/isob/lab2/pkg/kerberos/common"
)

type AS struct {
	TGSKey         string
	ClientsKeysMap map[string]string
	ServersKeysMap map[string]string
	ExpTime        time.Duration
	Listener       net.Listener
}

//supposed to user is registered
func (a *AS) HandleClientRequest() {
	conn, err := a.Listener.Accept()
	common.HandleError(err)
	buff := make([]byte, 1024)
	bytes, err := conn.Read(buff)
	common.HandleError(err)
	var req *common.ASClientRequest
	err = json.Unmarshal(buff[:bytes-1], req)
	common.HandleError(err)
	tmp := req.TS
	timeStamp := des.Decode(tmp, a.ClientsKeysMap[req.ID])
	timeParsed, err := time.Parse("2006-01-02 15:04:05", timeStamp)
	common.HandleError(err)
	req.TS = timeParsed.String()
	sessionKey := common.GenKey(32)
	addr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:8001")
	common.HandleError(err)
	resp := common.ASClientResponse{
		SessionKey: sessionKey,
		TS:         req.TS,
		TGSAddr:    addr,
		Exp:        timeParsed.Add(a.ExpTime).String(),
	}
	tgt := common.TicketGrantingTicket{
		ID:         req.ID,
		SessionKey: sessionKey,
		TS:         req.TS,
		Addr:       req.Addr,
		Exp:        resp.Exp,
	}
	respJSON, err := json.Marshal(resp)
	common.HandleError(err)
	tgtJSON, err := json.Marshal(tgt)
	common.HandleError(err)
	respEncrypted := des.Encode(string(respJSON), a.ClientsKeysMap[req.ID])
	tgtEncrypted := des.Encode(string(tgtJSON), a.TGSKey)
	msg := common.ASResponseEncrypted{
		ASClientResponse: respEncrypted,
		TGT:              tgtEncrypted,
	}
	msgJSON, err := json.Marshal(msg)
	common.HandleError(err)
	msgJSON = append(msgJSON, '\n')
	conn.Write(msgJSON)
}

func (a *AS) HandleTGSRequest() {

}
