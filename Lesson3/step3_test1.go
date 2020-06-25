package Lesson3

import "strconv"

// TODO Объявить функцию вида func(uint) uint, которую в дальнейшем можно будет вызвать по имени fn.
//  Функция на вход получает целое положительное число (uint), возвращает число того же типа,
//  в котором отсутствуют нечетные цифры или цифра 0; если же получившееся число равно 0, то
//  возвращается число 100. Цифры в возвращаемом числе имеют тот же порядок, что и в исходном числе.
func Caller(num uint) uint {
	fn := func(src uint) uint {
		result := ""
		srcString := strconv.FormatUint(uint64(src), 10)
		for _, r := range srcString {
			val, err := strconv.Atoi(string(r))
			if err != nil {
				panic(err)
			}
			if val != 0 && val%2 == 0 {
				result += string(val)
			}
		}
		resultNum, err := strconv.Atoi(result)
		if err != nil {
			panic(err)
		}
		if resultNum > 0 {
			return uint(resultNum)
		}
		return 100
	}
	return fn(num)
}
