package gosamplecode

import library "github.com/andresianipar/go-sample-code/internal"

// F21 function
func F21(c library.Celsius) library.Kelvin {
	var k library.Kelvin = library.Kelvin(c + 273.15)

	// fmt.Println(k)

	return k
}

// F22 function
func F22(k library.Kelvin) library.Celsius {
	var c library.Celsius = library.Celsius(k - 273.15)

	// fmt.Println(c)

	return c
}

// F23 function
func F23() int {
	// fmt.Println(library.I)

	return library.I
}
