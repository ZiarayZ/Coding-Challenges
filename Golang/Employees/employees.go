package main

import (
	"fmt"
)

func catch(v rune, m map[rune]rune) []rune {
	slice := []rune{v}
	for k, r := range m {
		if v == r {
			slice = append(slice, catch(k, m)...)
		}
	}
	return slice
}

func report(input map[rune]rune) map[rune][]rune {
	m := make(map[rune][]rune)

	//separate loop for populating map, so there's no chance of overwriting/erasing runes
	for k := range input {
		m[k] = []rune{}
	}

	//write in initial runes
	for k, r := range input {
		if k != r {
			m[r] = append(m[r], catch(k, input)...)
		}
	}

	return m
}

func main() {
	m := make(map[rune]rune)
	//example input
	m['A'] = 'A'
	m['B'] = 'A'
	m['C'] = 'B'
	m['D'] = 'B'
	m['E'] = 'D'
	m['F'] = 'E'
	fmt.Println(m)
	fmt.Println(report(m))
}
