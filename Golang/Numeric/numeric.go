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
	check := ""
	if ret == 0 {
		check = "not"
	}
	fmt.Println("Value", num, "is", check, "a valid numeric.")
}
