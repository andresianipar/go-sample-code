package gosamplecode

import (
	"sort"

	library "github.com/andresianipar/go-sample-code/internal"
)

// F6 function
func F6() ([]string, []string) {
	s1 := []string{"d", "c", "b", "a"}

	sort.Sort(library.StringSlice(s1))

	// fmt.Println(s1)

	s2 := []string{"d", "c", "b", "a"}

	sort.Strings(s2)

	// fmt.Println(s2)

	return s1, s2
}

// F7 function
func F7() []*library.Track {
	var tracks = []*library.Track{
		{
			Title:  "Go",
			Artist: "Delilah",
			Album:  "From the Roots Up",
			Year:   2012,
			Length: library.Length("3m38s"),
		},
		{
			Title:  "Go",
			Artist: "Moby",
			Album:  "Moby",
			Year:   1992,
			Length: library.Length("3m37s"),
		},
		{
			Title:  "Go Ahead",
			Artist: "Alicia Keys",
			Album:  "As I Am",
			Year:   2007,
			Length: library.Length("4m36s"),
		},
		{
			Title:  "Ready 2 Go",
			Artist: "Martin Solveig",
			Album:  "Smash",
			Year:   2011,
			Length: library.Length("4m24s"),
		},
	}

	sort.Sort(library.ByArtist(tracks))

	// library.PrintTracks(tracks)

	sort.Sort(sort.Reverse(library.ByArtist(tracks)))

	// library.PrintTracks(tracks)

	return tracks
}

// F8 function
func F8() []*library.Track {
	var tracks = []*library.Track{
		{
			Title:  "Go",
			Artist: "Delilah",
			Album:  "From the Roots Up",
			Year:   2012,
			Length: library.Length("3m38s"),
		},
		{
			Title:  "Go",
			Artist: "Moby",
			Album:  "Moby",
			Year:   1992,
			Length: library.Length("3m37s"),
		},
		{
			Title:  "Go Ahead",
			Artist: "Alicia Keys",
			Album:  "As I Am",
			Year:   2007,
			Length: library.Length("4m36s"),
		},
		{
			Title:  "Ready 2 Go",
			Artist: "Martin Solveig",
			Album:  "Smash",
			Year:   2011,
			Length: library.Length("4m24s"),
		},
	}

	sort.Sort(library.ByYear(tracks))

	// library.PrintTracks(tracks)

	sort.Sort(sort.Reverse(library.ByYear(tracks)))

	// library.PrintTracks(tracks)

	return tracks
}

// F9 function
func F9() []*library.Track {
	var tracks = []*library.Track{
		{
			Title:  "Go",
			Artist: "Delilah",
			Album:  "From the Roots Up",
			Year:   2012,
			Length: library.Length("3m38s"),
		},
		{
			Title:  "Go",
			Artist: "Moby",
			Album:  "Moby",
			Year:   1992,
			Length: library.Length("3m37s"),
		},
		{
			Title:  "Go Ahead",
			Artist: "Alicia Keys",
			Album:  "As I Am",
			Year:   2007,
			Length: library.Length("4m36s"),
		},
		{
			Title:  "Ready 2 Go",
			Artist: "Martin Solveig",
			Album:  "Smash",
			Year:   2011,
			Length: library.Length("4m24s"),
		},
	}

	sort.Sort(library.CustomStruct{T: tracks, CustomLess: func(x, y *library.Track) bool {
		if x.Title != y.Title {
			return x.Title < y.Title
		}

		if x.Year != y.Year {
			return x.Year < y.Year
		}

		if x.Length != y.Length {
			return x.Length < y.Length
		}

		return false
	}})

	// library.PrintTracks(tracks)

	return tracks
}

// F10 function
func F10() []int {
	values := []int{3, 1, 4, 2}

	// fmt.Println(sort.IntsAreSorted(values))

	sort.Ints(values)

	// fmt.Println(values)
	// fmt.Println(sort.IntsAreSorted(values))

	sort.Sort(sort.Reverse(sort.IntSlice(values)))

	// fmt.Println(values)
	// fmt.Println(sort.IntsAreSorted(values))

	return values
}
