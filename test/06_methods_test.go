package gosamplecode

import (
	"testing"

	library "github.com/andresianipar/go-sample-code/internal"

	methods "github.com/andresianipar/go-sample-code/pkg/06_methods"
)

// 6 methods > method declarations
func Test6_F1(t *testing.T) {
	// Test F1 function
	if f11, f12, f13, f14 := methods.F1(); 5 != f11 || 5 != f12 || 12 != f13 || 12 != f14 {
		t.Errorf("Should be 5, 5, 12, and 12")
	}
}

// 6 methods > methods with a pointer receiver
func Test6_F2_F3(t *testing.T) {
	// Test F2 function
	p := library.Point64{X: 2, Y: 4}

	if f21, f22, f23 := methods.F2(); p != f21 || p != f22 || 2 != f23 {
		t.Errorf("Should be {2 4}, {2 4}, and 2")
	}

	// Test F3 function
	if f3 := methods.F3(); 10 != f3 {
		t.Errorf("Should be 10")
	}
}

// 6 methods > composing types by struct embedding
func Test6_F4_F5(t *testing.T) {
	// Test F4 function
	if f4 := methods.F4(); 10 != f4 {
		t.Errorf("Should be 10")
	}

	// Test F5 function
	if f51, f52 := methods.F5(); 5 != f51 || 0 != f52 {
		t.Errorf("Should be 5 and 0")
	}
}

// 6 methods > method_values_and_expressions
func Test6_F6_F7_F8(t *testing.T) {
	p1 := library.Point64{X: 2, Y: 4}

	// Test F6 function
	if f61, f62 := methods.F6(); 5 != f61 || p1 != f62 {
		t.Errorf("Should be {2 4} and 2")
	}

	// Test F7 function
	if f71, f72 := methods.F7(); 5 != f71 || p1 != f72 {
		t.Errorf("Should be {2 4} and 2")
	}

	// Test F8 function
	p2 := library.Point64{X: 5, Y: 8}

	if f8 := methods.F8(); p2 != f8[0] {
		t.Errorf("Should be {5 8}")
	}
}
