package lib

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func Max[T int | int16 | int32 | int64 | float64 | float32](x T, y T) T {
	if x > y {
		return x
	}
	return y
}

func RandomInt(maxInt int) int {
	return rand.Intn(maxInt)
}
