package gosamplecode

import (
	"os"
	"strings"

	html "golang.org/x/net/html"

	library "github.com/andresianipar/go-sample-code/internal"
)

// F4 function
func F4() (func(int) int, func(int) int, string) {
	f1 := library.Square

	// fmt.Printf("%d\n", f1(4))
	// fmt.Printf("%T\n", f1)

	var f2 func(int) int

	// f2(4)

	// fmt.Println(nil == f2)

	s := strings.Map(library.AddRune, "ABCDE")

	// fmt.Println(s)

	return f1, f2, s
}

// F5 function
func F5() []string {
	resp, err := library.FetchURL1("https://golang.org")

	if nil != err {
		// fmt.Fprintf(os.Stderr, "%v\n", err)

		os.Exit(1)
	}

	doc, err := html.Parse(resp.Body)

	if nil != err {
		// fmt.Fprintf(os.Stderr, "%v\n", err)

		os.Exit(1)
	}

	tags := library.Visit2(nil, doc, library.StartElement, library.EndElement)

	// for _, tag := range tags {
	// 	if "" != tag {
	// 		fmt.Println(tag)
	// 	}
	// }

	return tags
}
