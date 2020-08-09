package gosamplecode

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	html "golang.org/x/net/html"
)

// 8 goroutines and channels > goroutines

// Spinner function
func Spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}

// Fib function
func Fib(x int) int {
	if x < 2 {
		return x
	}

	return Fib(x-1) + Fib(x-2)
}

// 8 goroutines and channels > example: concurrent clock server

// HandleConn1 function
func HandleConn1(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))

		if err != nil {
			return // e.g., client disconnected
		}
		time.Sleep((1 * time.Second))
	}
}

// 8 goroutines and channels > example: concurrent echo server

func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}

// HandleConn2 function
func HandleConn2(c net.Conn) {
	defer c.Close()

	input := bufio.NewScanner(c)

	for input.Scan() {
		// echo(c, input.Text(), 1*time.Second)
		go echo(c, input.Text(), 1*time.Second)
	}
	if input.Err() != nil {
		return
	}
}

// 8 goroutines and channels > channels > unbuffered channels

// MustCopy function
func MustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}

// 8 goroutines and channels > channels > unidirectional channel types

// ProduceNumbers function
func ProduceNumbers(out chan<- int) {
	for x := 0; x <= 4; x++ {
		out <- x
	}
	close(out)
}

// Squarer function
func Squarer(out chan<- int, in <-chan int) {
	for v := range in {
		out <- v * v
	}
	close(out)
}

// Printer function
func Printer(in <-chan int) int {
	var v int

	for v = range in {
		fmt.Println(v)
	}

	return v
}

// 8 goroutines and channels > example: concurrent web crawler

// FetchURL2 function
func FetchURL2(url string) (*http.Response, error) {
	resp, err := http.Get(url)

	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()

		return nil, fmt.Errorf("%v", err)
	}

	return resp, nil
}

// Visit3 function
func Visit3(links []string, n *html.Node, resp *http.Response) []string {
	if html.ElementNode == n.Type && "a" == n.Data {
		for _, a := range n.Attr {
			if a.Key == "href" {
				link, err := resp.Request.URL.Parse(a.Val)

				if err != nil {
					continue
				}
				links = append(links, link.String())
			}
		}
	}
	for c := n.FirstChild; nil != c; c = c.NextSibling {
		links = Visit3(links, c, resp)
	}

	return links
}

// Tokens is a counting semaphore used to
// enforce a limit of 20 concurrent requests.
var tokens = make(chan struct{}, 20)

// Crawl3 function
func Crawl3(url string) ([]string, error) {
	resp, err := FetchURL2(url)

	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}
	defer resp.Body.Close()
	if ct := resp.Header.Get("Content-Type"); ct != "text/html" && !strings.HasPrefix(ct, "text/html;") {
		return nil, fmt.Errorf("title: %s has type %s, not text/html", url, ct)
	}

	doc, err := html.Parse(resp.Body)

	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}
	tokens <- struct{}{} // acquire a token

	links := Visit3(nil, doc, resp)

	<-tokens // release the token

	return links, nil
}

// BreadthFirst3 function
func BreadthFirst3(c func(item string) ([]string, error), worklist chan []string) {
	var n int // number of pending sends to worklist
	// worklist := make(chan []string)

	// Start with the command-line arguments.
	n++
	go func() { worklist <- os.Args[1:] }()

	// Crawl the web concurrently.
	seen := make(map[string]bool)

	// for list := range worklist {
	for ; n > 0; n-- {
		list := <-worklist

		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				n++
				go func(link string) {
					links, err := c(link)

					if err != nil {
						fmt.Fprintf(os.Stderr, "%v\n", err)
					}
					worklist <- links
				}(link)
			}
		}
	}
}

// BreadthFirst4 function
// func BreadthFirst4(c func(item string) ([]string, error)) chan []string {
func BreadthFirst4(c func(item string) ([]string, error), worklist chan []string) {
	// worklist := make(chan []string)  // lists of URLs, may have duplicates
	unseenLinks := make(chan string) // de-duplicated URLs

	// Add command-line arguments to worklist.
	go func() { worklist <- os.Args[1:] }()
	// Create 20 crawler goroutines to fetch each unseen link.
	for i := 0; i < 20; i++ {
		go func() {
			for link := range unseenLinks {
				links, err := c(link)

				if err != nil {
					fmt.Fprintf(os.Stderr, "%v\n", err)
				}
				go func() { worklist <- links }()
			}
		}()
	}

	// The main goroutine de-duplicates worklist items
	// and sends the unseen ones to the crawlers.
	seen := make(map[string]bool)

	for list := range worklist {
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				unseenLinks <- link
			}
		}
	}
}

