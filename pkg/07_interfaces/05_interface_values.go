package gosamplecode

import (
	"bytes"

	library "github.com/andresianipar/go-sample-code/internal"
)

// F5 function
func F5() *bytes.Buffer {
	const debug = true
	// const debug = false
	var buff *bytes.Buffer

	if debug {
		buff = new(bytes.Buffer)
	}

	library.OutWrite(buff)

	// fmt.Println(buff)

	return buff
}
