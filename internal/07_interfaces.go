package gosamplecode

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"syscall"
	"text/tabwriter"
	"time"
)

// 7 interfaces > interfaces as contracts

// ByteCounter type
type ByteCounter int

func (c *ByteCounter) Write(p []byte) (int, error) {
	n := len(p)

	*c += ByteCounter(n)

	return n, nil
}

// 7 interfaces > parsing flags with flag.Value

// *celsiusFlag satisfies the flag.Value interface.
type celciusFlag struct{ Celsius }

func (cF *celciusFlag) Set(s string) error {
	var unit string
	var value float64

	// no error check needed
	// if there was a problem, no switch case will match
	fmt.Sscanf(s, "%f%s", &value, &unit)

	switch unit {
	case "C", "°C":
		cF.Celsius = Celsius(value)

		return nil
	case "F", "°F":
		cF.Celsius = Celsius((value - 32) * 5 / 9)

		return nil
	}

	return fmt.Errorf("invalid temperature %q", s)
}

// CelciusFlag function
func CelciusFlag(name string, value Celsius, usage string) *Celsius {
	f := celciusFlag{value}

	flag.CommandLine.Var(&f, name, usage)

	return &f.Celsius
}

// 7 interfaces > interface values

// OutWrite function
func OutWrite(out io.Writer) {
	if nil != out {
		out.Write([]byte("Hello, World!"))
	}
}

// 7 interfaces > sorting with sort.interface

// StringSlice type
type StringSlice []string

func (ss StringSlice) Len() int           { return len(ss) }
func (ss StringSlice) Less(i, j int) bool { return ss[i] < ss[j] }
func (ss StringSlice) Swap(i, j int)      { ss[i], ss[j] = ss[j], ss[i] }

// Track type
type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

// Length function
func Length(s string) time.Duration {
	d, err := time.ParseDuration(s)

	if nil != err {
		panic(s)
	}

	return d
}

// PrintTracks function
func PrintTracks(tracks []*Track) {
	const format = "%v\t%v\t%v\t%v\t%v\t\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)

	fmt.Fprintf(tw, format, "Title", "Artist", "Album", "Year", "Length")
	fmt.Fprintf(tw, format, "-----", "------", "-----", "----", "------")

	for _, t := range tracks {
		fmt.Fprintf(tw, format, t.Title, t.Artist, t.Album, t.Year, t.Length)
	}

	// calculate column widths and print table
	tw.Flush()
}

// ByArtist type
type ByArtist []*Track

func (x ByArtist) Len() int           { return len(x) }
func (x ByArtist) Less(i, j int) bool { return x[i].Artist < x[j].Artist }
func (x ByArtist) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

// ByYear type
type ByYear []*Track

func (x ByYear) Len() int           { return len(x) }
func (x ByYear) Less(i, j int) bool { return x[i].Year < x[j].Year }
func (x ByYear) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

// CustomStruct type
type CustomStruct struct {
	T          []*Track
	CustomLess func(x, y *Track) bool
}

func (x CustomStruct) Len() int           { return len(x.T) }
func (x CustomStruct) Less(i, j int) bool { return x.CustomLess(x.T[i], x.T[j]) }
func (x CustomStruct) Swap(i, j int)      { x.T[i], x.T[j] = x.T[j], x.T[i] }

// 7 interfaces > the http.handler interface

type dollars float32

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

// Database type
type Database map[string]dollars

func (db Database) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/list":
		for item, price := range db {
			fmt.Fprintf(w, "%s:\t%s\n", item, price)
		}
	case "/price":
		item := req.URL.Query().Get("item")
		price, ok := db[item]

		if !ok {
			// 404
			w.WriteHeader(http.StatusNotFound)

			fmt.Fprintf(w, "no such item: %q\n", item)

			return
		}

		fmt.Fprintf(w, "%s\n", price)
	default:
		// 404
		w.WriteHeader(http.StatusNotFound)

		fmt.Fprintf(w, "no such page: %s\n", req.URL)
	}
}

// List function
func (db Database) List(w http.ResponseWriter, req *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s:\t%s\n", item, price)
	}
}

// Price function
func (db Database) Price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price, ok := db[item]

	if !ok {
		// 404
		w.WriteHeader(http.StatusNotFound)

		fmt.Fprintf(w, "no such item: %q\n", item)

		return
	}

	fmt.Fprintf(w, "%s\n", price)
}

// 7 interfaces > type assertions

// Interface1 interface
type Interface1 interface {
	Do1()
}

// Type1 type
type Type1 int

// Do1 function
func (type1 Type1) Do1() {}

// Type2 type
type Type2 int

// Do1 function
func (type2 Type2) Do1() {}

// 7 interfaces > discriminating errors with type assertions

// ErrNotExists variable
var ErrNotExists = errors.New(("file does not exist"))

// IsNotExist function
func IsNotExist(err error) bool {
	if pe, ok := err.(*os.PathError); ok {
		err = pe.Err
	}

	return err == syscall.ENOENT || err == ErrNotExists
}

// 7 interfaces > querying behaviours with interface type assertions

// Type3 type
type Type3 int

// Do1 function
func (type3 Type3) Do1() {
	fmt.Println("Type3.Do1()")
}

// Type4 type
type Type4 int

// Do1 function
func (type4 Type4) Do1() {
	fmt.Println("Type4.Do1()")
}

// Do2 function
func (type4 Type4) Do2() {
	fmt.Println("Type4.Do2()")
}

// DoThis function
func DoThis(i1 Interface1) {
	type Do2 interface {
		Do2()
	}

	if i1, ok := i1.(Do2); ok {
		i1.Do2()

		return
	}

	i1.Do1()

	return
}
