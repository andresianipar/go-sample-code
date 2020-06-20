package gosamplecode

import (
	"log"
	"net/http"

	library "github.com/andresianipar/go-sample-code/internal"
)

// F11 function
func F11() {
	db := library.Database{
		"shoes": 50,
		"socks": 5,
	}

	log.Fatal(http.ListenAndServe("localhost:3000", db))
}

// F12 function
func F12() {
	db := library.Database{
		"shoes": 50,
		"socks": 5,
	}
	mux := http.NewServeMux()

	mux.Handle("/list", http.HandlerFunc(db.List))
	mux.Handle("/price", http.HandlerFunc(db.Price))

	log.Fatal(http.ListenAndServe("localhost:3000", db))
}

// F13 function
func F13() {
	db := library.Database{
		"shoes": 50,
		"socks": 5,
	}

	http.HandleFunc("/list", db.List)
	http.HandleFunc("/price", db.Price)

	log.Fatal(http.ListenAndServe("localhost:3000", nil))
}
