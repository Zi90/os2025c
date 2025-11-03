package main

import "fmt"

func main() {
	// var arrayBool [3]bool = [3]bool{true, false, true}
	// var arrayInt [3]int
	// fmt.Println(arrayBool[1]) // zero value(false)
	// arrayInt[1]++
	// arrayInt[1] = arrayInt[1] + 1
	// fmt.Println(arrayInt[1]) // zero value + 2

	// arrayBool := [3]bool{true, false, true}
	// var arrayInt [3]int
	// fmt.Println(reflect.TypeOf(arrayBool))
	// fmt.Printf("%#v\n", arrayBool)

	// fmt.Println(arrayBool[1])
	// arrayInt[1] = 2
	// fmt.Println(arrayInt[1])

	// arrayBool := [3]bool{true, false, true}
	// arrayInt := [3]int{-9, 11, 7}
	// for i := 0; i < len(arrayBool); i++ {
	// 	fmt.Println(i, arrayBool[i])
	// 	fmt.Println(i, arrayInt[i])
	// }

	numbers := [3]int{-9, 11, 7}
	for i, number := range numbers {
		fmt.Println(i, number)
	}
}
