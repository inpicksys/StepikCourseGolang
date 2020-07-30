package Lesson3

import (
	"archive/zip"
	"fmt"
)

func Executor() {
	r, err := zip.OpenReader("task.zip")
	if err != nil {
		panic(err)
	}
	defer r.Close()
	result := ZipWork(r)
	fmt.Println(result)
}
