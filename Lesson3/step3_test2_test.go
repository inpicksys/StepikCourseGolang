package Lesson3

import "testing"

func TestInterfaceWork(t *testing.T) {
	value1, value2, value3 := 10.0, 2.0, "+"
	result := InterfaceWork(value1, value2, value3)
	if result != "12.0000" {
		t.Error("Wrong answer! Expected 12.0000, got ", result)
	}

	value1, value2, value3 = 10.0, 2.0, "-"
	result = InterfaceWork(value1, value2, value3)
	if result != "8.0000" {
		t.Error("Wrong answer! Expected 8.0000, got ", result)
	}

	value1, value2, value3 = 10.0, 2.0, "*"
	result = InterfaceWork(value1, value2, value3)
	if result != "20.0000" {
		t.Error("Wrong answer! Expected 20.0000, got ", result)
	}

	value1, value2, value3 = 10.0, 2.0, "/"
	result = InterfaceWork(value1, value2, value3)
	if result != "5.0000" {
		t.Error("Wrong answer! Expected 5.0000, got ", result)
	}

	value1Wrong := true
	value2, value3 = 2.0, "-"
	result = InterfaceWork(value1Wrong, value2, value3)
	if result != "value=true: bool" {
		t.Error("Wrong answer! Expected value=true: bool, got ", result)
	}

	value3Wrong := "123"
	value1, value2 = 2.0, 3.0
	result = InterfaceWork(value1, value2, value3Wrong)
	if result != "неизвестная операция" {
		t.Error("Wrong answer! Expected неизвестная операция, got ", result)
	}
}
