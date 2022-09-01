package main

import (
	"fmt"
)

func splitInt(n int) ([]int, int) {
	slice := []int{}
	//loop and divide by 10 getting each digit with each division
	//due to rounding down just use "n > 0" as loop condition
	for ; n != 0; n /= 10 {
		slice = append(slice, n%10)
	}
	return slice, len(slice)
}

func paliCheck(n int) bool {
	//single digit is a palindrome, base case
	if n < 10 && n > -10 {
		return true
	}

	//otherwise split the digits and check
	slice, num := splitInt(n)
	for i := 0; i < (num+num%2)/2; i++ {
		if slice[i] != slice[num-1-i] {
			return false
		}
	}
	return true
}

func main() {
	//check both possible cases
	fmt.Println(paliCheck(9), paliCheck(45654))
	//ignoring negative cases, as they are treated the same as positive ones
}
