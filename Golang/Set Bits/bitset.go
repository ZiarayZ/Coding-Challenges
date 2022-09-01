package main

import (
	"fmt"
	"math"
)

func countSetBits(n int) int {
	value := 0
	//calculate the exponent (x) of 2 that fits into n: 2^x = n, then add 1
	iterat := int(math.Log(float64(n))/math.Log(2)) + 1
	//iterate through every possible bit that could be set
	for i := 0; i < iterat; i++ {
		value += (n >> i) & 1
	}
	return value
}

func main() {
	val := 31
	fmt.Println("There are", countSetBits(val), "set bits in the int:", val)
}
