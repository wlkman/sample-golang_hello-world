package main

import (
	"fmt"
)

func main() {
	var name string
	var age byte

	fmt.Println("please enter your name")
	fmt.Scanln(&name)

	fmt.Println("please enter your age")
	fmt.Scanln(&age)

	fmt.Printf("name is %v \nage is %v",name,age)
}
