package gosamplecode

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"runtime"
	"sort"
	"strings"
	"time"

	html "golang.org/x/net/html"
)

// 5 functions > recursion

// FetchURL1 function
func FetchURL1(url string) (*http.Response, error) {
	resp, err := http.Get(url)

	if nil != err {
		// fmt.Fprintf(os.Stderr, "%v\n", err)

		os.Exit(1)
	}

	if http.StatusOK != resp.StatusCode {
		resp.Body.Close()

		return nil, fmt.Errorf("%v", err)
	}

	return resp, nil
}

// Visit1 function
// Visit1 appends to links each link found in n and returns the result.
func Visit1(links []string, n *html.Node, resp *http.Response) []string {
	if html.ElementNode == n.Type && "a" == n.Data {
		for _, a := range n.Attr {
			if "href" == a.Key {
				link, err := resp.Request.URL.Parse(a.Val)

				if nil != err {
					continue
				}

				links = append(links, link.String())
			}
		}
	}

	for c := n.FirstChild; nil != c; c = c.NextSibling {
		links = Visit1(links, c, resp)
	}

	return links
}

// Outline function
func Outline(stack []string, n *html.Node) []string {
	if html.ElementNode == n.Type {
		// push tag
		stack = append(stack, n.Data)

		// fmt.Println(stack)
	}

	for c := n.FirstChild; nil != c; c = c.NextSibling {
		stack = Outline(stack, c)
	}

	return stack
}

// 5 functions > errors > error-handling strategies

// WaitForServer function
func WaitForServer(url string) error {
	const timeout = 1 * time.Minute
	deadline := time.Now().Add(timeout)

	for tries := 0; time.Now().Before(deadline); tries++ {
		_, err := http.Head(url)

		if nil == err {
			return nil
		}

		// log.Printf("server not responding (%s); retrying...", err)

		time.Sleep(time.Second << uint(tries))
	}

	return fmt.Errorf("server %s failed to respond after %s", url, timeout)
}

// 5 functions > function values

// Square function
func Square(n int) int { return n * n }

// AddRune function
func AddRune(r rune) rune { return r + 1 }

var depth int

// Visit2 function
func Visit2(stack []string, n *html.Node, pre, post func(n *html.Node) string) []string {
	if nil != pre {
		stack = append(stack, pre(n))
		depth++
	}

	for c := n.FirstChild; nil != c; c = c.NextSibling {
		stack = Visit2(stack, c, pre, post)
	}

	if nil != post {
		depth--
		stack = append(stack, post(n))
	}

	return stack
}

// StartElement function
func StartElement(n *html.Node) string {
	if html.ElementNode == n.Type {
		return fmt.Sprintf("%*s<%s>", (depth-1)*2, "", n.Data)
	}

	return ""
}

// EndElement function
func EndElement(n *html.Node) string {
	if html.ElementNode == n.Type {
		return fmt.Sprintf("%*s<%s>", (depth-1)*2, "", n.Data)
	}

	return ""
}

// 5 functions > anonymous functions

// Squares function
func Squares() func() int {
	var x int

	return func() int {
		x++

		return x * x
	}
}

// TopoSort function
func TopoSort(p map[string][]string) []string {
	var order []string
	var visitAll func(items []string)
	seen := make(map[string]bool)

	visitAll = func(items []string) {
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				visitAll(p[item])
				order = append(order, item)
			}
		}
	}

	var keys []string

	for key := range p {
		keys = append(keys, key)
	}

	sort.Strings(keys)
	visitAll(keys)

	return order
}

// Crawl1 function
func Crawl1(url string) []string {
	resp, err := FetchURL1(url)

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

	links := Visit1(nil, doc, resp)

	return links
}

// BreadthFirst1 function
// BreadthFirst1 calls c for each item in the urls.
// Any items returned by c are added to the urls.
// c is called at most once for each item.
func BreadthFirst1(c func(item string) []string, urls []string) []string {
	seen := make(map[string]bool)

	for len(urls) <= 100 {
		items := urls
		urls = nil

		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				urls = append(urls, c(item)...)
			}
		}
	}

	return urls
}

// 5 functions > variadic functions

// Sum function
func Sum(vals ...int) int {
	total := 0

	for _, val := range vals {
		total += val
	}

	return total
}

