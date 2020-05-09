package gosamplecode

import (
	"image/color"

	library "github.com/andresianipar/go-sample-code/internal"
)

// F4 function
func F4() float64 {
	var cp library.ColoredPoint1

	cp.X = 1
	cp.Y = 2

	// fmt.Println(cp)

	red := color.RGBA{255, 0, 0, 255}
	blue := color.RGBA{0, 0, 255, 255}
	var p = library.ColoredPoint1{Point64: library.Point64{X: 1, Y: 1}, Color: red}
	var q = library.ColoredPoint1{Point64: library.Point64{X: 5, Y: 4}, Color: blue}
	var distance = p.Distance(q.Point64)

	// fmt.Println(distance)

	p.ScaleBy(2)
	q.ScaleBy(2)

	distance = p.Distance(q.Point64)

	// fmt.Println(distance)

	return distance
}

// F5 function
func F5() (float64, float64) {
	red := color.RGBA{255, 0, 0, 255}
	blue := color.RGBA{0, 0, 255, 255}
	var p = library.ColoredPoint2{Point64: &library.Point64{X: 1, Y: 1}, Color: red}
	var q = library.ColoredPoint2{Point64: &library.Point64{X: 5, Y: 4}, Color: blue}
	var distance1 = p.Distance(*q.Point64)

	// fmt.Println(distance1)

	q.Point64 = p.Point64

	p.ScaleBy(2)

	// fmt.Println(*p.Point64, *q.Point64)

	var distance2 = p.Distance(*q.Point64)

	// fmt.Println(distance2)

	return distance1, distance2
}
