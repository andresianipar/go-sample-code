package gosamplecode

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"sync"
	"time"
)

// 1 tutorial > finding duplicate lines

// CountLines function
func CountLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)

	for input.Scan() {
		counts[input.Text()]++
	}
}

// Result struct
type Result struct {
	status string
	data   []Employee
}

// Employee struct
type Employee struct {
	id             int
	employeeName   string
	employeeSalary int
	employeeAge    int
	profileImage   string
}

// 1 tutorial > fetching urls concurrently

// Fetch function
func Fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)

	if nil != err {
		ch <- fmt.Sprint(err)

		return
	}

	nbytes, err := io.Copy(ioutil.Discard, resp.Body)

	resp.Body.Close()

	if nil != err {
		ch <- fmt.Sprint(err)

		return
	}

	secs := time.Since(start).Seconds()

	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}

// 1 tutorial > a web server

// Handler1 echoes the Path component of the request URL r
func Handler1(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

var mutex sync.Mutex
var count int

// Handler2 echoes the Path component of the requested URl
func Handler2(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	count++
	mutex.Unlock()

	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

// Counter echoes the number of calls so far
func Counter(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()

	fmt.Fprintf(w, "Count: %d\n", count)

	mutex.Unlock()
}

// Handler3 echoes the HTTP request
func Handler3(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Rebuilt URL to: %s\n", r.Host)
	fmt.Fprintf(w, "  Trying %s...\n", r.RemoteAddr)
	fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
	fmt.Fprintf(w, "Host: %q\n", r.Host)
	fmt.Fprintf(w, "User-Agent: %q\n", r.Header["User-Agent"])
	fmt.Fprintf(w, "Accept: %q\n\n", r.Header["Accept"])
	fmt.Fprintf(w, "%s\n", r.Proto)
	fmt.Fprintf(w, "Date:\n")
	fmt.Fprintf(w, "Content-Length: %s\n", r.Header["Content-Length"])
	fmt.Fprintf(w, "Content-Type: %s\n", r.Header["Content-Type"])
}
