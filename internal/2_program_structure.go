package gosamplecode

import (
	"fmt"
)

// Celsius type
type Celsius float64

// Fahrenheit type
type Fahrenheit float64

// Kelvin type
type Kelvin float64

const (
	// AbsoluteZeroC represents absolute zero
	AbsoluteZeroC Celsius = -273.5
	// FreezingC is the freezing point of water
	FreezingC Celsius = 0
	// BoilingC is the boiling point of water
	BoilingC Celsius = 100
)

// 2 program structure > type declarations

// String function
func (c Celsius) String() string {
	return fmt.Sprintf("%g°C", c)
}

// 2 program structure > packages and files

// String function
func (f Fahrenheit) String() string {
	return fmt.Sprintf("%g°F", f)
}

// String function
func (k Kelvin) String() string {
	return fmt.Sprintf("%g°K", k)
}

// 2 program structure > packages and files > package initialization

// I variable
var I int

func init() {
	I = 10
}
