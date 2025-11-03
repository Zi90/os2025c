package main

import (
	"fmt"
	"reflect"
)

func main() {
	// var arrayBool [3]bool = 3[bool]{true, false, true}
	// var arrayInt [3]int
	// fmt.Println(arrayBool[1]) // zero value(false)
	// arrayInt[1]++
	// arrayInt[1] = arrayInt[1] + 1
	// fmt.Println(arrayInt[1]) // zero value + 2

	arrayBool := [3]bool{true, false, true}
	var arrayInt [3]int
	fmt.Println(reflect.TypeOf(arrayBool))
	fmt.Printf("%#v\n" arrayBool)

	fmt.Println(arrayBool[1])
	arrayInt[1] = 2
	fmt.Println(arrayInt[1])

}
