package Lesson1

import "fmt"

//TODO Напишите программу, принимающая на вход число N(N≥4),
//а затем N целых чисел-элементов массива.
//На вывод нужно подать 4-ый элемент данного массива.

func arrFourthElement(size uint) int {
	var element int
	arr := make([]int, size)
	var i uint
	for i = 0; i < size; i++ {
		fmt.Scan(&element)
		arr[i] = element
	}
	fmt.Println(arr)
	return arr[3]
}
