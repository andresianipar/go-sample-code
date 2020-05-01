package gosamplecode

import (
	"os"
	"time"

	library "github.com/andresianipar/go-sample-code/internal"
)

// F12 function
func F12() []string {
	urls, err := library.BreadthFirst2(library.Crawl2, []string{"https://golang.org"})
	// urls, err := library.BreadthFirst2(library.Crawl2, []string{"https://golang.org/doc/gopher/frontpage.png"})

	if nil != err {
		// fmt.Fprintf(os.Stderr, "%v\n", err)

		os.Exit(1)
	}

	// for i, url := range urls {
	// 	fmt.Printf("%d:\t%s\n", i+1, url)
	// }

	return urls
}

// F13 function
func F13() {
	defer library.Trace("F13")()

	// simulate slow operation by sleeping
	time.Sleep(10 * time.Second)
}

// F14 function
func F14() (int, int) {
	r1 := library.Double(4)

	// fmt.Println(r1)

	r2 := library.Triple(4)

	// fmt.Println(r2)

	return r1, r2
}

// F15 function
func F15() (string, int64) {
	local, n, err := library.FetchBasePath("https://golang.org/doc/")

	if nil != err {
		// fmt.Fprintf(os.Stderr, "%v\n", err)

		os.Exit(1)
	}

	// fmt.Println(local)
	// fmt.Println(n)

	return local, n
}
