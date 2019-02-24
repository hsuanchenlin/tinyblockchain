// server.go
package main

import (
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	. "amis_test/src/server"
)

func main() {

	//accept tx loop
	//cal := new(Calculator)
	accs := new(Dealer)
	server := rpc.NewServer()
	server.Register(accs)
	server.HandleHTTP(rpc.DefaultRPCPath, rpc.DefaultDebugPath)
	listener, e := net.Listen("tcp", ":1234")
	if e != nil {
		log.Fatal("listen error:", e)
	}
	for {
		if conn, err := listener.Accept(); err != nil {
			log.Fatal("accept error: " + err.Error())
		} else {
			log.Printf("new connection established\n")
			go server.ServeCodec(jsonrpc.NewServerCodec(conn))
		}
	}
}



