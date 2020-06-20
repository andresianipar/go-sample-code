package gosamplecode

import (
	"sort"
	"testing"
	"time"

	interfaces "github.com/andresianipar/go-sample-code/pkg/07_interfaces"
)

// 7 interfaces > interfaces as contracts
func Test7_F1(t *testing.T) {
	// Test F1 function
	if f1 := interfaces.F1(); 12 != f1 {
		t.Errorf("Should be 12")
	}
}

// 7 interfaces > interface satisfaction
func Test7_F2(t *testing.T) {
	// Test F2 function
	if f2 := interfaces.F2(); nil == f2 {
		t.Errorf("Should not be nil")
	}
}

// 7 interfaces > parsing flags with flag.Value
func Test7_F3_F4(t *testing.T) {
	// Test F3 function
	if f3 := interfaces.F3(); 1*time.Second != *f3 {
		t.Errorf("Should be 20")
	}

	// Test F4 function
	if f4 := interfaces.F4(); 20 != *f4 {
		t.Errorf("Should be 20")
	}
}

// 7 interfaces > interface values
func Test7_F5(t *testing.T) {
	// Test F5 function
	if f5 := interfaces.F5(); nil == f5 {
		t.Errorf("Should not be nil")
	}
}

// 7 interfaces > sorting with sort.interface
func Test7_F6_F7_F8_F9_F10(t *testing.T) {
	// Test F6 function
	if f61, f62 := interfaces.F6(); "a" != f61[0] || "a" != f62[0] {
		t.Errorf("Should be \"a\" and \"a\"")
	}

	// Test F7 function
	if f7 := interfaces.F7(); 1992 != f7[0].Year {
		t.Errorf("Should be 1992")
	}

	// Test F8 function
	if f8 := interfaces.F8(); 2012 != f8[0].Year {
		t.Errorf("Should be 2012")
	}

	// Test F9 function
	if f9 := interfaces.F9(); 1992 != f9[0].Year {
		t.Errorf("Should be 1992")
	}

	// Test F10 function
	if f10 := interfaces.F10(); 4 != f10[0] || !!sort.IntsAreSorted(f10) {
		t.Errorf("Should be 4 and false")
	}
}

// 7 interfaces > type assertions
func Test7_F14(t *testing.T) {
	// Test F14 function
	if f141, f142 := interfaces.F14(); true != f141 || false != f142 {
		t.Errorf("Should be true and false")
	}
}

// 7 interfaces > discriminating errors with type assertions
func Test7_F15(t *testing.T) {
	// Test F15 function
	if f151, f152 := interfaces.F15(); true != f151 || true != f152 {
		t.Errorf("Should be true and true")
	}
}
