package gosamplecode

import (
	library "github.com/andresianipar/go-sample-code/internal"
)

// F17 function
func F17(c library.Celsius) library.Fahrenheit {
	var f library.Fahrenheit = library.Fahrenheit(c*9/5 + 32)

	// fmt.Println(f)

	return f
}

// F18 function
func F18(f library.Fahrenheit) library.Celsius {
	var c library.Celsius = library.Celsius((f - 32) * 5 / 9)

	// fmt.Println(c)

	return c
}

// F19 function
func F19() (library.Celsius, library.Fahrenheit, bool) {
	c := library.BoilingC - library.FreezingC
	boilingF := F17(library.BoilingC)
	freezingF := F17(library.FreezingC)
	f := boilingF - freezingF

	// fmt.Println(c)
	// fmt.Println(f)

	return c, f, c == library.Celsius(f)
}

// F20 function
func F20() string {
	var c library.Celsius = 100
	s := c.String()

	// fmt.Println(c)
	// fmt.Println(s)

	return s
}
