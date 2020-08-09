package gosamplecode

import (
	"os"

	html "golang.org/x/net/html"

	library "github.com/andresianipar/go-sample-code/internal"
)

// F1 function
func F1() []string {
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

	links := library.Visit1(nil, doc, resp)

	// for _, link := range links {
	// 	fmt.Println(link)
	// }

	return links
}

// F2 function
func F2() []string {
	resp, err := library.FetchURL1("https://golang.org")

	if nil != err {
		// fmt.Fprintf(os.Stderr, "%v\n", err)

		os.Exit(1)
	}

	doc, err := html.Parse(resp.Body)

	resp.Body.Close()

	if nil != err {
		// fmt.Fprintf(os.Stderr, "%v\n", err)

		os.Exit(1)
	}

	tags := library.Outline(nil, doc)

	// for _, tag := range tags {
	// 	fmt.Println(tag)
	// }

	return tags
}
