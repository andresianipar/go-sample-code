package gosamplecode

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	library "github.com/andresianipar/go-sample-code/internal"
)

// F4 function
func F4() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)

	for input.Scan() {
		if "exit" == input.Text() {
			break
		}

		counts[input.Text()]++
	}

	if err := input.Err(); nil != err {
		fmt.Fprintf(os.Stderr, "%v\n", err)

		os.Exit(1)
	}

	fmt.Printf("%s\t%s\n", "Content", "Count")

	for line, n := range counts {
		if 1 < n {
			fmt.Printf("%s\t%d\n", line, n)
		}
	}
}

// F5 function
func F5() {
	counts := make(map[string]int)
	files := os.Args[1:]

	if 0 == len(files) {
		library.CountLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)

			if nil != err {
				fmt.Fprintf(os.Stderr, "%v\n", err)

				continue
			}

			library.CountLines(f, counts)
			f.Close()
		}
	}

	fmt.Printf("%s\t%s\n", "Content", "Count")

	for line, n := range counts {
		if 1 < n {
			fmt.Printf("%s\t%d\n", line, n)
		}
	}
}

// F6 function
func F6() {
	counts := make(map[string]int)

	for _, fileName := range os.Args[1:] {
		data, err := ioutil.ReadFile(fileName)

		if nil != err {
			fmt.Fprintf(os.Stderr, "%v\n", err)

			continue
		}

		for _, line := range strings.Split(string(data), "\n") {
			if "" != line {
				counts[fileName+";"+line]++
			}
		}
	}

	fmt.Printf("%s\t\t%s\t%s\n", "File", "Content", "Count")

	for line, n := range counts {
		if 1 < n {
			tmp := strings.Split(line, ";")

			fmt.Printf("%s\t%s\t%d\n", tmp[0], tmp[1], n)
		}
	}
}
