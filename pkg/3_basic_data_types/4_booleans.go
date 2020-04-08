package gosamplecode

// F7 function
func F7() (bool, bool) {
	var x bool = true

	// fmt.Println(x)

	var c int = 55
	var r bool

	if 'a' <= c && 'z' >= c || 'A' <= c && 'Z' >= c || '0' <= c && '9' >= c {
		r = true
	}

	// fmt.Println(r)

	return !x == false, r
}

// F8 function
func F8(i int) bool {
	var r bool = i != 0

	// fmt.Println(r)

	return r
}
