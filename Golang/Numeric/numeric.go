package main

import (
	"fmt"
	"strings"
)

func isNumeric(val string) int {
	//strip whitespace from string as " 0.1 " still counts
	rs := []rune(strings.TrimSpace(val))
	nums := []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}
	//bools for duplicate "e" and "."
	eC := false
	dotC := false
	//bool for digit validation
	numCheck := false
	//for storing rs[i]
	chara := ' '
	//loop through each rune
	for i := 0; i < len(rs); i++ {
		//check for 0-1 "." in string for decimals, 0-1 "e" for values like: "2e10"
		//"e" and "." can only be sandwiched by numbers and cannot be the first or last character
		chara = rs[i]
		if ((i == 0 || i == len(rs)-1) && (chara == '.' || chara == 'e')) || (chara == '.' && dotC) || (chara == 'e' && eC) {
			//long check for various invalid "e" and "."
			return 0
		} else if rs[i] == '.' { //checks for a ".", makes sure it's only a single instance
			dotC = true
		} else if rs[i] == 'e' { //checks for a "e", makes sure it's only a single instance
			eC = true
		} else { //if chara is not a number 0-9, return 0
			numCheck = false
			for _, item := range nums {
				if item == chara {
					numCheck = true
				}
			}
			if !numCheck {
				return 0
			}
		}
	}
	return 1
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
