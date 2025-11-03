package main

import (
	"fmt"
	"log"
	"os2025c/pkg/keyboard"
)

func main() {
	fmt.Print("점수 입력 : ")
	score, err := keyboard.GetFloat()
	if err != nil {
		log.Fatal(err)
	}
	if score >= 80 {
		fmt.Printf("%.2f점은 합격입니다!\n", score)
	} else {
		fmt.Print("%.2f점은 불합격입니다!\n", score)
	}
}
