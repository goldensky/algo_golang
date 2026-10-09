package main

import (
	"fmt"
	"math"
)

func isPrime(number int) bool {
	sqrt_from_number := int(math.Sqrt(float64(number)))
	fmt.Println("sqrt_from_number", sqrt_from_number)

	for i := 2; i <= sqrt_from_number; i++ {
		if number%i == 0 {
			fmt.Println("Divided by ", i)
			return false
		}
	}
	return true
}

//number := 132
//result := isPrime(number)
//fmt.Println(result)
