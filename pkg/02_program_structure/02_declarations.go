package gosamplecode

const boilingF = 212.0

// F2 function
func F2() (float64, float64) {
	var f = boilingF
	var c = (f - 32) * 5 / 9

	// fmt.Printf("Boiling point = %g°F or %g°C\n", f, c)

	return f, c
}

// F3 function
func F3(f float64) float64 {
	var c = (f - 32) * 5 / 9

	// fmt.Printf("%g°C\n", c)

	return c
}
