package Lesson3

import "testing"

func TestInterfaceWork2(t *testing.T) {
	// 	1000010011 -> [      XXXX]
	cap := "1000010011"
	result := InterfaceWork2(cap)
	if result != "[      XXXX]" {
		t.Error("Wrong answer! Expected [      XXXX], got ", result)
	}
}
