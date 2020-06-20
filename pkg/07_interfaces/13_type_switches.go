package gosamplecode

import "fmt"

// F17 function
func F17() {
	var x interface{} = "Hello, World!"
	// var x interface{} = true

	switch x := x.(type) {
	case nil:
		fmt.Println("x is nil")
	case int, uint:
		fmt.Println("x is int or unint")
		fmt.Printf("x = %d\n", x)
	case string:
		fmt.Println("x is string")
		fmt.Printf("x = %s\n", x)
	default:
		fmt.Printf("unexpected type %T: %v\n", x, x)
	}
}
