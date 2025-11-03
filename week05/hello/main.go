package main // 실행 가능한 프로그램의 시작점이 되는 패키지 이름

import "fmt" // 출력 기능을 제공하는 표준 패키지(fmt)를 불러옴

func main() { // 프로그램이 시작될 때 가장 먼저 실행되는 함수 정의
	// name := "Go" // 문자열 "Go"를 name이라는 변수에 저장
	// fmt.Println("Hello", name) // name 변수의 값을 포함해 "Hello"를 출력 (줄 바꿈 포함)

	number := 7 // number라는 이름의 변수에 숫자 7을 저장

	fmt.Printf("Number is %d.\n", number) // 포맷 문자열을 사용해 number 값을 출력 (%d는 정수 자리 표시자, \n은 줄 바꿈)
}
