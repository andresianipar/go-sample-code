package gosamplecode

import (
	"log"
	"net"

	library "github.com/andresianipar/go-sample-code/internal"
)

// F2 function
// F2 is a TCP server that periodically writes the time.
func F2() {
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
		// library.HandleConn(conn) // handle one connection at a time
		go library.HandleConn1(conn) // handle connections concurrently
	}
}
