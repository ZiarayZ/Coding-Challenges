package main

import (
	"fmt"
)

func catchRune(v rune, m map[rune]rune) []rune {
	slice := []rune{v}
	for k, r := range m {
		if v == r {
			slice = append(slice, catchRune(k, m)...)
		}
	}
	return slice
}

func reportRune(input map[rune]rune) map[rune][]rune {
	m := make(map[rune][]rune)

	//separate loop for populating map, so there's no chance of overwriting/erasing runes
	for k := range input {
		m[k] = []rune{}
	}

	//write in initial runes
	for k, r := range input {
		if k != r {
			m[r] = append(m[r], catchRune(k, input)...)
		}
	}

	return m
}

func catchString(v string, m map[string]string) []string {
	slice := []string{v}
	for k, r := range m {
		if v == r {
			slice = append(slice, catchString(k, m)...)
		}
	}
	return slice
}

func reportString(input map[string]string) map[string][]string {
	m := make(map[string][]string)

	//separate loop for populating map, so there's no chance of overwriting/erasing runes
	for k := range input {
		m[k] = []string{}
	}

	//write in initial runes
	for k, r := range input {
		if k != r {
			m[r] = append(m[r], catchString(k, input)...)
		}
	}

	return m
}

func main() {
	mRune := make(map[rune]rune)
	mString := make(map[string]string)
	//example input
	mRune['A'] = 'A'
	mRune['B'] = 'A'
	mRune['C'] = 'B'
	mRune['D'] = 'B'
	mRune['E'] = 'D'
	mRune['F'] = 'E'
	mString["A"] = "A"
	mString["B"] = "A"
	mString["C"] = "B"
	mString["D"] = "B"
	mString["E"] = "D"
	mString["F"] = "E"
	fmt.Println()
	fmt.Println("Example test using runes:")
	fmt.Println(mRune)
	fmt.Println(reportRune(mRune))
	fmt.Println()
	fmt.Println("Example test with easy to read strings:")
	fmt.Println(mString)
	fmt.Println(reportString(mString))
}
