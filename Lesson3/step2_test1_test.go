package Lesson3

import "testing"

func TestAdding(t *testing.T) {
	str1, str2 := "%^80", "hhhhh20&&&&nd"
	result := Adding(str1, str2)
	if result != 100 {
		t.Error("Wrong answer! Expected 100, got ", result)
	}
}
