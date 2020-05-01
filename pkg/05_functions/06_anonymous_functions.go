package gosamplecode

import (
	"time"

	library "github.com/andresianipar/go-sample-code/internal"
)

// F6 function
func F6() int {
	f := library.Squares()

	// fmt.Println(f())
	// fmt.Println(f())
	// fmt.Println(f())
	// fmt.Println(f())

	f()
	f()
	f()
	f()

	return f()
}

// F7 function
func F7() []string {
	var prereqs = map[string][]string{
		"algorithms": {"data structures"},
		"calculus":   {"linear algebra"},
		"compilers": {
			"data structures",
			"formal languages",
			"computer organization",
		},
		"data structures":  {"discrete math"},
		"databases":        {"data structures"},
		"discrete math":    {"intro to programming"},
		"formal languages": {"discrete math"},
		"networks":         {"operating systems"},
		"operating systems": {
			"data structures",
			"computer organization",
		},
		"programming languages": {
			"data structures",
			"computer organization",
		},
	}

	seq := library.TopoSort(prereqs)

	// for i, course := range seq {
	// 	fmt.Printf("%d:\t%s\n", i+1, course)
	// }

	return seq
}

// F8 function
func F8() []string {
	urls := library.BreadthFirst1(library.Crawl1, []string{"https://golang.org"})

	// for i, url := range urls {
	// 	fmt.Printf("%d:\t%s\n", i+1, url)
	// }

	return urls
}

// F9 function
func F9() {
	// data := []string{"1", "2", "3", "4"}

	// for _, v := range data {
	// 	v := v

	// 	go func() {
	// 		fmt.Println(v)
	// 	}()
	// }

	time.Sleep(4 * time.Second)
}
