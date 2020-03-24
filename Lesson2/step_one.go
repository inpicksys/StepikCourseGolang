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
// которое среди значений ее аргументов x, y, z встречается чаще.
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

//TODO Последовательность Фибоначчи определена следующим образом: φ1=1, φ2=1, φn=φn-1+φn-2 при n>1.
// Начало ряда Фибоначчи выглядит следующим образом: 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, ...
// Напишите функцию, которая по указанному натуральному n возвращает φn.
func fibonacci(n int) int {
	prev := 0
	curr := 1
	for i := 1; i < n; i++ {
		temp := curr
		curr = prev + curr
		prev = temp
	}
	return curr
}

//TODO Напишите тело функции, которая умножает значения двух указателей
// и выводит получившееся произведение в консоль. Ввод данных уже написан за вас.
func test(x1 *int, x2 *int) {
	fmt.Println(*x1 * *x2)
}

// TODO Поменяйте значения переменных местами на которые ссылается указатель. После этого переменные нужно вывести.
func valueSwitcher(x1, x2 *int) {
	temp := *x1
	*x1 = *x2
	*x2 = temp
	fmt.Println(*x1, *x2)
}
