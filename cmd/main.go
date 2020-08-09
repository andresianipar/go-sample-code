package main

// library "github.com/andresianipar/go-sample-code/internal"
// tutorial "github.com/andresianipar/go-sample-code/pkg/01_tutorial"
// programStructure "github.com/andresianipar/go-sample-code/pkg/02_program_structure"
// basicDataTypes "github.com/andresianipar/go-sample-code/pkg/03_basic_data_types"
// compositeTypes "github.com/andresianipar/go-sample-code/pkg/04_composite_types"
// functions "github.com/andresianipar/go-sample-code/pkg/05_functions"
// methods "github.com/andresianipar/go-sample-code/pkg/06_methods"
// interfaces "github.com/andresianipar/go-sample-code/pkg/07_interfaces"
// goroutinesAndChannels "github.com/andresianipar/go-sample-code/pkg/08_goroutines_and_channels"

func f1() {
	// 1 tutorial > command-line arguments
	// tutorial.F1()
	// tutorial.F2()
	// tutorial.F3()

	// 1 tutorial > finding duplicate lines
	// tutorial.F4()
	// tutorial.F5()
	// tutorial.F6()

	// 1 tutorial > fetching a url
	// tutorial.F7()

	// 1 tutorial > fetching urls concurrently
	// tutorial.F8()

	// 1 tutorial > a web server
	// tutorial.F9()
	// tutorial.F10()
	// tutorial.F11()

	// 1 tutorial > loose ends
	// tutorial.F12("heads")
	// tutorial.F13(5)
}

func f2() {
	// 2 program structure > names
	// programStructure.F1()

	// 2 program structure > declarations
	// programStructure.F2()
	// programStructure.F3(32)

	// 2 program structure > variables
	// programStructure.F4()

	// 2 program structure > variables > pointers
	// programStructure.F5()
	// programStructure.F6()

	// v := 1

	// programStructure.F7(&v)
	// programStructure.F8()

	// 2 program structure > variables > the new function
	// programStructure.F9()
	// programStructure.F10()
	// programStructure.F11(1, 2)

	// 2 program structure > variables > lifetime of variables
	// programStructure.F12()

	// fmt.Println(*programStructure.Global)

	// 2 program structure > assignments
	// programStructure.F13()

	// 2 program structure > assignments > tuple assigment
	// programStructure.F14()
	// programStructure.F15(5, 10)
	// programStructure.F16(5)

	// 2 program structure > type declarations
	// var c1 library.Celsius = 36.5

	// programStructure.F17(c1)

	// var f library.Fahrenheit = 97.7

	// programStructure.F18(f)
	// programStructure.F19()
	// programStructure.F20()

	// 2 program structure > packages and files
	// var c2 library.Celsius = 0

	// programStructure.F21(c2)

	// var k library.Kelvin = 0

	// programStructure.F22(k)

	// 2 program structure > packages and files > package initialization
	// programStructure.F23()

	// 2 program structure > scope
	// programStructure.F24()
	// programStructure.F25()
}

func f3() {
	// 3 basic data types > integers
	// basicDataTypes.F1()
	// basicDataTypes.F2()
	// basicDataTypes.F3()
	// basicDataTypes.F4()

	// 3 basic data types > floating-point numbers
	// basicDataTypes.F5()

	// 3 basic data types > complex numbers
	// basicDataTypes.F6()

	// 3 basic data types > booleans
	// basicDataTypes.F7()
	// basicDataTypes.F8(1)

	// 3 basic data types > strings
	// basicDataTypes.F9()
	// basicDataTypes.F10()

	// 3 basic data types > strings > string literals
	// basicDataTypes.F11()

	// 3 basic data types > strings > UTF-8
	// basicDataTypes.F12("Hello, World!", "Hello")
	// basicDataTypes.F13("Hello, World!", "World!")
	// basicDataTypes.F14("Hello, World!", "o, ")
	// basicDataTypes.F15()

	// 3 basic data types > strings > strings and byte slices
	// basicDataTypes.F16("a/b/c.go")
	// basicDataTypes.F17("c.d.go")
	// basicDataTypes.F18("123456")
	// basicDataTypes.F19([]int{1, 2, 3, 4})

	// 3 basic data types > strings > conversions between strings and numbers
	// basicDataTypes.F20(10)
	// basicDataTypes.F21()

	// 3 basic data types > constants
	// basicDataTypes.F22()
	// basicDataTypes.F23()

	// 3 basic data types > constants > the constant generator iota
	// basicDataTypes.F24()
	// basicDataTypes.F25()
}

