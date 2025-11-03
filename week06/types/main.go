package main // 실행 가능한 프로그램의 시작점이 되는 패키지 이름. Go에서는 반드시 'main'이어야 함

import (
	"fmt"     // 출력 기능을 제공하는 표준 패키지
	"reflect" // 변수의 타입을 확인할 수 있는 패키지
)

func main() { // 프로그램이 시작될 때 가장 먼저 실행되는 함수

	// var id int16           // id라는 이름의 int16 타입 변수 선언
	// var name string        // name이라는 문자열 변수 선언
	// var gpa float32        // gpa라는 실수형(float32) 변수 선언

	// id = 999               // 선언된 id 변수에 값 할당
	// name = "Kim Inha"      // name 변수에 문자열 할당
	// gpa = 3.99             // gpa 변수에 실수값 할당

	// var id int16 = 999     // 선언과 동시에 값 할당
	// var name string = "Kim Inha"
	// var gpa float32 = 3.99

	// var id = 999           // 타입 생략 → 자동으로 int로 추론됨
	// var name = "Kim Inha"  // 자동으로 string 타입으로 추론됨
	// var gpa = 3.99         // 자동으로 float64로 추론됨

	// id := 999              // :=는 변수 선언과 동시에 값 대입 (짧은 선언)
	// name := "Kim Inha"
	// gpa := 3.99

	// fmt.Println("학번은", id, reflect.TypeOf(id), ", 이름은", name, reflect.TypeOf(name))
	// → 변수 값과 타입을 함께 출력

	// fmt.Println("평점 :", gpa, reflect.TypeOf(gpa)) // gpa 값과 타입 출력

	// 아래는 다양한 리터럴 값의 타입을 확인하는 코드들

	// fmt.Println(reflect.TypeOf(2.31))       // 실수 → float64
	// fmt.Println(reflect.TypeOf("Kim Inha")) // 문자열 → string
	// fmt.Println(reflect.TypeOf(true))       // 논리값 → bool
	// fmt.Println(reflect.TypeOf('A'))        // 문자 → int32 (rune)
	// fmt.Println(reflect.TypeOf(19))         // 정수 → int

	// fmt.Println(math.Round(2.21))           // 소수점 반올림 (math 패키지 필요)
	// fmt.Println(math.Ceil(2.21))            // 소수점 올림

	// fmt.Println(strings.Title("go. developer!")) // 문자열 첫 글자 대문자로 (strings 패키지 필요)

	// fmt.Println("Kim\n Inha\t \"\ \") // 이스케이프 문자 출력: 줄바꿈, 탭, 큰따옴표, 역슬래시

	// fmt.Println('2', '가') // 문자 리터럴 출력 (정수값으로 처리됨, 둘 다 유니코드로 변환하여 출력)

	// Zero value (초기값 없이 선언된 변수의 기본값 확인)

	// var 64f float64 // 변수 이름이 숫자로 시작하면 에러 발생
	// var f64 float64 // 선언만 하면 기본값 0.0
	// var t bool      // 기본값 false
	// var s string    // 기본값 ""
	// var i int16     // 기본값 0
	// var i16 int16   // 기본값 0

	// fmt.Println(f64, reflect.TypeOf(f64)) // 0, float64
	// fmt.Println(t, reflect.TypeOf(t))     // false, bool
	// fmt.Println(s, reflect.TypeOf(s))     // "", string
	// fmt.Println(i, reflect.TypeOf(i))     // 0, int16
	// fmt.Println(i16, reflect.TypeOf(i16)) // 0, int16

	var f64 float64 // 선언만 했기 때문에 기본값 0.0이 들어감

	// totalprice := 1000     // 변수 이름은 소문자로 시작하는 카멜 표기법이 권장됨
	// total_price := 1000    // Go에서는 snake_case보다 카멜 표기법을 더 자주 사용함
	totalPrice := 1000 // 카멜 표기법 예시. 변수 이름은 의미 있게 작성하는 것이 좋음

	fmt.Println(totalPrice)               // totalPrice 값 출력
	fmt.Println(f64, reflect.TypeOf(f64)) // f64의 값과 타입 출력 (0, float64)
}
