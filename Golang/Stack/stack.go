package main

import (
	"fmt"
)

// item on stack
type Node struct {
	value interface{}
	prev  *Node
}

//stack itself
type Stack struct {
	last    *Node
	minimum *Node
}

//stack methods
func (s *Stack) min() interface{} {
	return s.minimum.value
}
func (s *Stack) pop() interface{} {
	toRet := s.last.value
	//set last to new last, memory loss?
	s.last = s.last.prev
	return toRet
}
func (s *Stack) push(val interface{}) bool {
	s.last = &Node{value: val, prev: s.last}
	//only checks for int values, ignored for any other value
	testVal, valOk := val.(int)
	testMin, minOk := s.minimum.value.(int)
	if valOk && minOk && testVal <= testMin {
		s.minimum = s.last
	}
	//some bug occurred with the minimum value last push
	return minOk
}

func main() {
	testStack := Stack{}
	//test methods
	testStack.push(5)
	fmt.Println(testStack.pop())
	fmt.Println(testStack.min())
	testStack.push(1)
	testStack.push(4)
	fmt.Println(testStack.pop())
	fmt.Println(testStack.min())
	testStack.push(2)
	testStack.push(3)
	fmt.Println(testStack.pop())
	fmt.Println(testStack.min())
}
