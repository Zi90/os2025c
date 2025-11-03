package main

import (
	"bufio" // 입력을 버퍼로 처리
	"fmt"
	"os" // 표준 입력을 위한 패키지
)

func main() {
	r := bufio.NewReader(os.Stdin) // 표준 입력을 읽는 Reader 생성
	i, _ := r.ReadString('\n')     // 줄 끝까지 입력 받기, 에러는 무시
	fmt.Println(i)                 // 입력한 문자열 출력
}
