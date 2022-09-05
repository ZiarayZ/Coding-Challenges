package main

import (
	"fmt"
)

func toDecimal(val string) int {
	r := []rune(val)
	newVal := 0
	//reverse loop through string
	for i, j := len(r)-1, 1; i >= 0; i-- {
		//hit a one, add 'j' which doubles each loop: 1 2 4 8 16 32 64
		if r[i] == '1' {
			newVal += j
		}
		j *= 2
	}
	return newVal
}

func toBinary(val int) string {
	return "0"
}

func main() {
	bin := "10111"
	dec := 45
	fmt.Println("Converting", bin, "to decimal:", toDecimal(bin))
	fmt.Println("Converting", dec, "to binary:", toBinary(dec))
}
