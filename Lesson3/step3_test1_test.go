package Lesson3

import "testing"

func TestLambda(t *testing.T) {
	// 727178 28
	result := Lambda(727178)
	if result != 28 {
		t.Error("Wrong answer! Expected 28, got ", result)
	}
}
