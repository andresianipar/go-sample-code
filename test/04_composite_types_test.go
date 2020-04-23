package gosamplecode

import (
	"testing"

	compositeTypes "github.com/andresianipar/go-sample-code/pkg/04_composite_types"
)

// 4 composite types > arrays
func Test4_F1_F2_F3_F4_F5(t *testing.T) {
	// Test F1 function
	if f11, f12 := compositeTypes.F1(); 0 != f11 || 3 != f12 {
		t.Errorf("Should be 0 and 3")
	}

	// Test F2 function
	if f2 := compositeTypes.F2(); "$" != f2 {
		t.Errorf("Should be \"$\"")
	}

	// Test F3 function
	if f31, f32 := compositeTypes.F3(); 100 != f31 || -1 != f32 {
		t.Errorf("Should be 100 and -1")
	}

	// Test F4 function
	if f41, f42 := compositeTypes.F4(); false != f41 || true != f42 {
		t.Errorf("Should be false and true")
	}

	// Test F5 function
	if f5 := compositeTypes.F5(); false != f5 {
		t.Errorf("Should be false")
	}
}

// 4 composite types > slices
func Test4_F6_F7_F8_F9_F10_F11_F12(t *testing.T) {
	// Test F6 function
	if f61, f62, f63 := compositeTypes.F6(); "[April May June]" != f61 || "[June July August]" != f62 || "[June July August September October]" != f63 {
		t.Errorf("Should be \"[April May June]\", \"[June July August]\", and \"[June July August September October]\"")
	}

	// Test F7 function
	if f7 := compositeTypes.F7(); [...]int{5, 4, 3, 2, 1, 0} != f7 {
		t.Errorf("Should be [5 4 3 2 1 0]")
	}

	// Test F8 function
	if f81, f82 := compositeTypes.F8(); 0 != f81 || 0 != f82 {
		t.Errorf("Should be 0 and 0")
	}

	// Test F9 function
	if f91, f92 := compositeTypes.F9(); "Hello, World!" != f91 || nil == f92 {
		t.Errorf("Should be \"Hello, World!\"")
		t.Errorf("Should not be nil")
	}

	// Test F10 function
	if f10 := compositeTypes.F10(); "[one three five]" != f10 {
		t.Errorf("Should be \"[one three five]\"")
	}

	// Test F11 function
	if f11 := compositeTypes.F11(); nil == f11 {
		t.Errorf("Should not be nil")
	}

	// Test F12 function
	if f12 := compositeTypes.F12(); nil == f12 {
		t.Errorf("Should not be nil")
	}
}

// 4 composite types > maps
func Test4_F13_F14_F15_F16_F17_F18(t *testing.T) {
	// Test F13 function
	if f131, f132 := compositeTypes.F13(); 31 != f131["alice"] || 0 != f132["charlie"] {
		t.Errorf("Should not be 31 and 0")
	}

	// Test F14 function
	if f14 := compositeTypes.F14(); "[a b c f m t z]" != f14 {
		t.Errorf("Should be \"[a b c f m t z]\"")
	}

	// Test F15 function
	if f151, f152, f153 := compositeTypes.F15(); 31 != f151 || true != f152 || false != f153 {
		t.Errorf("Should be 31, true, and false")
	}

	// Test F16 function
	if f16 := compositeTypes.F16(); nil == f16 {
		t.Errorf("Should not be nil")
	}

	// Test F17 function
	if f171, f172 := compositeTypes.F17(); 1 != f171 || 1 != f172 {
		t.Errorf("Should be 1 and 1")
	}

	// Test F18 function
	if f18 := compositeTypes.F18(); true != f18 {
		t.Errorf("Should be true")
	}
}

// 4 composite types > structs
func Test4_F19_F20_F21_F22_F23_F24_F25_F26(t *testing.T) {
	// Test F19 function
	if f19 := compositeTypes.F19(); "Johnny Bravo" != f19.Name {
		t.Errorf("Should be \"Johnny Bravo\"")
	}

	// Test F20 function
	if f20 := compositeTypes.F20(); "Johnny Bravo" != f20.Name {
		t.Errorf("Should be \"Johnny Bravo\"")
	}

	// Test F21 function
	if f21 := compositeTypes.F21(compositeTypes.F20()); "Principal Software Engineer" != f21.Position || 10000 != f21.Salary {
		t.Errorf("Should be \"Principal Software Engineer\" and 10000")
	}

	// Test F22 function
	if f22 := compositeTypes.F22(compositeTypes.F20(), "Principal Software Engineer", 10000); "Principal Software Engineer" != (*f22).Position || 10000 != (*f22).Salary {
		t.Errorf("Should be \"Principal Software Engineer\" and 10000")
	}

	// Test F23 function
	if f231, f232 := compositeTypes.F23(); false != (f231 == f232) {
		t.Errorf("Should be false")
	}

	// Test F24 function
	if f24 := compositeTypes.F24(); 1 != f24 {
		t.Errorf("Should be 1")
	}

	// Test F25 function
	if f251, f252 := compositeTypes.F25(); true == ((f251.Circle.Center.X != f252.X) || (f251.Circle.Center.Y != f252.Y)) {
		t.Errorf("Should be true")
	}

	// Test F26 function
	if f261, f262 := compositeTypes.F26(); true == ((f261.X != f262.X) || (f261.Y != f262.Y)) {
		t.Errorf("Should be true")
	}
}

// 4 composite types > json
func Test4_F27_F28_F29(t *testing.T) {
	// Test F27 function
	if f27 := compositeTypes.F27(); nil == f27 {
		t.Errorf("Should not be nil")
	}

	// Test F28 function
	if f28 := compositeTypes.F28(); 3 != len(f28) {
		t.Errorf("Should be 3")
	}

	// Test F29 function
	if f29 := compositeTypes.F29(); nil == f29 {
		t.Errorf("Should not be nil")
	}
}
