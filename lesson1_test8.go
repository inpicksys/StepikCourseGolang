package main

import "fmt"

func firstDigit() int {
	//TODO Дано неотрицательное целое число <= 10000. Найдите и выведите первую цифру числа.
	var a int
	fmt.Scan(&a)
	if a < 10 {
		return a
	} else if a < 100 {
		return a / 10
	} else if a < 1000 {
		return a / 100
	} else if a < 10000 {
		return a / 1000
	} else if a == 10000 {
		return 1
	} else {
		return 0
	}
}
