package main

import (
	"fmt"
)

func distinct(val string) int {
	//create map to check distinct strings
	keys := make(map[string]bool)
	variation := 0
	//loop through every possible character
	for i := 0; i < len([]rune(val))-1; i++ {
		//loop against every other possible character
		for j := i; j <= len([]rune(val)); j++ {
			//check against map
			if _, value := keys[val[i:j]]; !value {
				keys[val[i:j]] = true
			}
			//empty character is a distinct substring, but it clashes with a space character, other characters can be added here to fix those
			if val[i:j] == " " {
				variation++
			}
		}
	}
	//return length of map
	return len(keys) + variation
}

func main() {
	fmt.Println(distinct("Hello World."))
}
