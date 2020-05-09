package gosamplecode

import (
	"image/color"
	"math"
)

// 6 methods > method declarations

// Point64 struct
type Point64 struct{ X, Y float64 }

// Distance function
func Distance(p, q Point64) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

// Distance function
func (p Point64) Distance(q Point64) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

// Path type
// A Path is a journey connecting the points with straight lines.
type Path []Point64

// Distance function
// Distance returns the distance traveled along the path.
func (path Path) Distance() float64 {
	sum := 0.0

	for i := range path {
		if i > 0 {
			sum += path[i-1].Distance(path[i])
		}
	}

	return sum
}

// 6 methods > methods with a pointer receiver

// ScaleBy function
func (p *Point64) ScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}

// 6 methods > methods with a pointer receiver > nil is a valid receiver value

// IntList struct
// An IntList is a linked list of integers.
// A nil *IntList represents the empty list.
type IntList struct {
	Value int
	Tail  *IntList
}

// Sum function
// Sum returns the sum of the list elements.
func (list *IntList) Sum() int {
	if nil == list {
		return 0
	}

	return list.Value + list.Tail.Sum()
}

// 6 methods > composing types by struct embedding

// ColoredPoint1 struct
type ColoredPoint1 struct {
	Point64
	Color color.RGBA
}

// ColoredPoint2 struct
type ColoredPoint2 struct {
	*Point64
	Color color.RGBA
}

// 6 methods > method values and expressions

// Add function
func (p Point64) Add(q Point64) Point64 {
	return Point64{p.X + q.X, p.Y + q.Y}
}

// Sub function
func (p Point64) Sub(q Point64) Point64 {
	return Point64{p.X - q.X, p.Y - q.Y}
}

// TranslateBy function
func (path Path) TranslateBy(offset Point64, add bool) {
	var op func(p, q Point64) Point64

	if add {
		op = Point64.Add
	} else {
		op = Point64.Sub
	}

	for i := range path {
		// Call either path[i].Add(offset) or path[i].Sub(offset).
		path[i] = op(path[i], offset)
	}
}
