package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv" // 문자열 → 숫자 변환
	"strings" // 문자열 다듬기
)

func main() {
	fmt.Print("Enter a score: ") // 사용자에게 입력 요청
	r := bufio.NewReader(os.Stdin)
	i, err := r.ReadString('\n') // 입력 받기
	if err != nil {
		log.Fatal(err) // 에러 발생 시 종료
	}

	i = strings.TrimSpace(i)                // 입력값의 공백 제거
	score, err := strconv.ParseFloat(i, 64) // 문자열을 실수로 변환
	if err != nil {
		log.Fatal(err)
	}

	if score >= 60 { // 조건문으로 합격 여부 판단
		fmt.Println(score, "Pass")
	} else {
		fmt.Println(score, "Fail")
	}

	// shadowing 예시
	var float64 float64 = 2.71 // 변수 이름이 타입 이름과 같음 (가능하지만 혼란 유발)
	var f float64 = 3.991
	fmt.Println(float64) // 2.71 출력
	fmt.Println(f)       // 3.991 출력

	var fmt float64 = 2.71 // fmt 패키지 이름과 같은 변수명 → 패키지 접근 불가
	fmt.Println(fmt)       // 2.71 출력, 하지만 fmt.Println()은 더 이상 사용 불가
}
