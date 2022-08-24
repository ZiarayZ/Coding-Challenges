package main

import (
	"fmt"
)

func unique(runeSlice []rune) []string {
	keys := make(map[rune]bool)
	list := []string{}

	//get every unique character in slice of runes
	//append it all to a slice of strings
	for _, entry := range runeSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, string(entry))
		}
	}

	return list
}

func distinct(val string) int {
	//get all unique chars, for all 1 length distinct substrings
	values := unique([]rune(val))
	//temporary return
	return len(values) + 1 //add 1 for the full string itself
}

func main() {
	//should return "9" as: "helo wrd."
	fmt.Println(distinct("hello world."))
}
