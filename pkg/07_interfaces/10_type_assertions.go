package gosamplecode

import library "github.com/andresianipar/go-sample-code/internal"

// F14 function
// The type assertion x.(T) asserts that the concrete value stored in x is of type T, and that x is not nil
// If x is a nil interface value, the type assertion fails
func F14() (bool, bool) {
	var i library.Interface1 = new(library.Type1)
	// T is an interface, so it asserts that the dynamic type of x implements T
	_, ok1 := i.(library.Interface1)

	// fmt.Println(t1)
	// true
	// fmt.Println(ok1)

	// T is not an interface, so it asserts that the dynamic type of x is identical to T
	_, ok1 = i.(*library.Type1)

	// fmt.Println(t1)
	// true
	// fmt.Println(ok1)

	// t2 := i.(*library.Type2)
	// panic: interface conversion: gosamplecode.Interface1 is *gosamplecode.Type1, not *gosamplecode.Type2
	// fmt.Println(t2)

	_, ok2 := i.(*library.Type2)
	// fmt.Println(t2)
	// false
	// fmt.Println(ok2)

	return ok1, ok2
}
