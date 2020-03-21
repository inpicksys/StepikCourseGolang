package Lesson2

import "fmt"

func f(text string) {
	fmt.Println(text)
}

//TODO Напишите функцию, находящую наименьшее из четырех введённых в этой же функции чисел.
func minimumFromFour(a, b, c, d int) int {
	var arr [3]int
	arr[0] = b
	arr[1] = c
	arr[2] = d
	temp_min := a
	for i := 0; i < len(arr); i++ {
		if arr[i] < temp_min {
			temp_min = arr[i]
		}
	}
	return temp_min
}

//TODO Напишите "функцию голосования", возвращающую то значение (0 или 1),
//  которое среди значений ее аргументов x, y, z встречается чаще.
func vote(x, y, z int) int {
	var arr [2]int
	arr[0] = y
	arr[1] = z
	temp_value := x
	for i := 0; i < len(arr); i++ {
		if arr[i] == temp_value {
			return arr[i]
		}
	}
	return arr[1]
}
