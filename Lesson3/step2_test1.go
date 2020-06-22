package Lesson3

import (
	"bytes"
	"strconv"
	"unicode"
)

// TODO Представьте что вы работаете в большой компании, где используется модульная архитектура.
//  Ваш коллега написал модуль с какой-то логикой (вы не знаете) и передает в вашу программу какие-то данные.
//  Вы же пишете функцию, которая считывает две переменных типа string, а возвращает число типа int64,
//  которое нужно получить сложением этим строк. Предварительно вам нужно убрать из них мусор и конвертировать в числа.
//  Под мусором имеются ввиду лишние символы и спец знаки. Разрешается использовать только определенные пакеты:
//  fmt, strconv, unicode, strings,  bytes.
func Adding(string1, string2 string) (sumOfStrings int64) {
	sumOfStrings = removeNonDigitsAndConvertToInt(&string1) + removeNonDigitsAndConvertToInt(&string2)
	return
}

func removeNonDigitsAndConvertToInt(s *string) (result int64) {
	resultStr := ""
	strRunes := bytes.Runes([]byte(*s))
	for i := 0; i < len(strRunes); i++ {
		if unicode.IsDigit(strRunes[i]) {
			resultStr += string(strRunes[i])
		}
	}
	result, _ = strconv.ParseInt(resultStr, 10, 64)
	return
}
