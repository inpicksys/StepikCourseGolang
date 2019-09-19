package main

import "fmt"

//TODO Напишите программу, которая последовательно делает следующие операции с каждым числом: число умножается на 2; к числу прибавляется 100.
func stepOneExtendedWork() int {
	var a int
	fmt.Scan(&a)
	a *= 2
	a += 100
	return a
}

//TODO вывести квадрат данного числа
func squarer() int {
	var a int
	fmt.Scan(&a)
	return a * a
}

//TODO Дано натуральное число, выведите его последнюю цифру.
func last_digit() int {
	var a int
	fmt.Scan(&a)
	return a % 10
}

//TODO Дано неотрицательное целое число. Найдите число десятков (то есть вторую цифру справа)
func pre_last_digit() uint {
	var a uint
	fmt.Scan(&a)
	return (a / 10) % 10
}

TODO Часовая стрелка повернулась с начала суток на d градусов. Определите, сколько сейчас целых часов h и целых минут m