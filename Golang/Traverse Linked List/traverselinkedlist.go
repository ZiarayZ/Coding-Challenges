package main

import (
	"fmt"
	"reflect"
)

func isSlice(arg interface{}) (val reflect.Value, ok bool) {
	val = reflect.ValueOf(arg)
	if val.Kind() == reflect.Slice {
		ok = true
	}
	return
}

func CreateAnyTypeSlice(arg interface{}) ([]interface{}, bool) {
	val, ok := isSlice(arg)

	if !ok {
		return nil, false
	}

	sliceLen := val.Len()
	out := make([]interface{}, sliceLen)

	for i := 0; i < sliceLen; i++ {
		out[i] = val.Index(i).Interface()
	}
	return out, true
}

func faa(arg interface{}) ([]interface{}, bool) {
	arr, ok := CreateAnyTypeSlice(arg)
	if !ok {
		return arr, false
	}
	return arr, true
}

//actual function used to traverse the "linked list" in this case it's just a slice of anything being used as an example
func traverse(M int, N int, arg []interface{}) ([]interface{}, bool) {
	val, ok := isSlice(arg)
	if !ok {
		return arg, false
	}
	newVal := val.Len()
	if M+N > newVal {
		return arg, false
	}

	newArg := make([]interface{}, newVal-N)

	for i := 0; i < M; i++ {
		newArg[i] = arg[i] //here would use the linked list get function
	}
	for i := M + N; i < newVal; i++ {
		newArg[i-N] = arg[i] //here would use the linked list get function
	}
	return newArg, true
}

func main() {
	slice := []int{2, 3, 5, 7, 11, 13, 17}
	fmt.Println(slice)
	newSlice, test := faa(slice)
	if test {
		slicedSlice, testOK := traverse(1, 4, newSlice)
		if testOK {
			fmt.Println(slicedSlice) //should output "[2, 13, 17]"
		} else {
			fmt.Println("Traverse Failed.")
		}
	} else {
		fmt.Println("Conversion Failed.")
	}
}
