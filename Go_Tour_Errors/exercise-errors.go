package main

import "fmt"

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("***cannot Sqrt negative number: %v***", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		e := ErrNegativeSqrt(x)
		return 0, e
	}

	// z**2 should be as close to x as possible
	z := float64(x)

	for i := 0; i < 10; i++ {
		z = z - (z*z-x)/(2*z)
	}

	return z, nil
}

func exerciseErrors() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-4))
}
