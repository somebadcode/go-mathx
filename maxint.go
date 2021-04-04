package mathx

// MaxInt will return the largest of the two integers.
func MaxInt(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

// MaxUint will return the largest of the two unsigned integers.
func MaxUint(a, b uint) uint {
	if a >= b {
		return a
	}
	return b
}

// MaxInt64 will return the largest of the two integers.
func MaxInt64(a, b int64) int64 {
	if a >= b {
		return a
	}
	return b
}

// MaxUint64 will return the largest of the two unsigned integers.
func MaxUint64(a, b uint64) uint64 {
	if a >= b {
		return a
	}
	return b
}