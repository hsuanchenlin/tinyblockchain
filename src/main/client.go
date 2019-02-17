package main

import (
	. "amis_test/src/client"
	"fmt"
	"log"
	"net"
	"net/rpc/jsonrpc"
)

func main() {

	client, err := net.Dial("tcp", "127.0.0.1:1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}
	// Synchronous call
	args := &Args{ From: "abc", To: "def", Value: 10}
	var reply int
	c := jsonrpc.NewClient(client)
	err = c.Call("Calculator.Add", args, &reply)
	if err != nil {
		log.Fatal("arith error:", err)
	}
	fmt.Printf("Result: %d+%d=%d\n", args.X, args.Y, reply)
}
