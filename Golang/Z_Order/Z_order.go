package main

//Given (x, y) coordinates, create a function such that each coordinate is uniquely mapped to an integer.
//Also make sure that given an integer, you should be able to find (x, y) coordinates.
//So F(x, y) = z and also that inverse F(z) = (x, y).

import (
	"fmt"
)

func hash(coords [2]int) int {
	return coords[0] + coords[1]
}

func main() {
	slice := [2]int{1, 2}
	fmt.Println(hash(slice))
}
