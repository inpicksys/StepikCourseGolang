package Lesson3

import (
	"archive/zip"
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
