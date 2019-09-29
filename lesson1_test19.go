package main

import "fmt"

//TODO Даны два числа. Определить цифры, входящие в запись как первого, так и второго числа.
/*
Программа получает на вход два числа. Гарантируется, что цифры в числах не повторяются.
Числа в пределах от 0 до 10000.

Выходные данные

Программа должна вывести цифры, которые имеются в обоих числах, через пробел.
Цифры выводятся в порядке их нахождения в первом числе.

Sample Input:

564 8954
Sample Output:

5 4
*/

func areDigitsInOther(a, b uint) {
	var i, j, digitA, digitB uint
	for i = 10000; i >= 10; i /= 10 {
		if i > a*10 {
			continue
		}
		digitA = a % i / (i / 10)
		//fmt.Println("DigitA: ", digitA)
		for j = 10000; j >= 10; j /= 10 {
			if j > b*10 {
				continue
			}
			digitB = b % j / (j / 10)
			if digitA == digitB {
				fmt.Printf("%d ", digitA)
			}
			//fmt.Println("DigitB: ", digitB)
		}
	}
}
