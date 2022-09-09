package main

import (
	"fmt"
)

// item on stack
type Node struct {
	value   interface{}
	prev    *Node
	minimum *Node
}

//stack itself
type Stack struct {
	last    *Node
	minimum *Node
}

//stack methods
func (s *Stack) min() interface{} {
	if s.minimum == nil {
		return nil
	}
	return s.minimum.value
}
func (s *Stack) pop() interface{} {
	//minimum needs to be changed or returned to nil
	toRet := s.last.value
	if s.last == s.minimum {
		s.minimum = s.last.minimum
	}
	//set last to new last, memory loss?
	s.last = s.last.prev
	return toRet
}
func (s *Stack) push(val interface{}) bool {
	s.last = &Node{value: val, prev: s.last}
	if s.minimum == nil {
		//this can cause issues by being a value that can't be compared
		s.minimum = s.last
		return true
	} else {
		var i int
		var iF float64
		isFl := false
		//assert data type to int or float if possible
		switch testVal := val.(type) {
		case int:
			i = testVal
		case int8:
			i = int(testVal)
		case int16:
			i = int(testVal)
		case int32:
			i = int(testVal)
		case int64:
			i = int(testVal)
		case bool:
			fmt.Printf("%t == %T\n", testVal, testVal)
		case float32:
			//may need to change condition here
			isFl = true
			iF = float64(testVal)
		case float64:
			//may need to change condition here
			isFl = true
			iF = testVal
		case uint8:
			i = int(testVal)
		case uint16:
			i = int(testVal)
		case uint32:
			i = int(testVal)
		case uint64:
			i = int(testVal)
		case string:
			fmt.Printf("%s == %T\n", testVal, testVal)
			// messy to use, but could be done
		default:
			// what is it then?
			fmt.Printf("%v == %T\n", testVal, testVal)
		}
		// grab current minimum as int and float values
		testMin, minOk := s.minimum.value.(int)
		floatMin, floaOk := s.minimum.value.(float64)
		//compares between floats first then ints as floats may have more precision
		if (floaOk && isFl && iF <= floatMin) || (minOk && !isFl && i <= testMin) {
			s.minimum = s.last
		} else {
			s.last.minimum = s.minimum
		}

		//if some bug occurred with the minimum value last push we return that
		//if both are false it failed to compare, if one is true it DID compare
		return minOk || floaOk
	}
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
