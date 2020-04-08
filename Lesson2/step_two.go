package Lesson2

import (
	"errors"
	"strings"
	"unicode"
)

//TODO На вход подается строка. Нужно определить, является ли она правильной или нет.
// Правильная строка начинается с заглавной буквы и заканчивается точкой.
// Если строка правильная - вывести Right иначе - вывести Wrong
func isPerfectString(text *string) string {
	textRune := []rune(*text)
	if unicode.IsUpper(textRune[0]) && textRune[len(textRune)-3] == 46 {
		return "Right"
	}
	return "Wrong"
}

//TODO На вход подается строка. Нужно определить, является ли она палиндромом.
// Если строка является палиндромом - вывести Палиндром иначе - вывести Нет
func isPalindrome(text *string) string {
	textRune := []rune(*text)
	length := len(textRune) - 2
	for i := 0; i < length/2; i++ {
		if textRune[i] != textRune[length-i] {
			return "Нет"
		}
	}
	return "Палиндром"
}

//TODO Даются две строки X и S. Нужно найти и вывести первое вхождение подстроки S в строке X.
// Если подстроки S нет в строке X - вывести -1
func findSubstring(str1, str2 *string) int {
	return strings.Index(*str1, *str2)
}

//TODO  На вход дается строка, из нее нужно сделать другую строку, оставив только нечетные символы (считая с нуля)
func onlyNotEvenLetters(text, result *string) {
	for counter := 1; counter < len(*text); {
		*result += string((*text)[counter])
		counter += 2
	}
}

//TODO Дается строка. Нужно удалить все символы, которые встречаются более одного раза и вывести получившуюся строку
func onlyUniqueLetters(text, result *string) {
	for _, letter := range *text {
		if !strings.Contains(*result, string(letter)) && strings.Count(*text, string(letter)) == 1 {
			*result += string(letter)
		}
		continue
	}
}

//TODO Ваша задача сделать проверку подходит ли пароль вводимый пользователем под заданные требования.
// Длина пароля должна быть не менее 5 символов, он должен содержать только цифры и буквы латинского алфавита.
// На вход подается строка-пароль. Если пароль соответствует требованиям - вывести "Ok", иначе вывести "Wrong password"
func isAppropriatePassword(password, verdict *string) {
	marker := true
	length := len(*password)
	for _, symbol := range *password {
		if unicode.Is(unicode.Latin, symbol) || unicode.IsDigit(symbol) {
			continue
		}
		marker = false
		break
	}
	if length >= 5 && marker {
		*verdict = "Ok"
	} else {
		*verdict = "Wrong password"
	}
}

//TODO Вы должны вызвать функцию divide которая делит два числа, но она возвращает не только результат деления,
// но и информацию об ошибке. В случае какой-либо ошибки вы должны вывести "ошибка", если ошибки нет -
// результат функции. Функция divide(a int, b int) (int, error) получает на вход два числа,
// которые нужно поделить и возвращает результат (int) и ошибку (error).
func divide(a, b int) (int, error) {
	if b != 0 {
		return a / b, nil
	} else {
		return -1, errors.New("dividing to zero")
	}
}
