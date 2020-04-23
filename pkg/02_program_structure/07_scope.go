package gosamplecode

// F24 function
func F24() string {
	var s string
	x := "Hello!"

	for i := 0; i < len(x); i++ {
		x := x[i]

		if '!' != x {
			x := x

			s += string(x)

			// fmt.Printf("x = %c, s = %s\n", x, s)
		}
	}

	// fmt.Println("")

	return s
}

// F25 function
func F25() {
	if x := F24(); "" == x {
		// fmt.Println(x)
	} else if y := F24(); x == y {
		// fmt.Println(x, y)
	} else {
		// fmt.Println(x, y)
	}
}
