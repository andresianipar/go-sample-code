package gosamplecode

import (
	"testing"
	"time"

	basicDataTypes "github.com/andresianipar/go-sample-code/pkg/03_basic_data_types"
)

// 3 basic data types > integers
func Test3_F1_F2_F3_F4(t *testing.T) {
	// Test F1 function
	if f11, f12 := basicDataTypes.F1(); 18446744073709551615 != f11 || 0 != f12 {
		t.Errorf("Should be 18446744073709551615 and 0")
	}

	// Test F2 function
	reversed := [3]string{"bronze", "silver", "gold"}

	if f2 := basicDataTypes.F2(); reversed != f2 {
		t.Errorf("Should be \"bronze, silver, gold\"")
	}

	// Test F3 function
	if f31, f32 := basicDataTypes.F3(); 3 != f31 || 3 != f32 {
		t.Errorf("Should be 3 and 3")
	}

	// Test F4 function
	if f41, f42 := basicDataTypes.F4(); 97 != int(f41) || 10 != int(f42) {
		t.Errorf("Should be 97 and 10")
	}
}

// 3 basic data types > floating-point numbers
func Test3_F5(t *testing.T) {
	// Test F5 function
	if f51, f52, _, _, _ := basicDataTypes.F5(); 0 != f51 || -0 != f52 {
		t.Errorf("Should be 0 and -0")
	}
}

// 3 basic data types > complex numbers
func Test3_F6(t *testing.T) {
	// Test F6 function
	if f61, f62 := basicDataTypes.F6(); -5 != f61 || 10 != f62 {
		t.Errorf("Should be -5 and 10")
	}
}

// 3 basic data types > booleans
func Test3_F7_F8(t *testing.T) {
	// Test F7 function
	if f71, f72 := basicDataTypes.F7(); true != f71 || true != f72 {
		t.Errorf("Should be true and true")
	}

	// Test F8 function
	if f8 := basicDataTypes.F8(1); true != f8 {
		t.Errorf("Should be true")
	}
}

// 3 basic data types > strings
func Test3_F9_F10_F11_F12_F13_F14_F15_F16_F17_F18_F19_F20_F21(t *testing.T) {
	// Test F9 function
	if f91, f92, f93, f94 := basicDataTypes.F9(); "Hello, World!" != f91 || 13 != f92 || 72 != f93 || 100 != f94 {
		t.Errorf("Should be \"Hello, World!\", 13, 72, and 100")
	}

	// Test F10 function
	if f101, f102, f103, f104 := basicDataTypes.F10(); "Hello" != f101 || "Hello" != f102 || "World!" != f103 || "Hello, World!" != f104 {
		t.Errorf("Should be \"Hello\", \"Hello\", \"World!\", and \"Hello, World!\"")
	}

	// Test F11 function
	if f11 := basicDataTypes.F11(); "" == f11 {
		t.Errorf("Should not be an empty string")
	}

	// Test F12 function
	if f12 := basicDataTypes.F12("Hello, World!", "Hello"); true != f12 {
		t.Errorf("Should be true")
	}

	// Test F13 function
	if f13 := basicDataTypes.F13("Hello, World!", "World!"); true != f13 {
		t.Errorf("Should be true")
	}

	// Test F14 function
	if f14 := basicDataTypes.F14("Hello, World!", "o, "); true != f14 {
		t.Errorf("Should be true")
	}

	// Test F15 function
	if f151, f152, f153, f154, f155 := basicDataTypes.F15(); 16 != f151 || 10 != f152 || "Hello, 世界！" != f153 || "世界" != f154 || "�" != f155 {
		t.Errorf("Should be 16, 10, \"Hello, 世界！\", \"世界\", and \"�\"")
	}

	// Test F16 function
	if f16 := basicDataTypes.F16("a/b/c.go"); "c" != f16 {
		t.Errorf("Should be \"c\"")
	}

	// Test F17 function
	if f17 := basicDataTypes.F17("c.d.go"); "c.d" != f17 {
		t.Errorf("Should be \"c.d\"")
	}

	// Test F18 function
	if f18 := basicDataTypes.F18("123456"); "123,456" != f18 {
		t.Errorf("Should be \"123,456\"")
	}

	// Test F19 function
	if f19 := basicDataTypes.F19([]int{1, 2, 3, 4}); "[1, 2, 3, 4]" != f19 {
		t.Errorf("Should be \"[1, 2, 3, 4]\"")
	}

	// Test F20 function
	if f201, f202, f203 := basicDataTypes.F20(10); "10" != f201 || "10" != f202 || "1010" != f203 {
		t.Errorf("Should be \"10\", \"10\", and \"1010\"")
	}

	// Test F21 function
	if f211, f212, f213 := basicDataTypes.F21(); 1234 != f211 || -1234 != f212 || 1234 != f213 {
		t.Errorf("Should be 1234, -1234, and 1234")
	}
}

// 3 basic data types > constants
func Test3_F22_F23_F24_F25(t *testing.T) {
	// Test F22 function
	const a time.Duration = 0
	const b = 5 * time.Minute

	if f221, f222 := basicDataTypes.F22(); a != f221 || b != f222 {
		t.Errorf("Should be time.Duration 0s and time.Duration 5m0s")
	}

	// Test F23 function
	if f231, f232, f233, f234 := basicDataTypes.F23(); 1 != f231 || 1 != f232 || 2 != f233 || 2 != f234 {
		t.Errorf("Should be 1, 1, 2, and 2")
	}

	// Test F24 function
	if f241, f242, f243, f244 := basicDataTypes.F24(); 0 != f241 || 5 != f242 || 10 != f243 || 15 != f244 {
		t.Errorf("Should be 1, 1, 2, and 2")
	}

	// Test F25 function
	if f251, f252, f253, f254 := basicDataTypes.F25(); 1024 != f251 || 1048576 != f252 || 1073741824 != f253 || 1099511627776 != f254 {
		t.Errorf("Should be 1, 1, 2, and 2")
	}
}
