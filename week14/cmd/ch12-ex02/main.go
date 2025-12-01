package main

import "fmt"

func safeDivide(a, b int) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("에러 발생:", err)
		}
	}()

	if b == 0 {
		panic("0으로 나눌 수 없습니다!")
	}

	result := a / b
	fmt.Println("결과:", result)
}

func main() {
	fmt.Println("첫 번쨰 호출")
	safeDivide(10, 2)

	fmt.Println("\n두 번쨰 호출")
	safeDivide(10, 0)

	defer fmt.Println("프로그램 종료 - 정리작업")
	fmt.Println("1. 프로그램 시작")
	panic("심각한 에러 발생!")
	fmt.Println("2. 이 줄은 실행되지 않음")
}
