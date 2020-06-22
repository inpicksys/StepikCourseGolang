package Lesson3

import (
	"bytes"
	"encoding/csv"
	"strconv"
	"strings"
	"unicode"
)

// TODO В привычных нам редакторах электронных таблиц присутствует удобное представление числа
//  с разделителем разрядов в виде пробела, кроме того в России целая часть от дробной отделяется запятой.
//  Набор таких чисел был экспортирован в формат CSV, где в качестве разделителя используется символ ";".
//  На стандартный ввод вы получаете 2 таких вещественных числа, в качестве результата требуется вывести
//  частное от деления первого числа на второе с точностью до четырех знаков после "запятой"
//  (на самом деле после точки, результат не требуется приводить к исходному формату).
//  Алгоритм работы:
//  1. Прочитать строку
//  2. Передать указатель на строку в функцию
//  3. Разбить строку на поля (разделитель - ';')
//  4. Удалить мусор из полученной строки
//  5. Прочитать из строки вещественные числа
//  6. Посчитать их частное и вернуть как результат функции
//  7. Вывести полученное значение на стандартный вывод (точность - 4 знака после запятой)

func GetQuotient(string1 *string) float32 {
	var result float32
	r := csv.NewReader(strings.NewReader(*string1))
	records, err := r.ReadAll()
	if err != nil {
		panic(err)
	}
	numFirstStr := strings.Join(records[0], "")
	numSecondStr := strings.Join(records[1], "")
	result = float32(removeNonDigitsAndConvertToFloat(&numFirstStr) / removeNonDigitsAndConvertToFloat(&numSecondStr))
	return result
}

func removeNonDigitsAndConvertToFloat(s *string) (result float64) {
	resultStr := ""
	strRunes := bytes.Runes([]byte(*s))
	for i := 0; i < len(strRunes); i++ {
		if unicode.IsDigit(strRunes[i]) {
			resultStr += string(strRunes[i])
		}
	}
	result, _ = strconv.ParseFloat(resultStr, 64)
	return
}
