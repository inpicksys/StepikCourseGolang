package Lesson1

import "fmt"

//TODO Дан массив, состоящий из целых чисел. Нумерация элементов начинается с 0.
//Напишите программу, которая выведет элементы массива, номера которых четны (0, 2, 4...).

func evenNumbersFromArray(size uint) {
	var element int
	arr := make([]int, size)
	var counter uint
	for counter = 0; counter < size; counter++ {
		fmt.Scan(&element)
		arr[counter] = element
	}
	for i, value := range arr {
		if i%2 == 0 {
			fmt.Printf("%d ", value)
		}
	}
}
