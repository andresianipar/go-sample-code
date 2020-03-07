package gosamplecode

// F12 function
func F12(flip string) (int, int) {
	var heads, tails int

	switch flip {
	case "heads":
		heads++
	case "tails":
		tails++
	default:
	}

	return heads, tails
}

// F13 function
func F13(x int) int {
	switch {
	case x > 0:
		return +1
	default:
		return 0
	case x < 0:
		return -1
	}
}
