package gosamplecode

import (
	"fmt"
	"os"
	"strconv"
	"unicode/utf8"

	library "github.com/andresianipar/go-sample-code/internal"
)

// F9 function
func F9() (string, int, byte, byte) {
	s := "Hello, World!"

	// fmt.Println(s)
	// fmt.Println(len(s))
	// fmt.Printf("%[1]d = %[1]c, %[2]d = %[2]c\n", s[0], s[len(s)-2])

	return s, len(s), s[0], s[len(s)-2]
}

// F10 function
func F10() (string, string, string, string) {
	s := "Hello, World!"

	// fmt.Printf("%s\n%s\n%s\n%s\n", s[0:5], s[:5], s[7:], s[:])

	return s[0:5], s[:5], s[7:], s[:]
}

// F11 function
func F11() string {
	const goUsage = `Go is a tool for managing Go source code.
	
Usage:
    go command [arguments]`

	// fmt.Println(goUsage)

	return goUsage
}

// F12 function
func F12(s, prefix string) bool {
	// fmt.Println(s[:len(prefix)])

	return len(s) >= len(prefix) && prefix == s[:len(prefix)]
}

// F13 function
func F13(s, suffix string) bool {
	// fmt.Println(s[len(s)-len(suffix):])

	return len(s) >= len(suffix) && suffix == s[len(s)-len(suffix):]
}

// F14 function
func F14(s, substring string) bool {
	for i := 0; i < len(s); i++ {
		if F12(s[i:], substring) {
			return true
		}
	}

	return false
}

// F15 function
func F15() (int, int, string, string, string) {
	s := "Hello, 世界！"

	// fmt.Println(len(s))
	// fmt.Println(utf8.RuneCountInString(s))

	// for i := 0; i < len(s); i++ {
	// 	fmt.Printf("%c\n", s[i])
	// }

	// for i := 0; i < len(s); {
	// 	r, size := utf8.DecodeRuneInString(s[i:])

	// 	fmt.Printf("%[1]d\t%[1]c\tsize = %d\n", r, size)

	// 	i += size
	// }

	// for _, r := range s {
	// 	fmt.Printf("%[1]d\t%[1]c\n", r)
	// }

	// fmt.Printf("% x\n", s)

	r := []rune(s)

	// fmt.Printf("%x\n", r)
	// fmt.Println(string(r))
	// fmt.Println(string(0x4e16) + string(0x754c))
	// fmt.Println(string(12345678))

	return len(s), utf8.RuneCountInString(s), string(r), string(0x4e16) + string(0x754c), string(12345678)
}

// F16 function
func F16(s string) string {
	bn := library.Basename1(s)

	// fmt.Println(bn)

	return bn
}

// F17 function
func F17(s string) string {
	bn := library.Basename2(s)

	// fmt.Println(bn)

	return bn
}

// F18 function
func F18(s string) string {
	stringWithCommas := library.Comma(s)

	// fmt.Println(stringWithCommas)

	return stringWithCommas
}

// F19 function
func F19(values []int) string {
	r := library.IntsToString(values)

	// fmt.Println(r)

	return r
}

// F20 function
func F20(x int) (string, string, string) {
	y := fmt.Sprintf("%d", x)
	z := strconv.Itoa(x)
	a := strconv.FormatInt(int64(x), 2)

	// fmt.Printf("%v\n%v\n%v\n", y, z, a)

	return y, z, a
}

// F21 function
func F21() (int, int64, uint64) {
	b, err := strconv.Atoi("1234")

	if nil != err {
		// fmt.Fprintf(os.Stderr, "%v\n", err)

		os.Exit(1)
	}

	c, err := strconv.ParseInt("-1234", 10, 64)

	if nil != err {
		// fmt.Fprintf(os.Stderr, "%v\n", err)

		os.Exit(1)
	}

	d, err := strconv.ParseUint("1234", 10, 64)

	if nil != err {
		// fmt.Fprintf(os.Stderr, "%v\n", err)

		os.Exit(1)
	}

	// fmt.Printf("%v\n%v\n%v\n", b, c, d)

	return b, c, d
}