// 8 goroutines and channels > example: concurrent directory traversal

// dirents1 returns the entries of directory dir.
func dirents1(dirname string) []os.FileInfo {
	entries, err := ioutil.ReadDir(dirname)

	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)

		return nil
	}

	return entries
}

// WalkDir1 function
// WalkDir1 recursively walks the file tree rooted at dir
// and sends the size of each found file on fileSizes.
func WalkDir1(dirname string, fileSizes chan<- int64) {
	for _, entry := range dirents1(dirname) {
		if entry.IsDir() {
			subdir := filepath.Join(dirname, entry.Name())

			WalkDir1(subdir, fileSizes)
		} else {
			fileSizes <- entry.Size()
		}
	}
}

// sema is a counting semaphore for limiting concurrency in dirents.
var sema = make(chan struct{}, 20)

// dirents2 returns the entries of directory dir.
func dirents2(dirname string) []os.FileInfo {
	sema <- struct{}{}        // acquire token
	defer func() { <-sema }() // relase token

	entries, err := ioutil.ReadDir(dirname)

	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)

		return nil
	}

	return entries
}

// WalkDir2 function
func WalkDir2(dirname string, n *sync.WaitGroup, fileSizes chan<- int64) {
	defer n.Done()
	for _, entry := range dirents2(dirname) {
		if entry.IsDir() {
			n.Add(1)

			subdir := filepath.Join(dirname, entry.Name())

			go WalkDir2(subdir, n, fileSizes)
		} else {
			fileSizes <- entry.Size()
		}
	}
}

// 8 goroutines and channels > cancellation

// Done channel
var Done = make(chan struct{})

func cancelled() bool {
	select {
	case <-Done:
		return true
	default:
		return false
	}
}

// dirents3 returns the entries of directory dir.
func dirents3(dirname string) []os.FileInfo {
	select {
	case sema <- struct{}{}: // acquire token
	case <-Done:
		return nil // cancelled
	}
	defer func() { <-sema }() // relase token

	entries, err := ioutil.ReadDir(dirname)

	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)

		return nil
	}

	return entries
}

// WalkDir3 function
func WalkDir3(dirname string, n *sync.WaitGroup, fileSizes chan<- int64) {
	defer n.Done()
	if cancelled() {
		return
	}
	for _, entry := range dirents3(dirname) {
		if entry.IsDir() {
			n.Add(1)

			subdir := filepath.Join(dirname, entry.Name())

			go WalkDir2(subdir, n, fileSizes)
		} else {
			fileSizes <- entry.Size()
		}
	}
}

// 8 goroutines and channels > example chat server

type client chan<- string // an outgoing message channel

var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan string) // all incoming client messages
)

// Broadcaster function
func Broadcaster() {
	clients := make(map[client]bool) // all connected clients

	for {
		select {
		case msg := <-messages:
			// Broadcast incoming message to all
			// clients' outgoing message channels.
			for cli := range clients {
				cli <- msg
			}
		case cli := <-entering:
			clients[cli] = true
		case cli := <-leaving:
			delete(clients, cli)
			close(cli)
		}
	}
}

func clientWriter(c net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(c, msg) // NOTE: ignoring network errors
	}
}

// HandleConn3 function
func HandleConn3(c net.Conn) {
	defer c.Close()

	ch := make(chan string) // outgoing client messages

	go clientWriter(c, ch)

	who := c.RemoteAddr().String()

	ch <- "You are " + who
	messages <- who + " has arrived"
	entering <- ch

	input := bufio.NewScanner(c)

	for input.Scan() {
		messages <- who + ": " + input.Text()
	}
	if input.Err() != nil {
		return
	}
	leaving <- ch
	messages <- who + " has left"
}
