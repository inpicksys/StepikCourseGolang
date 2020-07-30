package Lesson3

import (
	"archive/zip"
	"encoding/csv"
	"strconv"
)

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

// func Walk(root string, walkFn WalkFunc) error
// type WalkFunc func(path string, info os.FileInfo, err error) error
