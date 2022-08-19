package main

import (
	"fmt"
)

func countDigits(n int) int {
	count := 0

	for n > 0 {
		count++
		n /= 10
	}

	return count
}

func rotate(num1 int, num2 int) bool {
	digits1 := countDigits(num1)
	digits2 := countDigits(num2)
	if digits1 != digits2 {
		return false
	}

	return true
}

func recycle(nums []int) int {
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if rotate(nums[i], nums[j]) {
				fmt.Println("caught rotation")
			}
		}
	}
}

func main() {
	numbers := []int{}
	fmt.Println(recycle(numbers))
}
