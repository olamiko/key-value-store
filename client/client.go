package main

import (
	"fmt"
	"log"
	"net/rpc"
)

type Reply struct {
	Response string
}

var serverAddress string = "0.0.0.0"

func runClient() {
	client, err := rpc.Dial("tcp", serverAddress+":30800")

	if err != nil {
		log.Fatal(err)
	}

	// note that key_val is in a slice,

	var reply Reply
	err = client.Call("Listener.SetListener", []string{"name", "johnny"}, &reply)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Reply is %v\n", reply)

	// key is a string in this case
	err = client.Call("Listener.GetListener", "name", &reply)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("The key %v has value %v\n", "name", reply.Response)
}

func main() {
	runClient()
}
