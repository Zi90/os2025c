package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	// i, _ := r.ReadString('\n') // (, _ )= ignore error
	i, err := r.ReadString('\n')
	// fmt.Println(err)
	if err != nil {
		log.Fatal(err) // report the error and exit the program
	}
	fmt.Println(i)
}
