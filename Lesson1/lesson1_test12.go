package Lesson1

import "fmt"

//TODO Напишите программу, которая выводит квадраты натуральных чисел от 1 до 10. Квадрат каждого числа должен выводится в новой строке.

func squaresToTen() {
	for i := 1; i < 11; i++ {
		fmt.Println(i * i)
	}
}
