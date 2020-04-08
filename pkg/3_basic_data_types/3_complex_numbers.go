package gosamplecode

// F6 function
func F6() (float64, float64) {
	var x complex128 = complex(1, 2)
	var y complex128 = complex(3, 4)
	var r = real(x * y)
	var i = imag(x * y)

	// fmt.Println(x * y)
	// fmt.Println(r)
	// fmt.Println(i)

	return r, i
}
