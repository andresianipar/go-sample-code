package gosamplecode

// F2 function
func F2() interface{} {
	var any interface{}

	any = true
	any = 12.34
	any = "Hello, World!"
	any = map[string]int{"one": 1}

	// fmt.Println(any)

	return any
}
