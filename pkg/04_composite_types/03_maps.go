package gosamplecode

import (
	"fmt"
	"sort"

	library "github.com/andresianipar/go-sample-code/internal"
)

// F13 function
func F13() (map[string]int, map[string]int) {
	ages1 := make(map[string]int)

	ages1["alice"] = 31
	ages1["charlie"] = 34

	// fmt.Println(ages1["alice"])
	// fmt.Println(ages1["charlie"])

	ages2 := map[string]int{
		"alice":   31,
		"charlie": 34,
	}

	// fmt.Println(ages2["alice"])
	// fmt.Println(ages2["charlie"])

	delete(ages2, "charlie")

	// fmt.Println(ages2["alice"])
	// fmt.Println(ages2["charlie"])

	// for name, age := range ages1 {
	// 	fmt.Printf("%s\t%d\n", name, age)
	// }

	return ages1, ages2
}

// F14 function
func F14() string {
	alphabet := []string{"b", "z", "c", "f", "m", "t", "a"}

	sort.Strings(alphabet)

	// for _, letter := range alphabet {
	// 	fmt.Printf("%s\n", letter)
	// }

	return fmt.Sprintf("%s", alphabet)
}

// F15 function
func F15() (int, bool, bool) {
	ages1 := map[string]int{
		"alice":   31,
		"charlie": 34,
	}
	// ages2 := map[string]int{
	// 	"alice":   31,
	// 	"charlie": 34,
	// }

	alice, ok := ages1["alice"]

	// fmt.Println(alice)
	// fmt.Println(ok)

	// fmt.Println(library.Equal1(ages1, ages2))
	// fmt.Println(library.Equal1(map[string]int{"A": 0}, map[string]int{"B": 42}))

	return alice, ok, library.Equal1(map[string]int{"A": 0}, map[string]int{"B": 42})
}

// F16 function
func F16() map[string]bool {
	seen := make(map[string]bool)
	input := [...]string{"a", "b", "a", "c", "b", "d"}

	for _, letter := range input {
		if !seen[letter] {
			seen[letter] = true

			// fmt.Println(string(letter))
		}
	}

	// fmt.Println(seen)

	return seen
}

// F17 function
func F17() (int, int) {
	var m = make(map[string]int)
	list := []string{"a", "b", "c", "d"}

	m = library.Add(m, list)

	// fmt.Println(m)

	v := m["[\"a\" \"b\" \"c\" \"d\"]"]
	c := library.Count(m, list)

	// fmt.Println(v)
	// fmt.Println(c)

	return v, c
}

// F18 function
func F18() bool {
	var graph = make(map[string]map[string]bool)

	graph = library.AddEdge(graph, "A", "B")

	// fmt.Println(graph)
	// fmt.Println(library.HasEdge(graph, "A", "B"))
	// fmt.Println(library.HasEdge(graph, "A", "C"))

	return library.HasEdge(graph, "A", "B")
}
