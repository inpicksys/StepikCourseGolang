package Lesson2

import (
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

}
