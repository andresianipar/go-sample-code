package gosamplecode

// F1 function
func F1() (uint64, uint64) {
	var maxInt uint64 = 1<<64 - 1

	// fmt.Printf("Type: %T Value: %v\n", maxInt, maxInt)

	var remainder = maxInt % 5

	// fmt.Printf("%v\n", remainder)

	return maxInt, remainder
}

// F2 function
func F2() [3]string {
	var reversed [3]string
	medals := []string{"gold", "silver", "bronze"}

	for i := len(medals) - 1; i >= 0; i-- {
		// fmt.Println(medals[i])

		reversed[len(medals)-i-1] = medals[i]
	}

	return reversed
}

// F3 function
func F3() (int, int) {
	var apples int16 = 1
	var oranges int32 = 2
	var compote int = int(apples) + int(oranges)

	// fmt.Println(compote)

	f := 3.141
	i := int(f)

	// fmt.Println(f, i)

	return compote, i
}

// F4 function
func F4() (rune, rune) {
	ascii := 'a'
	newline := '\n'

	// fmt.Printf("%d %[1]c %[1]q\n", ascii)
	// fmt.Printf("%d %[1]q\n", newline)

	return ascii, newline
}
