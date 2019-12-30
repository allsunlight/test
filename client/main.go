package main

import (
	"fmt"
	"github/test/server"
	"log"
	"net/rpc"
)

func main(){
	client, err := rpc.DialHTTP("tcp",  "127.0.0.1:1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}
	args := &server.Args{7,8}
	var reply int
	err = client.Call("Arith.Multiply", args, &reply)
	if err != nil {
		log.Fatal("arith error:", err)
	}
	fmt.Printf("Arith: %d*%d=%d", args.A, args.B, reply)

	quotient := new(server.Quotient)
	divCall := client.Go("Arith.Divide", args, quotient, nil)
	replyCall := <-divCall.Done
	fmt.Println(replyCall)
	fmt.Println(quotient)
}
