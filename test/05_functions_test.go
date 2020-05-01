package gosamplecode

import (
	"testing"

	functions "github.com/andresianipar/go-sample-code/pkg/05_functions"
)

// 5 functions > recursion
func Test5_F1_F2(t *testing.T) {
	// Test F1 function
	if f1 := functions.F1(); 0 == len(f1) {
		t.Errorf("Should not be 0")
	}

	// Test F2 function
	if f2 := functions.F2(); 0 == len(f2) {
		t.Errorf("Should not be 0")
	}
}

// 5 functions > errors
func Test5_F3(t *testing.T) {
	// Test F3 function
	if f3 := functions.F3(); nil == f3 {
		t.Errorf("Should not be nil")
	}
}

// 5 functions > function values
func Test5_F4_F5(t *testing.T) {
	// Test F4 function
	if f41, f42, f43 := functions.F4(); 16 != f41(4) || nil != f42 || "BCDEF" != f43 {
		t.Errorf("Should be 16, nil, and \"BCDEF\"")
	}

	// Test F5 function
	if f5 := functions.F5(); 0 == len(f5) {
		t.Errorf("Should not be 0")
	}
}

// 5 functions > anonymous functions
func Test5_F6_F7_F8(t *testing.T) {
	// Test F6 function
	if f6 := functions.F6(); 25 != f6 {
		t.Errorf("Should be 25")
	}

	// Test F7 function
	if f7 := functions.F7(); 13 != len(f7) {
		t.Errorf("Should be 13")
	}

	// Test F8 function
	if f8 := functions.F8(); 0 == len(f8) {
		t.Errorf("Should not be 0")
	}
}

// 5 functions > variadic functions
func Test5_F10(t *testing.T) {
	// Test F10 function
	if f10 := functions.F10(); 15 != f10 {
		t.Errorf("Should be 15")
	}
}

// 5 functions > deferred function calls
func Test5_F12_F14_F15(t *testing.T) {
	// Test F12 function
	if f12 := functions.F12(); 0 == len(f12) {
		t.Errorf("Should not be 0")
	}

	// Test F14 function
	if f141, f142 := functions.F14(); 8 != f141 || 12 != f142 {
		t.Errorf("Should be 8 and 12")
	}

	// Test F15 function
	if f151, f152 := functions.F15(); "" == f151 || 0 == f152 {
		t.Errorf("Should not be an empty string and 0")
	}
}

// 5 functions > recover
func Test5_F19(t *testing.T) {
	// Test F19 function
	if f19 := functions.F19(); "Google" != f19 {
		t.Errorf("Should be \"Google\"")
	}
}
