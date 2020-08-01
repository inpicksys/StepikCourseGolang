package Lesson3

import (
	"bufio"
	"os"
	"testing"
)

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
