package Lesson3

import "testing"

func TestGetQuotient(t *testing.T) {
	str1 := "1 149,6088607594936;1 179,0666666666666"
	result := GetQuotient(&str1)
	if result != 0.9750 {
		t.Error("Wrong asnwer! Expected 0.9750, got ", result)
	}
}
