package Lesson3

import "testing"

func TestCaller(t *testing.T) {
	result := Caller(727178)
	if result != 28 {
		t.Error("Wrong answer! Expected 28, got ", result)
	}
}
