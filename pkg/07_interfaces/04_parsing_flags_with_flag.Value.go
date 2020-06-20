package gosamplecode

import (
	"flag"
	"time"

	library "github.com/andresianipar/go-sample-code/internal"
)

// F3 function
func F3() *time.Duration {
	var period = flag.Duration("period", 1*time.Second, "sleep period")

	flag.Parse()

	// fmt.Printf("Sleeping for %v...", *period)

	time.Sleep(*period)

	// fmt.Println()

	return period
}

// F4 function
func F4() *library.Celsius {
	var temp = library.CelciusFlag("temp", 20.0, "the temperature")

	flag.Parse()

	// fmt.Println(*temp)

	return temp
}
