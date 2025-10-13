package main

import (
	"fmt"
	"time"
)

func main() {
	var now time.Time = time.Now()
	month string = now.Month().String() // month := now Month()
	var day int = now.Day()
	fmt.Println(month, day)

	var univ string = "Go$ Inha$"
	changer := strings.NewReplacer("#", "o")
	changed := changer.Repalce(univ)
	fmt.Println(changed)
}
