package main

import (
	"fmt"
	"math"
)

/*
func Sqrt(x float64) float64 {
	z := 1.0
	prevZ := 0.0
	stop := false
	for ok := true; ok; ok = stop {
    	z -= (z*z - x) / (2*z)
		stop = z != prevZ
		prevZ = z
		fmt.Println(z)
	}
	return z;
}*/

func Sqrt(x float64) float64 {
	z := 1.0
	for ok := true; ok; ok = (z*z < x-math.Pow(10, -10) || z*z > x+math.Pow(10, -10)) {
		z -= (z*z - x) / (2 * z)
		fmt.Println(z)
	}
	return z
}

/*
func Sqrt(x float64) float64 {
	z := 1.0
	for i := 0; i < 10; i++ {
    	z -= (z*z - x) / (2*z)
		fmt.Println(z)
	}
	return z;
}*/

func main() {
	fmt.Println(Sqrt(2))
	// fmt.Println(toto(3))
}

/** Output

1.5
1.4166666666666667
1.4142156862745099
1.4142135623746899
1.4142135623746899
*/
