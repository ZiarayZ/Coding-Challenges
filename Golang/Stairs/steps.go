package main

import (
	"fmt"
	"math"
)

//climb 3 steps, or 2-1 or 1-2 or 1-1-1	= 4
//climb 2 steps, or 1-1					= 2
//climb 1 steps							= 1
func climb(n int) int {
	//return exact combos below 3 steps
	if n == 3 {
		return 4
	} else if n == 2 {
		return 2
	} else if n == 1 {
		return 1
	}

	//calculate every 3 steps is 4 combos
	expo := n / 3
	value := int(math.Pow(4, float64(expo)))
	//add the remainder after all the 3 steps
	//the remainder causes issues
	if n%3 == 2 {
		//TODO: introduce the 5 step method check
		value += (expo + 1) * 2
	} else if n%3 == 1 {
		//TODO: introduce the 4 step method check
		value += expo + 1
	}

	return value
}

func main() {
	steps := 10
	fmt.Println("There are", climb(steps), "ways to climb", steps, "steps.")
}
