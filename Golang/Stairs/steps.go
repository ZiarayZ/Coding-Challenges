package main

import (
	"fmt"
)

//climb 3 steps, or 2-1 or 1-2 or 1-1-1	= 4
//climb 2 steps, or 1-1					= 2
//climb 1 steps							= 1
func climb(n int) int {
	//return exact combos at or below 3 steps
	if n == 3 {
		return 4
	} else if n == 2 {
		return 2
	} else if n == 1 {
		return 1
	}

	//return tribonacci, formula: F(n) = F(n-1) + F(n-2) + F(n-3), with base cases above ^^
	return climb(n-1) + climb(n-2) + climb(n-3)
}

func main() {
	steps := 27
	fmt.Println("There are", climb(steps), "ways to climb", steps, "steps.")
}
