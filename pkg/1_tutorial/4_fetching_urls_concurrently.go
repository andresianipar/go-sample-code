package gosamplecode

import (
	library "github.com/andresianipar/go-sample-code/internal"
)

// F8 function
func F8() chan string {
	// start := time.Now()
	ch := make(chan string)
	urls := []string{"https://golang.org", "http://gopl.io", "https://godoc.org"}

	for _, url := range urls {
		go library.Fetch(url, ch)
	}

	for range urls {
		// fmt.Println(<-ch)
	}

	// fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())

	return ch
}
