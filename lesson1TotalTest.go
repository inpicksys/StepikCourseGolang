package main

import (
	"fmt"
)

//TODO Дано трехзначное число. Найдите сумму его цифр.
func sumOfDigits(data uint) uint {
	var i, result uint
	for i = 1000; i >= 10; i /= 10 {
		result += data % i / (i / 10)
	}
	return result
}

//TODO Дано трехзначное число. Переверните его, а затем выведите.
func reshaperThreeDigits(data int) {
	for i := 10; i <= 1000; i *= 10 {
		fmt.Print(data % i / (i / 10))
	}
}

//TODO Идёт k-я секунда суток. Определите, сколько целых часов h и целых минут m прошло с начала суток. Например, если
func clockExtReader(k uint) {
	var hours, minutes uint
	hours = k / 3600
	minutes = (k - hours*3600) / 60
	fmt.Printf("It is %d hours %d minutes.", hours, minutes)
}

//TODO Заданы три числа - a,b,c(a<b<c) - длины сторон треугольника. Нужно проверить, является ли треугольник прямоугольным. Если является, вывести "Прямоугольный". Иначе вывести "Не прямоугольный"
func is90Triangle(a, b, c uint) {
	if a*a+b*b == c*c {
		fmt.Println("Прямоугольный")
	} else {
		fmt.Println("Не прямоугольный")
	}
}

//TODO Даны три натуральных числа a, b, c. Определите, существует ли треугольник с такими сторонами.
func isTriangle(a, b, c uint) {
	if a+b > c && a+c > b && b+c > a {
		fmt.Println("Существует")
	} else {
		fmt.Println("Не существует")
	}
}

//TODO Даны два числа. Найти их среднее арифметическое.
func middleValue(a, b float32) float32 {
	return (a + b) / 2.0
}

//TODO По данным числам, определите количество чисел, которые равны нулю.
/*
Входные данные
Вводится натуральное число N, а затем N чисел.
Выходные данные
Выведите количество чисел, которые равны нулю.
*/
func howManyZeroes(num uint) uint {
	var i, counter uint
	var element int
	for i = 0; i < num; i++ {
		fmt.Scan(&element)
		if element == 0 {
			counter++
		}
	}
	return counter
}

//TODO Найдите количество минимальных элементов в последовательности.
/*
Входные данные
Вводится натуральное число N, а затем N чисел.
Выходные данные
Выведите количество минимальных элементов.
*/
func minimumInTheList(num uint) uint {
	var i, counter uint
	var minimum, element int
	fmt.Scan(&element)
	minimum = element
	counter = 1
	for i = 0; i < num-1; i++ {
		fmt.Scan(&element)
		if element == minimum {
			counter++
		} else if element < minimum {
			minimum = element
			counter = 1
		}
	}
	return counter
}

//TODO По данному числу определите его цифровой корень.
/*
Цифровой корень натурального числа — это цифра, полученная в результате итеративного процесса суммирования цифр,
на каждой итерации которого для подсчета суммы цифр берут результат, полученный на предыдущей итерации.
Этот процесс повторяется до тех пор, пока не будет получена одна цифра.
Например цифровой корень 65536 это 7 , потому что 6+5+5+3+6=25 и 2+5=7 .
*/
func digitalRoot(num uint) uint {
	var result uint
	for i := num; i > 0; i /= 10 {
		result += i % 10
	}
	if result > 10 {
		result = digitalRoot(result)
	}
	return result
}

//TODO Найдите самое большее число на отрезке от a до b, кратное 7 .
/*
Входные данные
Вводится два целых числа a и b (a≤b).
Выходные данные
Найдите самое большее число на отрезке от a до b, кратное 7 , или выведите "NO" - если таковых нет.
*/
func dividedBy7MaxOnInterval(a, b int) int {
	var max = a
	for i := a; i <= b; i++ {
		if i%7 == 0 && max < i {
			max = i
		}
	}
	return max
}

//TODO По данному числу n закончите фразу "На лугу пасется..." одним из возможных продолжений: "n коров", "n корова", "n коровы", правильно склоняя слово "корова".
/*
Входные данные
Дано число n (n<100).
Выходные данные
Программа должна вывести введенное число n и одно из слов (на латинице): korov, korova или korovy, например, 1 korova, 2 korovy, 5 korov. Между числом и словом должен стоять ровно один пробел.
*/
func speakLikeUkupnik(num uint) {
	cows := ""
	if num > 10 && num < 20 {
		cows = "korov"
	} else {
		switch num % 10 {
		case 1:
			cows = "korova"
		case 2:
			cows = "korovy"
		case 3:
			cows = "korovy"
		case 4:
			cows = "korovy"
		default:
			cows = "korov"
		}
	}
	fmt.Printf("%d %s", num, cows)
}

//TODO По данному числу N распечатайте все целые степени двойки, не превосходящие N, в порядке возрастания.
func powersOfTwo(border uint) {
	var i uint
	for i = 1; i <= border; i *= 2 {
		fmt.Printf("%d ", i)
	}
}

//TODO Дана строка, содержащая только английские буквы (большие и маленькие). Добавить символ ‘*’ (звездочка) между буквами (перед первой буквой и после последней символ ‘*’ добавлять не нужно).
func addTheAsterisk(str string) {
	fmt.Printf("%s", string(str[0]))
	for i := 1; i < len(str); i++ {
		fmt.Printf("*%s", string(str[i]))
	}
}

//TODO Дана строка, содержащая только десятичные цифры. Найти и вывести наибольшую цифру.

//TODO Из натурального числа удалить заданную цифру.
/*
Входные данные
Вводятся натуральное число и цифра, которую нужно удалить.
Выходные данные
Вывести число без заданных цифр.
*/
