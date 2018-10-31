/*
 * pdsim.go
 */

package main

import (
    "flag"
    "fmt"
    "html"
    "log"
    // "github.com/gorilla/mux"
    "github.com/mdunthorn/pdsim/proto/eis"
    "net"
    "net/http"
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
    log.Print("start program")

    // Parse the command line
    flag.Parse()
    log.Printf("start_port: %d, num_servers: %d", start_port, num_servers)

    // Start some simulators
    for i := 0; i < num_servers; i++ {
        port := start_port + i
        go start_listener(port)
    }

    // Set up an api router
    http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
        log.Printf("api got a request for %q", html.EscapeString(r.URL.Path))
	    fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
    })
    // r := mux.NewRouter()
    // http:.Handle("/", r)

    // Start the api server
    log.Printf("starting api server on port %d", 8080)
    log.Fatal(http.ListenAndServe(":8080", nil))
}
