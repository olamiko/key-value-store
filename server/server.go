package main

import (
	"github.com/olamiko/key-value-store/utils"
	"log"
	"net"
	//	"net/http"
	"net/rpc"
	//"strconv"
)

type Listener int

type Reply struct {
	Response string
}

var store utils.Store // = utils.NewStore("storage.kv")

func (l *Listener) GetListener(key string, reply *Reply) error {
	*reply = Reply{store.Get(key)}
	return nil
}

func (l *Listener) SetListener(key_val []string, reply *Reply) error {
	key := key_val[0]
	value := key_val[1]
	store.Set(key, value)
	*reply = Reply{"Added key successfully"}
	return nil
}

func StartServer(id string, storage string) {

	store = utils.NewStore(storage)

	addy, err := net.ResolveTCPAddr("tcp", "0.0.0.0:3080"+id)
	if err != nil {
		log.Fatal(err)
	}
	inbound, err := net.ListenTCP("tcp", addy)
	if err != nil {
		log.Fatal(err)
	}
	listener := new(Listener)
	rpc.Register(listener)
	//rpc.HandleHTTP()
	rpc.Accept(inbound)

	//l, err := net.Listen("tcp", ":30800")
	//if err != nil {
	//	log.Fatal("listen error: ", err)
	//}

	//go http.Serve(l, nil)
}

//func main() {
//	startServer()
//}
