package Lesson3

// TODO Вы должны объявить функцию вида func(uint) uint, которую в дальнейшем можно будет вызвать по имени fn.
//  Функция на вход получает целое положительное число (uint), возвращает число того же типа, в котором отсутствуют
//  нечетные цифры или цифра 0; если же получившееся число равно 0, то возвращается число 100.
//  Цифры в возвращаемом числе имеют тот же порядок, что и в исходном числе.
//  Используйте "fmt" и "strconv", другие пакеты импортировать нельзя.
func Lambda(num uint) uint {
	fn := func(num uint) uint {
		var counter, result, division uint
		division = 1
		for counter = 10; num > 0; num /= counter {
			currentDigit := num % counter
			if currentDigit != 0 && currentDigit%2 == 0 {
				result += currentDigit * division
				division *= 10
			}
		}
		if result == 0 {
			result = 100
		}
		return result
	}
	return fn(num)
}
