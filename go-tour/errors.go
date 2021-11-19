package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt struct {
	value float64
}

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("Sqrt error: value can't be negative: %g", e.value)
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt{value: x}
	}
	z := 1.0
	for ok := true; ok; ok = (z*z < x-math.Pow(10, -10) || z*z > x+math.Pow(10, -10)) {
		z -= (z*z - x) / (2 * z)
	}
	return z, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}

/** Output

1.4142135623746899 <nil>
0 Sqrt error: value can't be negative: -2
*/
