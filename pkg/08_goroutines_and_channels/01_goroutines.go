package gosamplecode

import (
	"fmt"
	"time"

	library "github.com/andresianipar/go-sample-code/internal"
)

// F1 function
func F1() int {
	go library.Spinner(100 * time.Millisecond)

	const n = 45
	fibN := library.Fib(n) // slow

	fmt.Printf("\rFibonacci(%d) = %d\n", n, fibN)

	return fibN
}
