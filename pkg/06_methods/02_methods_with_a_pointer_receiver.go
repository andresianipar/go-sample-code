package gosamplecode

import library "github.com/andresianipar/go-sample-code/internal"

// F2 function
func F2() (library.Point64, library.Point64, float64) {
	p := library.Point64{X: 1, Y: 2}

	// The receiver argument has the same type as the receiver parameter
	pPtr1 := &library.Point64{X: 1, Y: 2}

	pPtr1.ScaleBy(2)

	// fmt.Println(*pPtr1)

	// The receiver argument has the same type as the receiver parameter
	p = library.Point64{X: 1, Y: 2}
	pPtr2 := &p

	pPtr2.ScaleBy(2)

	// fmt.Println(p)

	// The receiver argument has the same type as the receiver parameter
	p = library.Point64{X: 1, Y: 2}

	(&p).ScaleBy(2)

	// fmt.Println(p)

	// The receiver argument is a variable of type T and the receiver parameter has type *T
	p = library.Point64{X: 1, Y: 2}

	p.ScaleBy(2)

	// fmt.Println(p)

	// The receiver argument has type *T and the receiver parameter has type T
	p = library.Point64{X: 1, Y: 2}
	pPtr3 := &p

	p.ScaleBy(2)

	distance := pPtr3.Distance(library.Point64{X: 4, Y: 4})

	// fmt.Println(distance)

	return *pPtr1, p, distance
}

// F3 function
func F3() int {
	list4 := library.IntList{Value: 4, Tail: nil}
	list3 := library.IntList{Value: 3, Tail: &list4}
	list2 := library.IntList{Value: 2, Tail: &list3}
	list1 := library.IntList{Value: 1, Tail: &list2}
	sum := list1.Sum()

	// fmt.Println(sum)

	return sum
}
