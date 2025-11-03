package main // 실행 가능한 프로그램의 시작점

import (
	"fmt"     // 출력 기능
	"reflect" // 변수의 타입을 확인하는 기능
)

func main() {
	var length float64 = 3.2 // 실수형 변수 선언 및 초기화
	var width int = 2        // 정수형 변수 선언 및 초기화

	fmt.Println("면적은", int(length)*width)               // 실수형 length를 정수형으로 변환 후 곱셈
	fmt.Println("length > width?", int(length) > width) // 형변환 후 비교 연산
	fmt.Println("형변환", reflect.TypeOf(int(length)))     // 형변환된 타입 출력 (int)
	fmt.Println("원본", reflect.TypeOf(length))           // 원래 타입 출력 (float64)
}
