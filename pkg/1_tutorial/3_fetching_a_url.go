package gosamplecode

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

// F7 function
func F7() map[string]interface{} {
	url := "http://dummy.restapiexample.com/api/v1/employees"
	resp, err := http.Get(url)

	if nil != err {
		fmt.Fprintf(os.Stderr, "%v\n", err)

		os.Exit(1)
	}

	body, err := ioutil.ReadAll(resp.Body)

	resp.Body.Close()

	if nil != err {
		fmt.Fprintf(os.Stderr, "%v\n", err)

		os.Exit(1)
	}

	var result map[string]interface{}

	json.Unmarshal([]byte(body), &result)

	// fmt.Println(resp.Status + "\n")

	// for key, value := range result {
	// 	if "data" == key {
	// 		fmt.Printf("%s\n", value)
	// 	}
	// }

	return result
}
