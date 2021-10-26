package main

import (
	"github.com/olamiko/key-value-store/utils"
	"log"
	"net"
	"net/http"
	"net/rpc"
)

type Listener int

type Reply struct {
	response string
}

var store utils.Store = utils.NewStore("storage.kv")

func (l *Listener) getListener(key string, reply *Reply) error {
	*reply = Reply{store.Get(key)}
	return nil
}

func (l *Listener) setListener(key_val []string, reply *Reply) error {
	key := key_val[0]
	value := key_val[1]
	*reply = Reply{store.Set(key, value)}
	return nil
}

func startServer() {
	listener := new(Listener)
	rpc.Register(listener)
	rpc.HandleHTTP()

	l, err := net.Listen("tcp", "30800")
	if err != nil {
		log.Fatal("listen error: ", err)
	}

	go http.Serve(l, nil)
}

func main() {
	startServer()
}
