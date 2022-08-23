package main

import (
	"fmt"
)

func calculate(A []int, B []int, gas int, pos int, loc int, count int) (int, bool) {
	N := len(A)
	N2 := len(B)
	if N != N2 {
		return -1, false
	} else if N == 1 {
		return 0, true
	} else if pos >= N-1 {
		pos = -1
	}
	pos++
	gas += A[pos]
	if pos == loc && count > 0 {
		return pos, true
	}
	if gas >= B[pos] {
		gas -= B[pos]
		count++
		return calculate(A, B, gas, pos, loc, count)
	} else {
		return pos, false
	}

}

func main() {
	//random set generator
	A := []int{14, 5, 19, 12, 11, 3, 18, 20, 4, 3, 8, 15, 7, 17, 10, 9, 16, 13, 1, 6}
	B := []int{5, 14, 2, 15, 19, 9, 11, 1, 6, 12, 13, 17, 3, 20, 4, 16, 18, 10, 7, 8}
	location, ok := calculate(A, B, 0, -1, 0, 0) //pos is 1 less than the actual location, because you haven't filled up there yet
	fmt.Println(location, ok)
}
