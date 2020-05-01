package gosamplecode

import library "github.com/andresianipar/go-sample-code/internal"

// F3 function
func F3() error {
	// log.SetPrefix("wait: ")
	// log.SetFlags(0)

	if err := library.WaitForServer("https://xyz.org"); nil != err {
		// fmt.Fprintf(os.Stderr, "site is down: %v\n", err)

		// os.Exit(1)

		// log.Fatalf("site is down: %v\n", err)

		return err
	}

	return nil
}
