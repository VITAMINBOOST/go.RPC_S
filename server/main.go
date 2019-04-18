package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"

	tasks "go.RPC_S/tasks"
)

func main() {
	task := new(tasks.Task)
	err := rpc.Register(task)

	// rsrc := new(tasks.Rsrc)
	// err := rpc.Register(rsrc)

	if err != nil {
		log.Fatal("Format of service Task isn't correct.", err)
	}

	rpc.HandleHTTP()

	listener, e := net.Listen("tcp", ":1234")

	if e != nil {
		log.Fatal("Listen error: ", e)
	}

	log.Printf("Serving RPC server on port %d", 1234)

	err = http.Serve(listener, nil)

	if err != nil {
		log.Fatal("Error serving: ", err)
	}
}
