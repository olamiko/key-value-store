package main

import (
	//	"github.com/olamiko/key-value-store/utils"
	"net"
)

var lbPort = ":30000"

func startLB() error {

	tcpAddr, err := net.ResolveTCPAddr("tcp4", lbPort)

	if err != nil {
		return err
	}

	//listen
	listener, err := net.ListenTCP("tcp", tcpAddr)

	if err != nil {
		return err
	}

	for {
		conn, err := listener.Accept()

		if err != nil {
			return err
		}

		//conn.Write([]byte("random"))
		//conn.Close()
		//fmt.Println("got here")

		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	defer conn.Close()

	conn.Write([]byte("random"))
}

func main() {
	startLB()
}
