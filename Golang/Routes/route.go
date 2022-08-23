package main

import (
	"fmt"
)

func calculate(A []int, B []int, gas int, pos int, loc int) (int, bool) {
	N := len(A)
	N2 := len(B)
	if N != N2 {
		return pos, false
	} else if pos == loc {
		return loc, true
	} else if pos >= N {
		pos = 0
	}
	gas += A[pos]
	fmt.Println(gas, B[pos])
	if gas >= B[pos] {
		gas -= B[pos]
		pos++
		return calculate(A, B, gas, pos, loc)
	} else {
		return pos, false
	}

}

func main() {
	A := []int{14, 5, 19, 12, 11, 3, 18, 20, 4, 3, 8, 15, 7, 17, 10, 9, 16, 13, 1, 6}
	B := []int{5, 14, 2, 15, 19, 9, 11, 1, 6, 12, 13, 17, 3, 20, 4, 16, 18, 10, 7, 8}
	location, ok := calculate(A, B, 0, 0, 0)
	fmt.Println(location, ok)
}
