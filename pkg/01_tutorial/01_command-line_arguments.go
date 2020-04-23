package gosamplecode

import (
	"os"
	"strings"
)

// F1 function
func F1() string {
	var s string

	for i := 1; i < len(os.Args); i++ {
		s += os.Args[i] + " "
	}

	return strings.TrimRight(s, " ")
}

// F2 function
func F2() string {
	fruits := [5]string{"Apple", "Pear", "Plum", "Tomato", "Peach"}
	s := ""

	for _, fruit := range fruits {
		s += fruit + ", "
	}

	s = strings.TrimRight(s, ", ")

	return s
}

// F3 function
func F3() string {
	fruits := []string{"Apple", "Pear", "Plum", "Tomato", "Peach"}

	return strings.Join(fruits, ", ")
}
