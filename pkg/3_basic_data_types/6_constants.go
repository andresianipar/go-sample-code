package gosamplecode

import (
	"time"
)

// F22 function
func F22() (time.Duration, time.Duration) {
	const noDelay time.Duration = 0
	const timeout = 5 * time.Minute
	// pf := fmt.Printf

	// pf("%T %[1]v\n", noDelay)
	// pf("%T %[1]v\n", timeout)
	// pf("%T %[1]v\n", time.Minute)

	return noDelay, timeout
}

// F23 function
func F23() (int, int, int, int) {
	const (
		a = 1
		b
		y = 2
		z
	)

	// fmt.Println(a, b, y, z)

	return a, b, y, z
}

// F24 function
func F24() (int, int, int, int) {
	const (
		a = iota * 5
		b
		c
		d
	)

	// fmt.Println(a, b, c, d)

	return a, b, c, d
}

// F25 function
func F25() (int, int, int, int) {
	const (
		_ = 1 << (10 * iota)
		KiB
		MiB
		GiB
		TiB
	)

	// fmt.Println(KiB, MiB, GiB, TiB)

	return KiB, MiB, GiB, TiB
}
