package gosamplecode

import (
	"fmt"

	library "github.com/andresianipar/go-sample-code/internal"
)

// F6 function
func F6() (string, string, string) {
	months := [...]string{1: "January", 2: "February", 3: "March", 4: "April", 5: "May", 6: "June", 7: "July", 8: "August", 9: "September", 10: "October", 11: "November", 12: "December"}
	Q2 := months[4:7]
	summer := months[6:9]

	// fmt.Println(Q2)
	// fmt.Println(summer)

	endlessSummer := summer[:5]

	// fmt.Println(endlessSummer)

	return fmt.Sprintf("%s", Q2), fmt.Sprintf("%s", summer), fmt.Sprintf("%s", endlessSummer)
}

// F7 function
func F7() [6]int {
	a := [...]int{0, 1, 2, 3, 4, 5}

	library.Reverse(a[:])

	// fmt.Println(a)

	return a
}

// F8 function
func F8() (int, int) {
	var s1 []int

	s1 = nil
	s1 = []int(nil)

	var s2 []int

	s2 = []int{}

	// fmt.Println(s1)
	// fmt.Println(s2)

	return len(s1), len(s2)
}

// F9 function
func F9() (string, []int) {
	var runes []rune

	for _, r := range "Hello, World!" {
		runes = append(runes, r)
	}

	// fmt.Printf("%q\n", runes)

	var x []int

	x = append(x, 1)
	x = append(x, 2, 3, 4, 5)
	x = append(x, x...)

	// fmt.Println(x)

	return string(runes), x
}

// F10 function
func F10() string {
	data := []string{"one", "", "three", "", "five"}

	// fmt.Printf("%q\n", library.Nonempty1(data))
	// fmt.Printf("%q\n", data)

	data = []string{"one", "", "three", "", "five"}
	// data = library.Nonempty1(data)

	// fmt.Printf("%q\n", data)

	data = []string{"one", "", "three", "", "five"}
	data = library.Nonempty2(data)

	// fmt.Printf("%q\n", data)

	return fmt.Sprintf("%s", data)
}

// F11 function
func F11() []int {
	var stack []int

	stack = append(stack, 1)
	stack = append(stack, 2)
	stack = append(stack, 3)
	stack = append(stack, 4)

	// top := stack[len(stack)-1]

	// fmt.Println(top)
	// fmt.Println(stack)

	stack = stack[:len(stack)-1]

	// fmt.Println(stack)

	return stack
}

// F12 function
func F12() []int {
	s := []int{1, 2, 3, 4, 5}

	s = library.Remove(s, 2)

	// fmt.Println(s)

	return s
}
