package gosamplecode

import (
	"encoding/json"
	"log"

	library "github.com/andresianipar/go-sample-code/internal"
)

// F27 function
func F27() []uint8 {
	data1, err := json.Marshal(library.Movies)

	if nil != err {
		log.Fatalf("JSON marshalling failed: %s", err)
	}

	// fmt.Printf("%s\n", data1)

	// data2, err := json.MarshalIndent(library.Movies, "", "  ")

	// if nil != err {
	// 	log.Fatalf("JSON marshalling failed: %s", err)
	// }

	// fmt.Printf("%s\n", data2)

	return data1
}

// F28 function
func F28() []struct{ Title string } {
	var titles []struct{ Title string }
	data := "[{\"title\":\"Casablanca\",\"released\":1942,\"actors\":[\"Humphrey Bogart\",\"Ingrid Bergman\"]},{\"title\":\"Cool Hand Luke\",\"released\":1967,\"color\":true,\"actors\":[\"Paul Newman\"]},{\"title\":\"Bullitt\",\"released\":1968,\"color\":true,\"actors\":[\"Steve McQueen\",\"Jacqueline Bisset\"]}]"

	if err := json.Unmarshal([]byte(data), &titles); nil != err {
		log.Fatalf("JSON unmarshalling failed: %s", err)
	}

	// fmt.Println(titles)

	return titles
}

// F29 function
func F29() *library.IssuesSearchResult {
	result, err := library.SearchIssues([]string{"repo:golang/go", "is:open", "generic"})

	if nil != err {
		log.Fatal(err)
	}

	// fmt.Printf("%d issues:\n", result.TotalCount)

	// for _, item := range result.Items {
	// 	fmt.Printf("#%-5d %9.9s %.55s\n", item.Number, item.User.Login, item.Title)
	// }

	return result
}
