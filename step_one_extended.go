package main

import "fmt"

//TODO Напишите программу, которая последовательно делает следующие операции с каждым числом:
//Число умножается на 2;
//Затем к числу прибавляется 100.
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
