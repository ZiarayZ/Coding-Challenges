package main

import (
	"fmt"
	"math"
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
	number1 := num1
	number2 := num2
	digits1 := countDigits(num1)
	digits2 := countDigits(num2)
	pow1 := int(math.Pow10(digits1 - 1))
	pow2 := int(math.Pow10(digits2 - 1))

	//if num1 has less digits than num2 then this loop will never return true
	for i := 0; i < digits1-1 && digits1 >= digits2; i++ {
		firstDigit := num1 / pow1
		left := ((num1 * 10) + firstDigit) - (firstDigit * pow1 * 10)
		num1 = left
		if num1 == number2 {
			return true
		}
	}
	//if num2 has less digits than num1 then this loop will never return true
	for i := 0; i < digits2-1 && digits1 <= digits2; i++ {
		firstDigit := num2 / pow2
		left := ((num2 * 10) + firstDigit) - (firstDigit * pow2 * 10)
		num2 = left
		if number1 == num2 {
			return true
		}
	}

	return false
}

func recycle(nums []int) int {
	count := 0
	paired := []int{}
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			//ignore paired j's
			truth := false
			for _, a := range paired {
				if a == j {
					truth = true
				}
			}
			//check rotations
			if !truth && rotate(nums[i], nums[j]) {
				count++
				paired = append(paired, j)
				j = len(nums)
			}
		}
	}
	return count
}

func main() {
	numbers := []int{69, 96, 102, 21, 5}
	fmt.Println(recycle(numbers))
}
