package main

//Given (x, y) coordinates, create a function such that each coordinate is uniquely mapped to an integer.
//Also make sure that given an integer, you should be able to find (x, y) coordinates.
//So F(x, y) = z and also that inverse F(z) = (x, y).

import (
	"fmt"
)

//x = even, y = odd
//byte size of 4 or 8 depending on system, so will restrict to 4 for 32-bit systems
//using unsigned ints for bit operations, 2 16 bits become a 32 bit
// coords -> single value
func interleave(coords [2]uint16) uint32 {
	var num uint32 = 0
	var i uint16 = 0
	for i = 0; i < 16; i++ { //loop could be made more efficient later, break after checking all set bits
		num |= uint32(coords[0]&(1<<i)) << i
		num |= uint32(coords[1]&(1<<i)) << (i + 1)
	}
	return num
}

// single value -> coords
func recall(value uint32) [2]uint16 {
	coords := [2]uint16{0, 0}
	var i uint16 = 0
	for i = 0; i < 16; i++ {
		coords[0] |= uint16((value >> (i)) & (1 << i))
		coords[1] |= uint16((value >> (i + 1)) & (1 << i))
	}
	return coords
}

func main() {
	slice := [2]uint16{127, 127}
	fmt.Println(interleave(slice))
	var number uint32 = 16383
	fmt.Println(recall(number))
}
