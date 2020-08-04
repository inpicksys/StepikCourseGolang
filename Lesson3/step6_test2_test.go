package Lesson3

import "testing"

func TestRegisterBeforeLunch(t *testing.T) {
	scheduledDate1, scheduledDate2 := "2020-05-15 08:00:00", "2020-05-15 14:00:00"
	result1, result2 := RegisterBeforeLunch(&scheduledDate1), RegisterBeforeLunch(&scheduledDate2)
	if result1 != scheduledDate1 {
		t.Error("Wrong answer! Expected 2020-05-15 08:00:00, got ", result1)
	}
	if result2 != "2020-05-16 14:00:00" {
		t.Error("Wrong answer! Expected 2020-05-16 14:00:00, got ", result2)
	}
}
