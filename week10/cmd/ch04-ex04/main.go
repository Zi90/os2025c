package main

import (
	"os2025c/pkg/greeting"
	"os2025c/pkg/greeting/deutsch"
	"os2025c/pkg/greeting/korean"
)

//import "greeting" // custom package(standard path)

func main() {
	greeting.Hello()
	greeting.Hi()
	deutsch.GutenTag()
	korean.Anyung()
}
