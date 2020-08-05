package Lesson3

import (
	"fmt"
	"strconv"
	"testing"
)

func TestAdding(t *testing.T) {
	str1, str2 := "%^80", "hhhhh20&&&&nd"
	result := Adding(str1, str2)
	if result != 100 {
		t.Error("Wrong answer! Expected 100, got ", result)
	}
}

func TestGetQuotient(t *testing.T) {
	str1 := "1 149,6088607594936;1 179,0666666666666"
	result, _ := strconv.ParseFloat(fmt.Sprintf("%.4f", GetQuotient(&str1)), 64)
	if result != 0.9750 {
		t.Error("Wrong asnwer! Expected 0.9750, got ", result)
	}
}
