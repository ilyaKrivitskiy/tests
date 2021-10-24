package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	if e < 0 {
		return fmt.Sprintf("cannot Sqrt negative number: %f", float64(e))
	} else {
		return ""
	}
}

func Sqrt(x float64) (float64, error) {
	z := 1.0
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	} else if x == 0 {
		return 0, nil
	} else {
		for math.Abs(z*z-x) > 0.00000001 {
			z -= (z*z - x) / (2 * z)
		}
		return z, nil
	}
}

func main() {
	two, negTwo := 2, -2
	res1, err := Sqrt(float64(two))
	if err != nil {
		panic(err)
	}
	fmt.Printf("Sqrt of %f = %f \n", float64(two), res1)

	res2, err2 := Sqrt(float64(negTwo))
	if err2 != nil {
		panic(err2)
	}
	fmt.Printf("Sqrt of %f = %f", float64(negTwo), res2)

}