func f4() {
	// 4 composite types > arrays
	// compositeTypes.F1()
	// compositeTypes.F2()
	// compositeTypes.F3()
	// compositeTypes.F4()
	// compositeTypes.F5()

	// 4 composite types > slices
	// compositeTypes.F6()
	// compositeTypes.F7()
	// compositeTypes.F8()

	// 4 composite types > slices > the append function
	// compositeTypes.F9()

	// 4 composite types > slices > in-place slice techniques
	// compositeTypes.F10()
	// compositeTypes.F11()
	// compositeTypes.F12()

	// 4 composite types > maps
	// compositeTypes.F13()
	// compositeTypes.F14()
	// compositeTypes.F15()
	// compositeTypes.F16()
	// compositeTypes.F17()
	// compositeTypes.F18()

	// 4 composite types > structs
	// compositeTypes.F19()

	// 4 composite types > structs > struct literals
	// compositeTypes.F20()
	// compositeTypes.F21(compositeTypes.F20())
	// compositeTypes.F22(compositeTypes.F20(), "Principal Software Engineer", 10000)

	// 4 composite types > structs > comparing structs
	// compositeTypes.F23()
	// compositeTypes.F24()

	// 4 composite types > structs > struct embedding and anonymous fields
	// compositeTypes.F25()
	// compositeTypes.F26()

	// 4 composite types > json
	// compositeTypes.F27()
	// compositeTypes.F28()
	// compositeTypes.F29()

	// 4 composite types > text and html templates
	// compositeTypes.F30()
	// compositeTypes.F31()
	// compositeTypes.F32()
}

func f5() {
	// 5 functions > recursion
	// functions.F1()
	// functions.F2()

	// 5 functions > errors > error-handling strategies
	// functions.F3()

	// 5 functions > function values
	// functions.F4()
	// functions.F5()

	// 5 functions > anonymous functions
	// functions.F6()
	// functions.F7()
	// functions.F8()
	// functions.F9()

	// 5 functions > variadic functions
	// functions.F10()
	// functions.F11()

	// 5 functions > deferred function calls
	// functions.F12()
	// functions.F13()
	// functions.F14()
	// functions.F15()

	// 5 functions > panic
	// functions.F16()
	// functions.F17()
	// functions.F18()

	// 5 functions > recover
	// functions.F19()
}

func f6() {
	// 6 methods > method declarations
	// methods.F1()

	// 6 methods > methods with a pointer receiver
	// methods.F2()

	// 6 methods > methods with a pointer receiver > nil is a valid receiver value
	// methods.F3()

	// 6 methods > composing types by struct embedding
	// methods.F4()
	// methods.F5()

	// 6 methods > method values and expressions
	// methods.F6()
	// methods.F7()
	// methods.F8()
}

func f7() {
	// 7 interfaces > interfaces as contracts
	// interfaces.F1()

	// 7 interfaces > interface satisfaction
	// interfaces.F2()

	// 7 interfaces > parsing flags with flag.Value
	// interfaces.F3()
	// interfaces.F4()

	// 7 interfaces > interface values
	// interfaces.F5()

	// 7 interfaces > sorting with sort.interface
	// interfaces.F6()
	// interfaces.F7()
	// interfaces.F8()
	// interfaces.F9()
	// interfaces.F10()

	// 7 interfaces > the http.handler interface
	// interfaces.F11()
	// interfaces.F12()
	// interfaces.F13()

	// 7 interfaces > type assertions
	// interfaces.F14()

	// 7 interfaces > discriminating errors with type assertions
	// interfaces.F15()

	// 7 interfaces > querying behaviours with interface type assertions
	// interfaces.F16()

	// 7 interfaces > type switches
	// interfaces.F17()
}

func f8() {
	// 8 goroutines and channels > goroutines
	// goroutinesAndChannels.F1()

	// 8 goroutines and channels > example: concurrent clock server
	// goroutinesAndChannels.F2()

	// 8 goroutines and channels > example: concurrent echo server
	// goroutinesAndChannels.F3()

	// 8 goroutines and channels > channels > unbuffered channels
	// goroutinesAndChannels.F4()

	// 8 goroutines and channels > channels > pipelines
	// goroutinesAndChannels.F5()

	// 8 goroutines and channels > channels > unidirectional channel types
	// goroutinesAndChannels.F6()

	// 8 goroutines and channels > channels > buffered channels
	// goroutinesAndChannels.F7()

	// 8 goroutines and channels > example: concurrent web crawler
	// goroutinesAndChannels.F8()

	// 8 goroutines and channels > multiplexing with select
	// goroutinesAndChannels.F9()
	// goroutinesAndChannels.F10()
	// goroutinesAndChannels.F11()

	// 8 goroutines and channels > example: concurrent directory traversal
	// goroutinesAndChannels.F12()
	// goroutinesAndChannels.F13()
	// goroutinesAndChannels.F14()

	// 8 goroutines and channels > cancellation
	// goroutinesAndChannels.F15()

	// 8 goroutines and channels > example chat server
	// goroutinesAndChannels.F16()
}

func main() {
	// 1 tutorial
	// f1()

	// 2 program structure
	// f2()

	// 3 basic data types
	// f3()

	// 4 composite types
	// f4()

	// 5 functions
	// f5()

	// 6 methods
	// f6()

	// 7 interfaces
	// f7()

	// 8 goroutines and channels
	// f8()
}
