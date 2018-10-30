/*
 * pdsim.go
 */

package main

import (
    "flag"
    "fmt"
    "log"
    "github.com/mdunthorn/pdsim/proto/eis"
    "net"
    "os"
)

var start_port int
var num_servers int
func init() {
    flag.IntVar(&start_port, "start_port", 2001, "start port")
    flag.IntVar(&num_servers, "num_servers", 1, "number of servers")
}

func start_listener(port int) {
    lstr := fmt.Sprintf(":%d", port)
    l, err := net.Listen("tcp", lstr)
	if err != nil {
		log.Fatal(err)
	} else {
        log.Printf("listening on port %d...\n", port)
    }
	defer l.Close()
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}
        go handler.Handle(conn)
	}
}

func main() {
    flag.Parse()
    log.Print("start program")
    log.Printf("start_port: %d, num_servers: %d", start_port, num_servers)
    for i := 0; i < num_servers; i++ {
        port := start_port + i
        go start_listener(port)
    }
    port := start_port + num_servers
    start_listener(port)
}
