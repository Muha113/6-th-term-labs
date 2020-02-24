package main

import (
	"net"

	"github.com/sirupsen/logrus"

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
	logrus.Println("Starting TGS...")
	tgs.Listener, err = net.Listen("tcp", "127.0.0.1:8001")
	common.HandleError(err, "TGSServer-17: ")
	logrus.Printf("Listening on -> %s ...\n", tgs.Listener.Addr().String())
	tgs.HandleClientRequest()
	tgs.Listener.Close()
}
