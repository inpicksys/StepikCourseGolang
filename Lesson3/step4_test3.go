package Lesson3

import (
	"bufio"
	"io"
	"strconv"
)

//TODO В тестовом файле содержится длинный ряд чисел, разделенных символом ";".
// Требуется найти, на какой позиции находится число 0 и указать его в качестве ответа.
// Требуется вывести именно позицию числа, а не индекс (то есть порядковый номер, нумерация с 1).
func FileWork(r *bufio.Reader) uint {
	for i := 1; ; i++ {
		raw, err := r.ReadString(';')
		raw = raw[:len(raw)-1]
		if err != nil && err != io.EOF {
			panic(err)
		}
		data, err := strconv.Atoi(raw)
		if err != nil {
			panic(err)
		}
		if data == 0 {
			return uint(i)
		}
	}
	return 0
}
