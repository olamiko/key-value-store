package main

import (
	"bufio"
	"fmt"
	//"github.com/olamiko/key-value-store/loadbalancer"
	//	"io/ioutil"
	"log"
	"net"
	"os"
	//"strconv"
	//"strings"
)

var scanner = bufio.NewScanner(os.Stdin)

var lbPort string = ":30000"

func clientConn(input string) string {
	tcpAddr, err := net.ResolveTCPAddr("tcp", lbPort)

	if err != nil {
		log.Fatal(err)
	}

	conn, err := net.DialTCP("tcp", nil, tcpAddr)

	if err != nil {
		log.Fatal(err)
	}

	_, err = conn.Write([]byte(input))

	buffer := make([]byte, 1024)
	_, err = conn.Read(buffer)

	if err != nil {
		log.Fatal(err)
	}

	return string(buffer)
}

func runClient() {

	//REPL
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("> Welcome to the kv store interface \n")
	fmt.Print("> Allowed operations are: \n")
	fmt.Print("> set foo=bar (no spaces between key=value!) \n")
	fmt.Print("> get foo \n")
	fmt.Println(" ")

	for {
		fmt.Print("> ")
		scanner.Scan()
		if scanner.Err() != nil {
			fmt.Println("> " + scanner.Err().Error()) // Handle error.
		}

		response := clientConn(scanner.Text())
		fmt.Print("> " + response + "\n")
	}

}

func main() {
	for {
		runClient()
	}
}
