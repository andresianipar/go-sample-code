package gosamplecode

import (
	"flag"
	"fmt"
	"sync"
	"time"

	library "github.com/andresianipar/go-sample-code/internal"
)

// F12 function
func F12() (int64, int64) {
	// Determine the initial directories.
	flag.Parse()

	roots := flag.Args()

	if len(roots) == 0 {
		roots = []string{"."}
	}

	// Traverse the file tree.
	fileSizes := make(chan int64)

	go func() {
		for _, root := range roots {
			library.WalkDir1(root, fileSizes)
		}
		close(fileSizes)
	}()

	// Print the results.
	var nFiles, nBytes int64

	for size := range fileSizes {
		nFiles++
		nBytes += size
	}
	fmt.Printf("%d files %.1f MB\n", nFiles, float64(nBytes)/1e6)

	return nFiles, nBytes
}

var verbose = flag.Bool("v", false, "show verbose progress messages")

// F13 function
func F13() (int64, int64) {
	// Determine the initial directories.
	flag.Parse()

	roots := flag.Args()

	if len(roots) == 0 {
		roots = []string{"."}
	}

	// Traverse the file tree.
	fileSizes := make(chan int64)

	go func() {
		for _, root := range roots {
			library.WalkDir1(root, fileSizes)
		}
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

// F14 function
func F14() (int64, int64) {
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
		go library.WalkDir2(root, &n, fileSizes)
	}
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
