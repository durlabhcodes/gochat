package server

import (
	"bufio"
	"fmt"
	"net"
)

//ServerInit func
func ServerInit() {
	listener, _ := net.Listen("tcp", ":8585")
	conn, _ := listener.Accept()

	for {
		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Println("Message : ", message)
	}
}
