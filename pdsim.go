/*
 * pdsim.go
 */

package main

import (
    "fmt"
    "log"
    "github.com/mdunthorn/pdsim/proto/eis"
    "net"
)

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
    log.Print("start program")
    go start_listener(2001)
    go start_listener(2002)
    start_listener(2003)
}
