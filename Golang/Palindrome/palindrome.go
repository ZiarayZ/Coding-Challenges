package main

import (
	"fmt"
)

func splitInt(n int) []int {
	slice := []int{}
	//loop and divide by 10 getting each digit with each division
	//due to rounding down just use "n > 0" as loop condition
	for ; n > 0; n /= 10 {
		slice = append(slice, n%10)
	}
	return slice
}

func paliCheck(n int) bool {
	//single digit is a palindrome, base case
	if n < 10 && n > -10 {
		return true
	}

	//otherwise split the digits and check
	slice := splitInt(n)
	return len(slice) <= 1
}

func main() {
	fmt.Println(paliCheck(9), paliCheck(45654))
}
