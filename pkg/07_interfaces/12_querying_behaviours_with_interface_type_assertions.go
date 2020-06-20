package gosamplecode

import library "github.com/andresianipar/go-sample-code/internal"

// F16 function
func F16() {
	var i library.Interface1 = new(library.Type3)

	library.DoThis(i)

	i = new(library.Type4)

	library.DoThis(i)
}
