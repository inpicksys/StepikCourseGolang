package Lesson1

import "fmt"

//TODO Напишите программу, которая в последовательности чисел находит сумму
//двузначных чисел, кратных 8. Программа в первой строке получает на вход число n -
//количество чисел в последовательности, во второй строке -- n чисел,
//входящих в данную последовательность.

func dividesToEight() uint {
	var counter, a, result uint
	fmt.Scan(&counter)
	for ; counter > 0; counter-- {
		fmt.Scan(&a)
		if a > 99 || a < 10 {
			continue
		} else if a%8 == 0 {
			result += a
		}
	}
	return result
}
