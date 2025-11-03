package main

import (
	"bufio"
	"fmt"
	"log" // 에러 발생 시 로그 출력 후 종료
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	i, err := r.ReadString('\n') // 입력 받기, 에러 변수도 함께 처리
	if err != nil {
		log.Fatal(err) // 에러 발생 시 프로그램 종료
	}
	fmt.Println(i) // 정상 입력 시 출력
}
