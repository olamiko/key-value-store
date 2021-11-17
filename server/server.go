package server

import (
	//"github.com/olamiko/key-value-store/utils"
	"net"
)

func startServer(port string) error {

	tcpAddr, err := net.ResolveTCPAddr("tcp4", port)

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

		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	defer conn.Close()

	conn.Write([]byte("random"))
}
