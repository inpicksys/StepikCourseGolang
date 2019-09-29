package main

import "fmt"

//TODO Дан массив, состоящий из целых чисел. Напишите программу,
//которая подсчитывает количество положительных чисел среди элементов массива.

func positiveElementsCounter(size uint) uint {
	var result, index uint
	var element int
	arr := make([]int, size)
	for index = 0; index < size; index++ {
		fmt.Scan(&element)
		arr[index] = element
		if element > 0 {
			result++
		}
	}
	return result
}
