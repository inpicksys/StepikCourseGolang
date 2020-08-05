package Lesson3

import (
	"archive/zip"
	"bufio"
	"encoding/csv"
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

//TODO В тестовом архиве task.zip содержится набор папок и файлов. Один из этих файлов является
// файлом с данными в формате CSV, прочие же файлы структурированных данных не содержат.
// Требуется найти и прочитать этот единственный файл со структурированными данными
// (это таблица 10х10, разделителем является запятая), а в качестве ответа необходимо указать
// число, находящееся на 5 строке и 3 позиции (индексы 4 и 2 соответственно).
func ZipWork(r *zip.ReadCloser) (data int) {
	for _, file := range r.File {
		if !file.FileInfo().IsDir() {
			read, _ := file.Open()
			lines, err := csv.NewReader(read).ReadAll()
			if err != nil {
				continue
			}
			if len(lines) == 10 && len(lines[4]) == 10 {
				data, _ = strconv.Atoi(lines[4][2])
			}
		}
	}
	return
}

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
