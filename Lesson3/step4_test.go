package Lesson3

import (
	"archive/zip"
	"bufio"
	"os"
	"testing"
)

func TestZipWork(t *testing.T) {
	rc, err := zip.OpenReader("../test_task.zip")
	if err != nil {
		t.Error("No such file! See log: ", err)
	}
	result := ZipWork(rc)
	if result != 42 {
		t.Error("Wrong answer! Expected 42, got ", result)
	}
}

func TestFileWork(t *testing.T) {
	file, err := os.Open("../test_file.txt")
	if err != nil {
		panic(err)
	}
	reader := bufio.NewReader(file)
	result := FileWork(reader)
	if result != 4 {
		t.Error("Wrong answer! Expected 4, got ", result)
	}
}
