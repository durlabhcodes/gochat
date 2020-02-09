package client

import (
	"bufio"
	"fmt"
	"net"
	"os"

	"github.com/sirupsen/logrus"
)

//ClientInit method
func ClientInit() {
	logrus.Info("Initiating Client")

	conn, _ := net.Dial("tcp", "127.0.0.1:8585")
	for {
		reader := bufio.NewReader(os.Stdin)
		line, _ := reader.ReadString('\n')
		fmt.Fprintf(conn, "%s", line)
	}
}
