package gosamplecode

import (
	"flag"
	"strings"
)

// F4 function
func F4() (bool, float64, string) {
	var b, f, s = true, 4.5, "yes"

	// fmt.Printf("b = %t, f = %f, s = %s\n", b, f, s)

	return b, f, s
}

// F5 function
func F5() *int {
	var x int
	p := &x

	*p = 1

	// fmt.Printf("x = %d\n", *p)

	return p
}

// F6 function
func F6() *int {
	v := 1

	// fmt.Println(&v)

	return &v
}

// F7 function
func F7(p *int) int {
	*p++

	// fmt.Println(*p)

	return *p
}

// F8 function
func F8() string {
	var n = flag.Bool("n", false, "Omit trailing newline")
	var s = flag.String("s", " ", "Separator")

	flag.Parse()

	args := strings.Join(flag.Args(), *s)

	// fmt.Println(args)

	if !*n {
		args = "No arguments"

		// fmt.Printf(args)
	}

	return args
}

// F9 function
func F9() *int {
	p := new(int)

	// fmt.Println(*p)

	*p++

	// fmt.Println(*p)

	return p
}

// F10 function
func F10() (*int, *int) {
	var dummy int

	return &dummy, new(int)
}

// F11 function
func F11(old, new int) int {
	r := new - old

	// fmt.Println(r)

	return r
}

// Global variable
var Global *int

// F12 function
func F12() {
	var x int

	x = 1

	Global = &x

	// fmt.Println(*Global)
}
