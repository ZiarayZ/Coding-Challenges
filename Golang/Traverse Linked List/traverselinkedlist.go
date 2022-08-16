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

func traverse(M int, N int, arg []interface{}) []interface{} {
	return arg
}

func main() {
	slice := []int{2, 3, 5, 7, 11, 13, 17}
	fmt.Println(slice)
	newSlice, test := faa(slice)
	if test {
		fmt.Println(traverse(1, 4, newSlice))
	} else {
		fmt.Println("Conversion Failed.")
	}
}
