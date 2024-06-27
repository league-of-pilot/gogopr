package udego

import (
	"fmt"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("%v cannot Sqrt negative number: ", float64(e))
}

func sqrtPositive(x float64) float64 {
	var z float64 = 1
	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
	}
	return z
}

func Sqrt20(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	z := sqrtPositive(x)
	return z, nil
}

func Main20() {
	fmt.Println(Sqrt20(2))
	fmt.Println(Sqrt20(-2))
}
