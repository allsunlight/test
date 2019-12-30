package main

import (
	"github/test/client/client"
	"github/test/server"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"time"
)

func main()  {
	arith := new(server.Arith)
	rpc.Register(arith)
	rpc.HandleHTTP()
	l, e := net.Listen("tcp", ":1234")
	if e != nil {
		log.Fatal("listen error:", e)
	}
	go http.Serve(l, nil)

	time.Sleep(time.Hour)
	client.Ec()
}
