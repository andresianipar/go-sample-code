package gosamplecode

import (
	"time"

	library "github.com/andresianipar/go-sample-code/internal"
)

// F19 function
func F19() library.Worker {
	var johnnyBravo library.Worker

	johnnyBravo.ID = 1009
	johnnyBravo.Name = "Johnny Bravo"
	johnnyBravo.Address = "Detroit"
	johnnyBravo.DoB = time.Date(1998, time.May, 15, 0, 0, 0, 0, time.UTC)

	position := &johnnyBravo.Position

	*position = "Senior Software Engineer"
	johnnyBravo.Salary = 5000

	// fmt.Println(johnnyBravo)

	return johnnyBravo
}

// F20 function
func F20() library.Worker {
	w := library.Worker{
		ID:       1010,
		Name:     "Johnny Bravo",
		DoB:      time.Date(1998, time.May, 15, 0, 0, 0, 0, time.UTC),
		Address:  "Detroit",
		Position: "Senior Software Engineer",
		Salary:   5000,
	}

	// fmt.Println(w)

	return w
}

// F21 function
func F21(w library.Worker) library.Worker {
	w.Position = "Principal Software Engineer"
	w.Salary = 10000

	// fmt.Println(w)

	return w
}

// F22 function
func F22(w library.Worker, position string, salary int) *library.Worker {
	w.Position = position
	w.Salary = salary

	// fmt.Println(w)

	return &w
}

// F23 function
func F23() (library.Worker, library.Worker) {
	w1 := F20()
	w2 := F20()

	// fmt.Println(w1 == w2)

	w1.Salary += 10000

	// fmt.Println(w1 == w2)

	return w1, w2
}

// F24 function
func F24() int {
	type address struct {
		hostname string
		port     int
	}

	hits := make(map[address]int)

	hits[address{"golang.org", 443}]++

	// fmt.Println(hits)

	return hits[address{"golang.org", 443}]
}

// F25 function
func F25() (library.Wheel1, library.Wheel2) {
	var w1 library.Wheel1

	w1.Circle.Center.X = 8
	w1.Circle.Center.Y = 8
	w1.Circle.Radius = 5
	w1.Spokes = 20

	// fmt.Println(w1)

	var w2 library.Wheel2

	w2.X = 8
	w2.Y = 8
	w2.Radius = 5
	w2.Spokes = 20

	// fmt.Println(w2)

	return w1, w2
}

// F26 function
func F26() (library.Wheel2, library.Wheel2) {
	w1 := library.Wheel2{
		library.Circle2{
			library.Point{8, 8},
			5,
		},
		20,
	}
	w2 := library.Wheel2{
		Circle2: library.Circle2{
			Point:  library.Point{X: 8, Y: 8},
			Radius: 5,
		},
		Spokes: 20,
	}

	// fmt.Println(w1)
	// fmt.Println("")
	// fmt.Println(w2)

	return w1, w2
}
