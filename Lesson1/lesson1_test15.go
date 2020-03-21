package Lesson1

import "fmt"

//TODO Последовательность состоит из натуральных чисел и завершается числом 0. Определите количество элементов этой последовательности. которые равны ее наибольшему элементу.

func equalsMaxCounter() uint {
	var a, max, counter uint
	fmt.Scan(&a)
	for a != 0 {
		if a > max {
			max = a
			counter = 1
		} else if max == a {
			counter++
		}
		fmt.Scan(&a)
	}
	return counter
}
