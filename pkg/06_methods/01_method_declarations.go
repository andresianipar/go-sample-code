package gosamplecode

import library "github.com/andresianipar/go-sample-code/internal"

// F1 function
func F1() (float64, float64, float64, float64) {
	p := library.Point64{X: 1, Y: 2}
	q := library.Point64{X: 4, Y: 6}
	distance1 := library.Distance(p, q)
	distance2 := p.Distance(q)

	// fmt.Println(distance1)
	// fmt.Println(distance2)

	perim := library.Path{
		{X: 1, Y: 1},
		{X: 5, Y: 1},
		{X: 5, Y: 4},
		{X: 1, Y: 1},
	}
	distance3 := library.Path.Distance(perim)
	distance4 := perim.Distance()

	// fmt.Println(distance3)
	// fmt.Println(distance4)

	return distance1, distance2, distance3, distance4
}
