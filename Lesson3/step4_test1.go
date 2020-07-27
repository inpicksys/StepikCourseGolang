package Lesson3

import (
	"bufio"
	"io"
	"os"
	"strconv"
)

//TODO на стандартный ввод подаются целые числа в диапазоне 0-100,
// каждое число подается на стандартный ввод с новой строки (количество чисел не известно).
// Требуется прочитать все эти числа и вывести в стандартный вывод их сумму.
func readFromStdinAndReturnSum() int {
	scanner := bufio.NewScanner(os.Stdin)
	result := 0
	for scanner.Scan() {
		dataRaw := scanner.Text()
		data, err := strconv.ParseInt(dataRaw, 10, 128)
		if data >= 0 && data <= 100 {
			result += int(data)
		}
		if err != nil && err != io.EOF {
			panic(err)
		}
	}
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	writer.WriteString(strconv.Itoa(result))
	return result
}
