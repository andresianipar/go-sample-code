package gosamplecode

import (
	"os"

	html "golang.org/x/net/html"

	library "github.com/andresianipar/go-sample-code/internal"
)

// F19 function
func F19() string {
	resp, err := library.FetchURL("https://www.google.com")

	if nil != err {
		// fmt.Fprintf(os.Stderr, "%v\n", err)

		os.Exit(1)
	}

	doc, err := html.Parse(resp.Body)

	if nil != err {
		// fmt.Fprintf(os.Stderr, "%v\n", err)

		os.Exit(1)
	}

	title, err := library.SoleTitle(doc)

	if nil != err {
		// fmt.Fprintf(os.Stderr, "%v\n", err)

		os.Exit(1)
	}

	// fmt.Println(title)

	return title
}
