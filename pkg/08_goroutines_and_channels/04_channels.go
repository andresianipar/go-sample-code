package gosamplecode

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"time"

	library "github.com/andresianipar/go-sample-code/internal"
)

// F4 function
func F4() {
	conn, err := net.Dial("tcp", "localhost:8000")

	if err != nil {
		log.Fatal(err)
	}
	// defer conn.Close()
	// library.MustCopy(os.Stdout, conn)
	// go library.MustCopy(os.Stdout, conn)
	// done := make(chan struct{})
	done := make(chan bool)
	go func() {
		io.Copy(os.Stdout, conn) // NOTE: ignoring errors
		log.Println("Done")
		// done <- struct{}{} // signal the main goroutine
		done <- true // signal the main goroutine
	}()
	library.MustCopy(conn, os.Stdin)
	conn.Close()
	<-done // wait for background goroutine to finish
}

// F5 function
func F5() int {
	naturals := make(chan int)
	squares := make(chan int)

	// Counter
	go func() {
		for x := 0; x <= 4; x++ {
			naturals <- x
			time.Sleep(1 * time.Second)
		}
		close(naturals)
	}()

	// Squarer
	go func() {
		for {
			x, ok := <-naturals

			if !ok {
				break // channel was closed and drained
			}
			squares <- x * x
		}
		close(squares)
	}()

	// Printer (in main goroutine)
	// for {
	// 	x, ok := <-squares

	// 	if !ok {
	// 		break // channel was closed and drained
	// 	}
	// 	fmt.Println(x)
	// }

	var x int

	for x = range squares {
		fmt.Println(x)
	}

	return x
}

// F6 function
func F6() int {
	naturals := make(chan int)
	squares := make(chan int)

	go library.ProduceNumbers(naturals)
	go library.Squarer(squares, naturals)
	return library.Printer(squares)
}

// F7 function
func F7() (int, int) {
	ch := make(chan string, 3)

	ch <- "A"
	ch <- "B"
	ch <- "C"
	// ch <- "D"         // block
	fmt.Println(<-ch) // <-["A" "B" "C"]<-
	ch <- "D"
	fmt.Println(<-ch) // <-["B" "C" "D"]<-
	ch <- "E"
	fmt.Println(<-ch) // <-["C" "D" "E"]<-
	fmt.Println(<-ch) // <-["D" "E"]<-
	fmt.Println(<-ch) // <-["E"]<-

	c := cap(ch)
	l := len(ch)

	fmt.Println(c)
	fmt.Println(l)

	return c, l
}
