package main

import (
	"fmt"
)

func isNumeric(val string) int {
	return 0
}

func main() {
	num := "7"
	ret := isNumeric(num)
	num = "'" + num + "'"
	check := "is"
	if ret == 0 {
		check = "is not"
	}
	fmt.Println("Value", num, check, "a valid numeric.")
}
