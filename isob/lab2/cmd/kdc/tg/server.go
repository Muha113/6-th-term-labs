package main

import (
	"net"

	"github.com/Muha113/6-th-term-labs/isob/lab2/pkg/kerberos/common"
	"github.com/Muha113/6-th-term-labs/isob/lab2/pkg/kerberos/tgs"
)

func main() {
	tgs := tgs.TGS{
		TGSKey:     "tgsMasterKey",
		SessionKey: "",
		Listener:   nil,
	}
	var err error
	tgs.Listener, err = net.Listen("tcp", "127.0.0.1:8001")
	common.HandleError(err, "TGSServer-17: ")
	tgs.HandleClientRequest()
	tgs.Listener.Close()
}
