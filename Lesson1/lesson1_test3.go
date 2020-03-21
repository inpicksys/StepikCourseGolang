package Lesson1

import "fmt"

func someFunc() {
	//TODO Исправить ошибку компиляции
	a := 8
	b := 10
	a = a + b
	b = b + a
	fmt.Println(a)
}
