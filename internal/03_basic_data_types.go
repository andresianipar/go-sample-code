package gosamplecode

import (
	"bytes"
	"fmt"
	"strings"
)

// 3 basic data types > strings > strings and byte slices

// Basename1 function
// Basename1 removes directory components and a .suffix.
// e.g., a => a, a.go => a, a/b/c.go => c, a/b.c.go => b.c
func Basename1(s string) string {
	// Discard last'/' and everything before.
	for i := len(s) - 1; i >= 0; i-- {
		if '/' == s[i] {
			s = s[i+1:]
			break
		}
	}

	// Preserve everything before last '.'
	for i := len(s) - 1; i >= 0; i-- {
		if '.' == s[i] {
			s = s[:i]
			break
		}
	}

	return s
}

// Basename2 function
func Basename2(s string) string {
	// -1 if "/" not found
	slash := strings.LastIndex(s, "/")

	s = s[slash+1:]

	if dot := strings.LastIndex(s, "."); 0 <= dot {
		s = s[:dot]
	}

	return s
}

// Comma function
func Comma(s string) string {
	n := len(s)

	if 3 >= n {
		return s
	}

	return Comma(s[:n-3]) + "," + s[n-3:]
}

// IntsToString function
func IntsToString(values []int) string {
	var buff bytes.Buffer

	buff.WriteByte('[')

	for i, v := range values {
		if i > 0 {
			buff.WriteString(", ")
		}

		fmt.Fprintf(&buff, "%d", v)
	}

	buff.WriteByte(']')

	return buff.String()
}
