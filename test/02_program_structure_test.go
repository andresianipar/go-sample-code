package gosamplecode

import (
	"testing"

	programStructure "github.com/andresianipar/go-sample-code/pkg/02_program_structure"

	library "github.com/andresianipar/go-sample-code/internal"
)

// 2 program structure > names
func Test2_F1(t *testing.T) {
	// Test F1 function
	f1 := programStructure.F1()

	if true != f1 {
		t.Errorf("Should be true")
	}
}

// 2 program structure > declarations
func Test2_F2_F3(t *testing.T) {
	// Test F2 function
	f2F, f2C := programStructure.F2()

	if 212 != f2F || 100 != f2C {
		t.Errorf("Should be 212 and 100")
	}

	// Test F3 function
	f3FP := programStructure.F3(32)
	f3BP := programStructure.F3(212)

	if 0 != f3FP || 100 != f3BP {
		t.Errorf("Should be 0 and 100")
	}
}

// 2 program structure > variables
func Test2_F4_F5_F6_F7_F8_F9_F10_F11_F12(t *testing.T) {
	// Test F4 function
	f4B, f4F, f4S := programStructure.F4()

	if true != f4B || 4.5 != f4F || "yes" != f4S {
		t.Errorf("Should be true, 4.5, and \"yes\"")
	}

	// Test F5 function
	f5 := programStructure.F5()

	if 0 == *f5 {
		t.Errorf("Should not be 0")
	}

	// Test F6 function
	f61 := programStructure.F6()
	f62 := programStructure.F6()

	if f61 == f62 {
		t.Errorf("Should have different values")
	}

	if *f61 != *f62 {
		t.Errorf("Should have the same value")
	}

	// Test F7 function
	v := 1
	f7 := programStructure.F7(&v)

	if 2 != f7 {
		t.Errorf("Should be 2")
	}

	// Test F8 function
	f8 := programStructure.F8()

	if "" == f8 {
		t.Errorf("Should not be empty string")
	}

	// Test F9 function
	f9 := programStructure.F9()

	if 1 != *f9 {
		t.Errorf("Should be 1")
	}

	// Test F10 function
	f101, f102 := programStructure.F10()

	if 0 != *f101 || 0 != *f102 {
		t.Errorf("Should be 0")
	}

	// Test F11 function
	f11 := programStructure.F11(1, 2)

	if 1 != f11 {
		t.Errorf("Should be 1")
	}

	// Test F12 function
	programStructure.F12()

	if 1 != *programStructure.Global {
		t.Errorf("Should be 1")
	}
}

// 2 program structure > assignments
func Test2_F13_F14_F15_F16(t *testing.T) {
	// Test F13 function
	f13 := programStructure.F13()

	if 5 != f13 {
		t.Errorf("Should be 5")
	}

	// Test F14 function
	f141, f142 := programStructure.F14()

	if 2 != f141 || 1 != f142 {
		t.Errorf("Should be 2 and 1")
	}

	// Test F15 function
	f15 := programStructure.F15(5, 10)

	if 5 != f15 {
		t.Errorf("Should be 5")
	}

	// Test F16 function
	f16 := programStructure.F16(5)

	if 5 != f16 {
		t.Errorf("Should be 5")
	}
}

// 2 program structure > type declarations
func Test2_F17_F18_F19_F20(t *testing.T) {
	// Test F17 function
	var c library.Celsius = 36.5
	f17 := programStructure.F17(c)

	if 97.7 != f17 {
		t.Errorf("Should be 97.7")
	}

	// Test F18 function
	var f library.Fahrenheit = 97.7
	f18 := programStructure.F18(f)

	if 36.5 != f18 {
		t.Errorf("Should be 97.7")
	}

	// Test F19 function
	f19C, f19F, f19B := programStructure.F19()

	if 100 != f19C || 180 != f19F || true == f19B {
		t.Errorf("Should be 100, 180, and true")
	}

	// Test F20 function
	f20 := programStructure.F20()

	if "100°C" != f20 {
		t.Errorf("Should be 100°C")
	}
}

// 2 program structure > packages and files
func Test2_F21_F22_F23(t *testing.T) {
	// Test F21 function
	var c library.Celsius = 0
	f21 := programStructure.F21(c)

	if 273.15 != f21 {
		t.Errorf("Should be 273.15")
	}

	// Test F22 function
	var k library.Kelvin = 0
	f22 := programStructure.F22(k)

	if -273.15 != f22 {
		t.Errorf("Should be -273.15")
	}

	// Test F23 function
	f23 := programStructure.F23()

	if 10 != f23 {
		t.Errorf("Should be 10")
	}
}

// 2 program structure > scope
func Test2_F24(t *testing.T) {
	// Test F24 function
	f24 := programStructure.F24()

	if "Hello" != f24 {
		t.Errorf("Should be \"Hello\"")
	}
}
