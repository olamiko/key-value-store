package main

import (
	"bufio"
	"fmt"
	"log"
	"net/rpc"
	"os"
	"strings"
)

type Reply struct {
	Response string
}

var serverAddress string = "0.0.0.0"

func parseInput(input string) string {

	var reply Reply
	client, err := rpc.Dial("tcp", serverAddress+":30800")
	if err != nil {
		log.Fatal(err)
	}

	splitSlice := strings.Split(input, " ")
	if len(splitSlice) > 1 {
		command := splitSlice[0]
		request := splitSlice[1]

		if strings.EqualFold(command, "get") {

			err = client.Call("Listener.GetListener", request, &reply)
			if err != nil {
				log.Fatal(err)
			}
		} else if strings.EqualFold(command, "set") {
			key_val := strings.Split(request, "=")
			if len(key_val) < 2 {
				return "set: incorrect input"
			}
			err = client.Call("Listener.SetListener", key_val, &reply)
			if err != nil {
				log.Fatal(err)
			}
		} else {
			return "Unknown command, try again."
		}

		return reply.Response
	}

	return "Incorrect command"

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

		response := parseInput(scanner.Text())
		fmt.Print("> " + response + "\n")
	}

}

func main() {
	for {
		runClient()
	}
}
