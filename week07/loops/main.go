// 1.
// package main

// import (
// 	"fmt"
// 	"reflect"
// )

// func main() {
// 	var length float64 = 3.2
// 	var width int = 2
// 	fmt.Println("면적은", int(length)*width)
// 	fmt.Println("length > width?", int(length) > width)
// 	fmt.Println("형변환", reflect.TypeOf(int(length)))
// 	fmt.Println("원본", reflect.TypeOf(length))
// }

//2.
// package main

// import (
// 	"fmt"
// 	"time"
// )

// func main() {
// 	var now time.Time = time.Now()
// 	month string = now.Month().String() // month := now Month()
// 	var day int = now.Day()
// 	fmt.Println(month, day)

// 	var univ string = "Go$ Inha$"
// 	changer := strings.NewReplacer("#", "o")
// 	changed := changer.Repalce(univ)
// 	fmt.Println(changed)
// }

//3.
// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// )

// func main() {
// 	r := bufio.NewReader(os.Stdin)
// 	i, _ := r.ReadString('\n')
// 	fmt.Println(i)
// }

//4.
// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"log"
// 	"os"
// )

// func main() {
// 	r := bufio.NewReader(os.Stdin)
// 	// i, _ := r.ReadString('\n') // (, _ )= ignore error
// 	i, err := r.ReadString('\n')
// 	// fmt.Println(err)
// 	if err != nil {
// 		log.Fatal(err) // report the error and exit the program
// 	}
// 	fmt.Println(i)
// }

// 4.
// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"log"
// 	"os"
// 	"strconv"
// 	"strings"
// )

// func main() {

// 	fmt.Print("Enter a score: ")
// 	r := bufio.NewReader(os.Stdin)
// 	i, err := r.ReadString('\n')
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	i = strings.TrimSpace(i)                // 문자열 주위에 붙은 공란 및 탭 키 등 제거
// 	score, err := strconv.ParseFloat(i, 64) // 정리된 문자열을 실수타입으로 변환
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	if score >= 60 {
// 		fmt.Println(score, "Pass")
// 	} else {
// 		fmt.Println(score, "Fail")
// 	}

//shadowing
// var float64 float64 = 2.71
// var f float64 = 3.991
// fmt.Println(float64)
// fmt.Println(f)

// var fmt float64 = 2.71
// fmt.Println(fmt)
// }

