package math

import (
	stdmath "math"
)

// int32

// http://golang.org/src/math/abs.go
func Abs32(x int32) int32 {
	switch {
	case x < 0:
		return -x
	case x == 0:
		return 0 // return correctly abs(-0)
	}
	return x
}

// int64

func Min64(x, y int64) int64 {
	return int64(stdmath.Min(float64(x), float64(y)))
}

func Max64(x, y int64) int64 {
	return int64(stdmath.Max(float64(x), float64(y)))
}

func Abs64(x int64) int64 {
	switch {
	case x < 0:
		return -x
	case x == 0:
		return 0 // return correctly abs(-0)
	}
	return x
}
