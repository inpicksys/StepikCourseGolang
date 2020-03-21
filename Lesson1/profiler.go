package Lesson1

import "fmt"

func profile() (int, string) {
	var age int
	var name string
	fmt.Println("What's your name?")
	fmt.Scan(&name)
	fmt.Println("What's your age?")
	fmt.Scan(&age)
	return age, name
}
