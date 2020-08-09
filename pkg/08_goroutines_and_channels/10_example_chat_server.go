package gosamplecode

import (
	"log"
	"net"

	library "github.com/andresianipar/go-sample-code/internal"
)

// F16 function
func F16() {
	listener, err := net.Listen("tcp", "localhost:8000")

	if err != nil {
		log.Fatal(err)
	}
	go library.Broadcaster()
	for {
		conn, err := listener.Accept()

		if err != nil {
			log.Print(err)
			continue
		}
		go library.HandleConn3(conn)
	}
}
