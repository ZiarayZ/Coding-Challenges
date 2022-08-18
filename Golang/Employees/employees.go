package main

import (
	"fmt"
)

func catch(input map[rune]rune, target rune) rune {

}

func report(input map[rune]rune) map[rune][]rune {

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
	fmt.Println(report(m))
}
