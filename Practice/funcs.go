package Practice

import (
	"math"
	"strconv"
)

//TODO На вход подаются a и b - катеты прямоугольного треугольника. Нужно найти длину гипотенузы
func pifagorus(a, b, c *float64) {
	*c = math.Sqrt((*a)*(*a) + (*b)*(*b))
}

//TODO Дана строка, содержащая только английские буквы (большие и маленькие).
// Добавить символ ‘*’ (звездочка) между буквами
// (перед первой буквой и после последней символ ‘*’ добавлять не нужно).
func stringSplitter(text, result *string) {
	for counter := 0; counter < len(*text)-1; counter++ {
		*result += string((*text)[counter])
		*result += string('*')
	}
	*result += string((*text)[len(*text)-1])
}

//TODO Дана строка, содержащая только десятичные цифры. Найти и вывести наибольшую цифру.
func findMaxInDigitRow(transcript *string) int {
	tempMax := 0
	for _, letter := range *transcript {
		digit, _ := strconv.Atoi(string(letter))
		if digit > tempMax {
			tempMax = digit
		}
	}
	return tempMax
}

//TODO На вход подается целое число. Необходимо возвести в квадрат каждую цифру числа и вывести получившееся число.
// Например, у нас есть число 9119. Первая цифра - 9. 9 в квадрате - 81. Дальше 1. Единица в квадрате - 1.
// В итоге получаем 811181
func whatTheFunc(source int) *string {
	var result string
	for _, letter := range strconv.Itoa(source) {
		digit, _ := strconv.Atoi(string(letter))
		result += strconv.Itoa(digit * digit)
	}
	return &result
}

//TODO Требуется вычислить период колебаний (t) математического маятника (мы округлили некоторые значения
// для удобства проверки), для этого нужно найти циклическую частоту колебания пружинного маятника (w),
// в формуле w встречается масса которую также нужно найти, все нужные формулы приведены ниже.
//Напишите три функции, каждая из которых будет выполнять конкретную формулу. Название функций обязательно
//должны соответствовать букве формулы: T(), W() и M(). Для того чтобы найти t - необходимо сначало найти w,
//и тд. Так что используйте результат функции W() в формуле функции T() - то-есть вызывайте функцию W() в T().
//Анологично и с W(), M().
//t= 6 \ w, w = sqrt(k \ m), m = p * v
func T(k, p, v float64) float64 {
	return 6 / W(k, p, v)
}

func W(k, p, v float64) float64 {
	return math.Sqrt(k / M(p, v))
}

func M(p, v float64) float64 {
	return p * v
}
