package gosamplecode

import (
	"fmt"

	library "github.com/andresianipar/go-sample-code/internal"
)

// F16 function
func F16() {
	s := "Joker"

	switch s {
	case "Spades":
	case "Hearts":
	case "Diamonds":
	case "Clubs":
	default:
		panic(fmt.Sprintf("invalid suit %q", s))
	}
}

// F17 function
func F17() {
	library.F(4)
}

// F18 function
func F18() {
	defer library.PrintStack()

	library.F(4)
}
