package gosamplecode

import (
	"testing"

	goroutinesAndChannels "github.com/andresianipar/go-sample-code/pkg/08_goroutines_and_channels"
)

// 8 goroutines and channels > goroutines
func Test8_F1(t *testing.T) {
	// Test F1 function
	if f1 := goroutinesAndChannels.F1(); 1134903170 != f1 {
		t.Errorf("Should be 1134903170")
	}
}

// 8 goroutines and channels > channels
func Test8_F5_F6_F7_F8(t *testing.T) {
	// Test F5 function
	if f5 := goroutinesAndChannels.F5(); 16 != f5 {
		t.Errorf("Should be 16")
	}

	// Test F6 function
	if f6 := goroutinesAndChannels.F6(); 16 != f6 {
		t.Errorf("Should be 16")
	}

	// Test F7 function
	if f71, f72 := goroutinesAndChannels.F7(); 3 != f71 || 0 != f72 {
		t.Errorf("Should be 3 and 0")
	}
}

// 8 goroutines and channels > example: concurrent directory traversal
func Test8_F12_F13_F14(t *testing.T) {
	// Test F12 function
	if f121, f122 := goroutinesAndChannels.F12(); 0 == f121 || 0 == f122 {
		t.Errorf("Should not be 0 and 0")
	}

	// Test F13 function
	if f131, f132 := goroutinesAndChannels.F13(); 0 == f131 || 0 == f132 {
		t.Errorf("Should not be 0 and 0")
	}

	// Test F14 function
	if f141, f142 := goroutinesAndChannels.F14(); 0 == f141 || 0 == f142 {
		t.Errorf("Should not be 0 and 0")
	}
}
