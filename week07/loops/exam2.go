package main

import (
	"fmt"
	"strings"
	"time" // 현재 시간 관련 기능
)

func main() {
	var now time.Time = time.Now()          // 현재 시간 정보를 가져옴
	var month string = now.Month().String() // 월 정보를 문자열로 변환
	var day int = now.Day()                 // 일(day) 정보 추출
	fmt.Println(month, day)                 // 월과 일을 출력

	var univ string = "Go$ Inha$"            // 문자열 선언
	changer := strings.NewReplacer("$", "o") // '$'를 'o'로 바꾸는 치환기 생성
	changed := changer.Replace(univ)         // 문자열 치환 실행
	fmt.Println(changed)                     // 치환된 문자열 출력
}
