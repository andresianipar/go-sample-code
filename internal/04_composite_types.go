package gosamplecode

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// 4 composite types > slices

// Reverse function
// Reverse reverses a slice of ints in place.
func Reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

// 4 composite types > slices > in-place slice techniques

// Nonempty1 function
// nonempty1 returns a slice holding only the non-empty strings.
// The underlying array is modified during the call.
func Nonempty1(strings []string) []string {
	i := 0

	for _, s := range strings {
		if "" != s {
			strings[i] = s
			i++
		}
	}

	return strings[:i]
}

// Nonempty2 function
func Nonempty2(strings []string) []string {
	out := strings[:0]

	for _, s := range strings {
		if "" != s {
			out = append(out, s)
		}
	}

	return out
}

// Remove function
func Remove(slice []int, i int) []int {
	copy(slice[i:], slice[i+1:])

	return slice[:len(slice)-1]
}

// 4 composite types > maps

// Equal1 function
func Equal1(x, y map[string]int) bool {
	if len(x) != len(y) {
		return false
	}

	for i, xv := range x {
		if yv, ok := y[i]; !ok || yv != xv {
			return false
		}
	}

	return true
}

// Equal2 function
func Equal2(x, y map[string]bool) bool {
	if len(x) != len(y) {
		return false
	}

	for i, xv := range x {
		if yv, ok := y[i]; !ok || yv != xv {
			return false
		}
	}

	return true
}

func k(list []string) string { return fmt.Sprintf("%q", list) }

// Add function
func Add(m map[string]int, list []string) map[string]int {
	m[k(list)]++

	return m
}

// Count function
func Count(m map[string]int, list []string) int { return m[k(list)] }

// AddEdge function
func AddEdge(graph map[string]map[string]bool, from, to string) map[string]map[string]bool {
	edges, ok := graph[from]

	if !ok {
		edges = make(map[string]bool)
		graph[from] = edges
	}

	edges[to] = true

	return graph
}

// HasEdge function
func HasEdge(graph map[string]map[string]bool, from, to string) bool { return graph[from][to] }

// 4 composite types > structs

// Worker struct
type Worker struct {
	ID            int
	Name, Address string
	DoB           time.Time
	Position      string
	Salary        int
	ManagerID     int
}

func (w Worker) String() string {
	return fmt.Sprintf(
		"ID\t\t%d\nName\t\t%s\nAddress\t\t%s\nDate of Birth\t%v\nPosition\t%s\nSalary\t\t$%d",
		w.ID,
		w.Name,
		w.Address,
		w.DoB.Format("2 January 2006"),
		w.Position,
		w.Salary,
	)
}

// 4 composite types > structs > struct embedding and anonymous fields

// Point struct
type Point struct {
	X, Y int
}

// Circle1 struct
type Circle1 struct {
	Center Point
	Radius int
}

// Wheel1 struct
type Wheel1 struct {
	Circle Circle1
	Spokes int
}

func (w Wheel1) String() string {
	return fmt.Sprintf(
		"X\t%d\nY\t%d\nRadius\t%d\nSpokes\t%d",
		w.Circle.Center.X,
		w.Circle.Center.Y,
		w.Circle.Radius,
		w.Spokes,
	)
}

// Circle2 struct
type Circle2 struct {
	Point
	Radius int
}

// Wheel2 struct
type Wheel2 struct {
	Circle2
	Spokes int
}

func (w Wheel2) String() string {
	return fmt.Sprintf(
		"X\t%d\nY\t%d\nRadius\t%d\nSpokes\t%d",
		w.X,
		w.Y,
		w.Radius,
		w.Spokes,
	)
}

// 4 composite types > json

// Movie struct
type Movie struct {
	Title  string   `json:"title"`
	Year   int      `json:"released"`
	Color  bool     `json:"color,omitempty"`
	Actors []string `json:"actors"`
}

// Movies variable
var Movies = []Movie{
	{
		Title:  "Casablanca",
		Year:   1942,
		Color:  false,
		Actors: []string{"Humphrey Bogart", "Ingrid Bergman"},
	},
	{
		Title:  "Cool Hand Luke",
		Year:   1967,
		Color:  true,
		Actors: []string{"Paul Newman"},
	},
	{
		Title:  "Bullitt",
		Year:   1968,
		Color:  true,
		Actors: []string{"Steve McQueen", "Jacqueline Bisset"},
	},
}

const issuesURL = "https://api.github.com/search/issues"

// IssuesSearchResult struct
type IssuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue
}

// Issue struct
type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	User      *User
	CreatedAt time.Time `json:"created_at"`
	Body      string
}

// User struct
type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}

// SearchIssues function
func SearchIssues(terms []string) (*IssuesSearchResult, error) {
	q := url.QueryEscape(strings.Join(terms, " "))
	resp, err := http.Get(issuesURL + "?q=" + q)

	if nil != err {
		return nil, err
	}

	if http.StatusOK != resp.StatusCode {
		resp.Body.Close()

		return nil, fmt.Errorf("Search query failed: %s", resp.Status)
	}

	var result IssuesSearchResult

	if err := json.NewDecoder(resp.Body).Decode(&result); nil != err {
		resp.Body.Close()

		return nil, err
	}

	resp.Body.Close()

	return &result, nil
}

// 4 composite types > text and html templates

// DaysAgo function
func DaysAgo(t time.Time) int {
	return int(time.Since(t).Hours() / 24)
}
