package gosamplecode

// F13 function
func F13() int {
	x := 10

	x /= 2

	// fmt.Printf("%d\n", x)

	return x
}

// F14 function
func F14() (int, int) {
	x, y := 1, 2

	y, x = x, y

	// fmt.Printf("x = %d, y = %d\n", x, y)

	return x, y
}

// F15 function
func F15(x, y int) int {
	// fmt.Printf("x = %d, y = %d\n", x, y)

	for 0 != y {
		x, y = y, x%y

		// fmt.Printf("x = %d, y = %d\n", x, y)
	}

	// fmt.Println(x)

	return x
}

// F16 function
func F16(n int) int {
	x, y := 0, 1

	for i := 0; i < n; i++ {
		x, y = y, x+y
	}

	// fmt.Println(x)

	return x
}
