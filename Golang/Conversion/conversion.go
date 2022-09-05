package main

import (
	"fmt"
)

func toDecimal(val string) int {
	return 0
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
