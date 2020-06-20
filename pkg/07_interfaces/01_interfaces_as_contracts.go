package gosamplecode

import (
	"fmt"

	library "github.com/andresianipar/go-sample-code/internal"
)

// F1 function
func F1() library.ByteCounter {
	var c library.ByteCounter

	c.Write([]byte("Hello"))

	// fmt.Println(c)

	c = 0

	var name = "Dolly"

	fmt.Fprintf(&c, "Hello, %s", name)
	// fmt.Println(c)

	return c
}
