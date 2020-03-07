package gosamplecode

import (
	"log"
	"net/http"

	library "github.com/andresianipar/go-sample-code/internal"
)

// F9 function
func F9() {
	http.HandleFunc("/", library.Handler1)

	log.Fatal(http.ListenAndServe("localhost:3000", nil))
}

// F10 function
func F10() {
	http.HandleFunc("/", library.Handler2)
	http.HandleFunc("/count", library.Counter)

	log.Fatal(http.ListenAndServe("localhost:3000", nil))
}

// F11 function
func F11() {
	http.HandleFunc("/", library.Handler3)

	log.Fatal(http.ListenAndServe("localhost:3000", nil))
}
