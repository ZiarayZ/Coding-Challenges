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
	newVal := ""
	power := 1
	//find largest power of 2 after val
	for power < val {
		power <<= 1
	}
	power >>= 1 //revert it so it is largest possible power of 2 below val
	//loop from largest power of 2 till 1, creating the binary string from left to right
	for i := power; i > 0; i /= 2 {
		if i <= val {
			val -= i
			newVal += "1"
		} else {
			newVal += "0"
		}
	}
	return newVal
}

func main() {
	bin := "10111"
	dec := 45
	fmt.Println("Converting", bin, "to decimal:", toDecimal(bin))
	fmt.Println("Converting", dec, "to binary:", toBinary(dec))
}
