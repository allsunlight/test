package main

import (
	"github.com/allsunlight/test/client/client"
	"github.com/allsunlight/test/server"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"time"
)
//require github.com/allsunlight/test v0.0.0-20191230103711-1fa4bf1b594d

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
