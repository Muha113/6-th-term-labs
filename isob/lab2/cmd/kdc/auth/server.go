package main

import (
	"net"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/Muha113/6-th-term-labs/isob/lab2/pkg/kerberos/common"

	"github.com/Muha113/6-th-term-labs/isob/lab2/pkg/kerberos/as"
)

func main() {
	as := as.AS{
		TGSKey:         "tgsMasterKey",
		ClientsKeysMap: make(map[string]string),
		ServersKeysMap: make(map[string]string),
		ExpTime:        time.Hour * 200,
		Listener:       nil,
	}
	var err error
	as.ClientsKeysMap["user1"] = common.Hash("user1PasswdKey")
	as.ServersKeysMap["TestServer"] = "serverMasterKey"
	logrus.Println("Starting AS...")
	as.Listener, err = net.Listen("tcp", "127.0.0.1:8000")
	common.HandleError(err, "ASServer-23: ")
	logrus.Printf("Listening on -> %s ...\n", as.Listener.Addr().String())
	as.HandleClientRequest()
	as.HandleTGSRequest()
	as.Listener.Close()
}
