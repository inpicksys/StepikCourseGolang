package main

import "fmt"

//TODO Требуется написать программу, при выполнении которой с клавиатуры считываются два натуральных числа A и B (каждое не более 100, A < B). Вывести сумму всех чисел от A до B  включительно.

func summer() uint {
	var a, b, result uint
	fmt.Scan(&a, &b)
	for i := a; i <= b; i++ {
		result += i
	}
	return result
}
