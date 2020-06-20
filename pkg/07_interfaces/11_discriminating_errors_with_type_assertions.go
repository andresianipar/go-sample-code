package gosamplecode

import (
	"os"

	library "github.com/andresianipar/go-sample-code/internal"
)

// F15 function
func F15() (bool, bool) {
	_, err := os.Stat("/no/such/file")
	ok1 := os.IsNotExist(err)

	if ok1 {
		// fmt.Println((err))
		// fmt.Printf("%#v\n", err)
	} else {
		// fmt.Println(fileInfo)
	}

	ok2 := library.IsNotExist(err)

	// fmt.Println(ok2)

	return ok1, ok2
}