// Errorf function
func Errorf(linenum int, format string, args ...interface{}) {
	// fmt.Fprintf(os.Stderr, "Line %d: ", linenum)
	// fmt.Fprintf(os.Stderr, format, args...)
	// fmt.Fprintln(os.Stderr)
}

// 5 functions > deferred function calls

// Crawl2 function
func Crawl2(url string) ([]string, error) {
	resp, err := FetchURL1(url)

	if nil != err {
		// fmt.Fprintf(os.Stderr, "%v\n", err)

		os.Exit(1)
	}

	defer resp.Body.Close()

	if ct := resp.Header.Get("Content-Type"); "text/html" != ct && !strings.HasPrefix(ct, "text/html;") {
		// resp.Body.Close()

		return nil, fmt.Errorf("title: %s has type %s, not text/html", url, ct)
	}

	doc, err := html.Parse(resp.Body)

	// resp.Body.Close()

	if nil != err {
		// fmt.Fprintf(os.Stderr, "%v\n", err)

		os.Exit(1)
	}

	links := Visit1(nil, doc, resp)

	return links, nil
}

// BreadthFirst2 function
// BreadthFirst2 calls c for each item in the urls.
// Any items returned by c are added to the urls.
// c is called at most once for each item.
func BreadthFirst2(c func(item string) ([]string, error), urls []string) ([]string, error) {
	seen := make(map[string]bool)

	for len(urls) <= 100 {
		items := urls
		urls = nil

		for _, item := range items {
			if !seen[item] {
				seen[item] = true

				r, err := c(item)

				if nil != err {
					return nil, fmt.Errorf("%v", err)
				}

				urls = append(urls, r...)
			}
		}
	}

	return urls, nil
}

// Trace function
func Trace(msg string) func() {
	start := time.Now()

	log.Printf("enter %s", msg)

	return func() { log.Printf("exit %s (%s)", msg, time.Since(start)) }
}

// Double function
func Double(x int) (result int) {
	defer func() {
		// fmt.Printf("double(%d) = %d\n", x, result)
	}()

	return x + x
}

// Triple function
func Triple(x int) (result int) {
	defer func() { result += x }()

	return x + x
}

// FetchBasePath function
// FetchBasePath downloads the URL and returns the
// name and length of the local file.
func FetchBasePath(url string) (filename string, n int64, err error) {
	resp, err := http.Get(url)

	if nil != err {
		return "", 0, err
	}

	defer resp.Body.Close()

	local := path.Base(resp.Request.URL.Path)

	if "/" == local {
		local = "index.html"
	}

	f, err := os.Create(local)

	if nil != err {
		return "", 0, err
	}

	n, err = io.Copy(f, resp.Body)

	// Close file, but prefer error from Copy, if any.
	if closeErr := f.Close(); nil == err {
		err = closeErr
	}

	return local, n, err
}

// 5 functions > panic

// F function
func F(x int) {
	// fmt.Printf("f(%d)\n", x+0/x)

	defer fmt.Printf("defer f(%d)\n", x)

	F(x - 1)
}

// PrintStack function
func PrintStack() {
	var buff [4096]byte

	n := runtime.Stack(buff[:], false)

	os.Stdout.Write(buff[:n])
}

// 5 functions > recover

func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if nil != pre {
		pre(n)
	}

	for c := n.FirstChild; nil != c; c = c.NextSibling {
		forEachNode(c, pre, post)
	}

	if nil != post {
		post(n)
	}
}

// SoleTitle function
// SoleTitle returns the text of the first non-empty title element
// in doc, and an error if there was not exactly one.
func SoleTitle(doc *html.Node) (title string, err error) {
	type bailout struct{}

	defer func() {
		switch p := recover(); p {
		case nil:
			// no panic
		case bailout{}:
			// "expected" panic
			err = fmt.Errorf("multiple title elements")
		default:
			// unexpected panic; carry on panicking
			panic(p)
		}
	}()

	// Bail out of recursion if we find more than one non-empty title.
	forEachNode(doc, func(n *html.Node) {
		if html.ElementNode == n.Type && "title" == n.Data && nil != n.FirstChild {
			if "" != title {
				// multiple title elements
				panic(bailout{})
			}

			title = n.FirstChild.Data
		}
	}, nil)

	if "" == title {
		return "", fmt.Errorf("no title element")
	}

	return title, nil
}
