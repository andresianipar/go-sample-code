package gosamplecode

import (
	"flag"
	"fmt"
	"os"
	"sync"
	"time"

	library "github.com/andresianipar/go-sample-code/internal"
)

// F15 function
func F15() (int64, int64) {
	// Determine the initial directories.
	flag.Parse()

	roots := flag.Args()

	if len(roots) == 0 {
		roots = []string{"."}
	}

	// Traverse each root of the file tree in parallel.
	var n sync.WaitGroup
	fileSizes := make(chan int64)

	for _, root := range roots {
		n.Add(1)
		go library.WalkDir3(root, &n, fileSizes)
	}
	// Cancel traversal when input is detected.
	go func() {
		os.Stdin.Read(make([]byte, 1)) // read a single byte
		close(library.Done)
	}()
	go func() {
		n.Wait()
		close(fileSizes)
	}()

	// Print the results periodically.
	var tick <-chan time.Time

	if *verbose {
		tick = time.Tick(500 * time.Millisecond)
	}

	// Print the results.
	var nFiles, nBytes int64

loop:
	for {
		select {
		case <-library.Done:
			// Drain fileSizes to allow existing goroutines to finish.
			for range fileSizes {
				// Do nothing.
			}
			break loop
		case size, ok := <-fileSizes:
			if !ok {
				break loop // fileSizes was closed
			}
			nFiles++
			nBytes += size
		case <-tick:
			fmt.Printf("%d files %.1f MB\n", nFiles, float64(nBytes)/1e6)
		}
	}
	fmt.Printf("%d files %.1f MB\n", nFiles, float64(nBytes)/1e6) // final totals

	return nFiles, nBytes
}
