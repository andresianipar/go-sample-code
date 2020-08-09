package gosamplecode

import (
	"log"
	"net"

	library "github.com/andresianipar/go-sample-code/internal"
)

// F3 function
func F3() {
	listener, err := net.Listen("tcp", "localhost:8000")

	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()

		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		go library.HandleConn2(conn) // handle connections concurrently
	}
}
