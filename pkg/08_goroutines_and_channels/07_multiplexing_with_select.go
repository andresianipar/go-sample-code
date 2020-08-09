package gosamplecode

import (
	"fmt"
	"os"
	"time"
)

// F9 function
func F9() {
	fmt.Println("Commencing countdown.")

	tick := time.Tick(1 * time.Second)

	for countdown := 10; countdown > 0; countdown-- {
		fmt.Println(countdown)
		<-tick
	}
	// launch()
}

// F10 function
func F10() {
	abort := make(chan struct{})

	go func() {
		os.Stdin.Read(make([]byte, 1)) // read a single byte
		abort <- struct{}{}
	}()
	fmt.Println("Commencing countdown. Press return to abort.")

	tick := time.Tick(1 * time.Second)

	for countdown := 10; countdown > 0; countdown-- {
		fmt.Println(countdown)
		select {
		case <-tick:
		// Do nothing.
		case <-abort:
			fmt.Println("Launch aborted!")

			return
		}
	}
	// launch()
}

// F11 function
func F11() {
	ch := make(chan int, 1)

	for i := 0; i < 10; i++ {
		select {
		case x := <-ch:
			fmt.Println(x) // "0" "2" "4" "6" "8"
		case ch <- i:
		}
	}
}
