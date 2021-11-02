package loadbalancer

import (
	"github.com/olamiko/key-value-store/utils"
	"log"
	"net"
	"net/rpc"
	"strconv"
)

var replicaSlice []string

var quit chan struct{}

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

func getIpAddress(id int) string {
	if id < 10 {
		return "0.0.0.0:3080" + strconv.Itoa(id)
	} else {
		return "0.0.0.0:308" + strconv.Itoa(id)
	}
}

func startServer(id int, storage string) {

	store = utils.NewStore(storage)
	ipAddress := getIpAddress(id)

	address, err := net.ResolveTCPAddr("tcp", ipAddress)
	if err != nil {
		log.Fatal(err)
	}
	inbound, err := net.ListenTCP("tcp", address)
	if err != nil {
		log.Fatal(err)
	}
	listener := new(Listener)
	rpc.Register(listener)
	rpc.Accept(inbound)

}

func StartServers(replicas int) {
	id := replicas

	quit := make(chan struct{})

	for id > 0 {
		storage := "storage-" + strconv.Itoa(id) + ".kv"

		go func() {
			for {
				select {
				case <-quit:
					return
				default:
					startServer(id, storage)
				}
			}
		}()
		replicaSlice = append(replicaSlice, getIpAddress(id))
		id = -1
	}
}
