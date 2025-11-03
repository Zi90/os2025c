package korean

import "fmt"

func Anyunghaseyo() {
	fmt.Println("안녕하세요!")
}

func Anyung() { // 함수명이 소문자로 시작하는 경우 외부에서 함수 소환 안 됨
	fmt.Println("안녕!")
}
