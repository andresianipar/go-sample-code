package gosamplecode

import library "github.com/andresianipar/go-sample-code/internal"

// F10 function
func F10() int {
	total := library.Sum(1, 2, 3, 4, 5)

	// fmt.Println(total)

	total = library.Sum([]int{1, 2, 3, 4, 5}...)

	// fmt.Println(total)

	// fmt.Printf("%T != %T\n", func(...int) {}, func([]int) {})

	return total
}

// F11 function
func F11() {
	linenum, name := 15, "count"

	library.Errorf(linenum, "undefined: %s", name)
}
