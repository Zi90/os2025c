package main

import (
	"fmt"
	"reflect"
)

func main() {
	// var id int16
	// var name string
	// var gpa float32

	// id = 999
	// name = "Kim Inha"
	// gpa = 3.99

	// var id int16 = 999
	// var name string = "Kim Inha"
	// var gpa float32 = 3.99

	// var id = 999
	// var name = "Kim Inha"
	// var gpa = 3.99

	// id := 999
	// name := "Kim Inha"
	// gpa := 3.99

	// fmt.Println("학번은", id, reflect.TypeOf(id), ", 이름은", name, reflect.TypeOf(gpa))
	// fmt.Println("평점 :", gpa, reflect.TypeOf(gpa))
	// fmt.Println(reflect.TypeOf(2.31))
	// fmt.Println(reflect.TypeOf("Kim Inha"))
	// fmt.Println(reflect.TypeOf(true))
	// fmt.Println(reflect.TypeOf('A'))
	// fmt.Println(reflect.TypeOf(19))
	// fmt.Println(math.Round(2.21))
	// fmt.Println(math.Ceil(2.21))
	// fmt.Println(strings.Title("go. developer!"))
	// fmt.Println("Kim\nInha\t\"\\")
	// fmt.Println('2', '가')

	// Zero value
	// var 64f float64 //error
	// var f64 float64
	// var t bool
	// var s string
	// var i int16
	// var i16 int16

	// fmt.Println(f64, reflect.TypeOf(f64))
	// fmt.Println(t, reflect.TypeOf(t))
	// fmt.Println(s, reflect.TypeOf(s))
	// fmt.Println(i, reflect.TypeOf(i))
	// fmt.Println(i16, reflect.TypeOf(i16))

	var f64 float64
	// totalprice := 1000
	// total_price := 1000
	totalPrice := 1000 // 카멜 표기법, Go 언어에서는 암묵적인 관례

	fmt.Println(totalPrice)
	fmt.Println(f64, reflect.TypeOf(f64))
}
