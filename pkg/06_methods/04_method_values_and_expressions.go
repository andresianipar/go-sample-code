package gosamplecode

import library "github.com/andresianipar/go-sample-code/internal"

// F6 function
func F6() (float64, library.Point64) {
	p := library.Point64{X: 1, Y: 2}
	q := library.Point64{X: 4, Y: 6}

	distanceFromP := p.Distance
	distance := distanceFromP(q)

	// fmt.Println(distance)

	scaleP := p.ScaleBy

	scaleP(2)

	// fmt.Println(p)

	return distance, p
}

// F7 function
func F7() (float64, library.Point64) {
	p := library.Point64{X: 1, Y: 2}
	q := library.Point64{X: 4, Y: 6}

	distanceFromTo := library.Point64.Distance
	distance := distanceFromTo(p, q)

	// fmt.Println(distance)

	scaleBy := (*library.Point64).ScaleBy

	scaleBy(&p, 2)

	// fmt.Println(p)

	return distance, p
}

// F8 function
func F8() library.Path {
	p := library.Point64{X: 1, Y: 2}
	q := library.Point64{X: 4, Y: 6}
	path := library.Path{p}

	path.TranslateBy(q, true)

	// fmt.Println(path)

	return path
}
