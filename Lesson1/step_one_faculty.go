package Lesson1

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
func lastDigit() int {
	var a int
	fmt.Scan(&a)
	return a % 10
}

//TODO Дано неотрицательное целое число. Найдите число десятков (то есть вторую цифру справа)
func preLastDigit() uint {
	var a uint
	fmt.Scan(&a)
	return (a / 10) % 10
}

//TODO Часовая стрелка повернулась с начала суток на d градусов. Определите, сколько сейчас целых часов h и целых минут m
func clockReader() (int, int) {
	var a int
	fmt.Scan(&a)
	a *= 2
	var h int = a / 60
	var m int = a % 60
	return h, m
}
