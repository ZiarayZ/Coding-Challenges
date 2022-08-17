package main

//Given a string A, partition A such that every substring of the partition is a palindrome. Return the minimum cuts needed for a palindrome partitioning of A.
// attempt: locate largest palindromes, cut and repeat on remaining substrings

import (
	"fmt"
	"strings"
)

func partite(A string) ([]string, int) {
	//removes whitespace in and between for palindromes
	A = strings.TrimSpace(A)
	A = strings.ReplaceAll(A, " ", "")
	srune := []rune(A)
	returnSlice := []string{string(srune[0]), string(srune[1:])}
	return returnSlice, 1
}

func main() {
	value, cuts := partite("hollo")
	fmt.Println(value, cuts)
}
