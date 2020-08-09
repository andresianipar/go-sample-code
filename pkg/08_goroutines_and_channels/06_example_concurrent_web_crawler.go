package gosamplecode

import (
	"fmt"

	library "github.com/andresianipar/go-sample-code/internal"
)

// F8 function
func F8() chan []string {
	worklist := make(chan []string) // lists of URLs, may have duplicates

	// library.BreadthFirst3(library.Crawl3, worklist)
	library.BreadthFirst4(library.Crawl3, worklist)
	go func() {
		for url := range worklist {
			fmt.Printf("%s\n", url)
		}
	}()

	return worklist
}
