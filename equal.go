package mathx

import "math"

// DefaultPrecision is intended for when comparing numbers that are close to 0 and 1 but want to maintain precision decent.
const DefaultPrecision = 1e-9

// Equal will compare two floats using DefaultPrecision and return a boolean.
func Equal(x, y float64) bool {
	return EqualP(x, y, DefaultPrecision)
}

// EqualP will compare two floats using the specified precision and return a boolean.
func EqualP(x, y, p float64) bool {
	return math.Abs(x-y) <= p
}
