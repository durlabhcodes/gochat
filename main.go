package main

import (
	"flag"

	"github.com/layzbot/gochat/client"
	"github.com/layzbot/gochat/server"
	"github.com/sirupsen/logrus"
)

func main() {

	mode := flag.Int("mode", 1, "Define the start mode. 0 means server and 1 means client")
	flag.Parse()

	switch *mode {
	case 0:
		logrus.Info("Initiating Server")
		server.ServerConfig()
		break
	case 1:
		logrus.Info("Initiating Client")
		client.ClientConfig()
		break
	}
}
