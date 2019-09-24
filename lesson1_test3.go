package main

import "fmt"

func some_func() {
	//TODO Исправить ошибку компиляции
	var a int = 8
	var b int = 10
	a = a + b
	b = b + a
	fmt.Println(a)
}
