package basic

import (
	"fmt"
	"math"
)

func IsPrime(num int) {
	if num <= 1 {
		fmt.Println("Num is not a prime number")
	}
	sqrt := int(math.Sqrt(float64(num)))
	for i := 2; i <= sqrt; i++ {
		if num%i == 0 {
			fmt.Println("Num is not a prime number")
			return
		}
	}
	fmt.Println("Num is a prime number")
}
