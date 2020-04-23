package gosamplecode

import (
	"crypto/sha256"
)

// F1 function
func F1() (int, int) {
	var q [3]int = [3]int{1, 2}

	// fmt.Println(q[2])

	r := [...]int{1, 2, 3}

	// fmt.Printf("%T\n", r)

	return q[2], len(r)
}

// F2 function
func F2() string {
	type Currency int

	const (
		USD Currency = iota
		EUR
		GBP
		RMB
	)

	symbol := [...]string{USD: "$", EUR: "€", GBP: "£", RMB: "¥"}

	// fmt.Println(symbol)
	// fmt.Println(symbol[USD])

	return symbol[USD]
}

// F3 function
func F3() (int, int) {
	r := [...]int{99: -1}

	// fmt.Println(len(r))
	// fmt.Println(r[99])

	return len(r), r[99]
}

// F4 function
func F4() (bool, bool) {
	a := [2]int{1, 2}
	b := [...]int{3, 4}
	c := [3]int{1, 2, 3}
	d := [...]int{1, 2, 3}

	// fmt.Println(a == b)
	// fmt.Println(c == d)

	return a == b, c == d
}

// F5 function
func F5() bool {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))

	// fmt.Printf("%x\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1)

	return c1 == c2
}
