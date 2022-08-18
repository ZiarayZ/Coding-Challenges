package main

//Given a string A, partition A such that every substring of the partition is a palindrome. Return the minimum cuts needed for a palindrome partitioning of A.
// attempt: locate largest palindromes, cut and repeat on remaining substrings

import (
	"fmt"
	"strings"
)

func cut(B string) []string {
	n := len([]rune(B))
	//if the string is a single letter return that
	if n == 1 {
		return []string{B}
	}

	//declare vars
	maxLength := 1
	start := 0
	flag := true

	//double loop to compare runes
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			flag = true

			for k := 0; k < ((j-i)/2)+1; k++ {
				if B[i+k] != B[j-k] {
					flag = false
				}
			}

			//if largest palindrome is the entire input just return the input
			if i == 0 && j+1 == n {
				return []string{B}
			}

			//record last valid largest palindrome index
			if flag && (j-i+1) > maxLength {
				start = i
				maxLength = j - i + 1
			}
		}
	}
	returnSlice := []string{}
	//repeat process for before largest palindrome
	if start > 0 {
		firstSlice := cut(B[0:start])
		returnSlice = append(returnSlice, firstSlice...)
	}
	//cut out largest palindrome and return it
	returnSlice = append(returnSlice, B[start : start+maxLength])
	//repeat process for before and after largest palindrome
	if start+maxLength < n {
		lastSlice := cut(B[start+maxLength:])
		returnSlice = append(returnSlice, lastSlice...)
	}
	//return all cut palindromes
	return returnSlice
}

func partite(A string) ([]string, int) {
	//removes whitespace in and between for palindromes
	A = strings.TrimSpace(A)
	A = strings.ReplaceAll(A, " ", "")
	sliced := cut(A)
	//cuts needed = length of slice - 1
	return sliced, len(sliced) - 1
}

func main() {
	value, cuts := partite("hollo")
	fmt.Println(value, cuts)
}
