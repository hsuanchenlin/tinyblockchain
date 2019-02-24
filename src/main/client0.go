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
	args := &Args{ From: "abc", To: "ggg", Value: 99}
	var reply string
	c := jsonrpc.NewClient(client)
	err = c.Call("Dealer.AcceptTX", args, &reply)
	if err != nil {
		log.Fatal("rpc error:", err)
	}
	fmt.Printf("Result0: %s\n",reply)

	args = &Args{ From: "rrr", To: "ttt", Value: 33}
	//c = jsonrpc.NewClient(client)
	err = c.Call("Dealer.AcceptTX", args, &reply)
	if err != nil {
		log.Fatal("rpc error:", err)
	}
	fmt.Printf("Result1: %s\n",reply)

}

