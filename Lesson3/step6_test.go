package Lesson3

import "testing"

func TestConvertTimeToUnixFormat(t *testing.T) {
	srcString := "1986-04-16T05:20:00+06:00"
	result := ConvertTimeToUnixFormat(srcString)
	if result != "Wed Apr 16 05:20:00 +0600 1986" {
		t.Error("Wrong answer! Expected Wed Apr 16 05:20:00 +0600 1986, got ", result)
	}
}

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

func TestGetDifferenceBetweenDates(t *testing.T) {
	date1, date2 := "13.03.2018 14:00:15", "12.03.2018 14:00:15"
	result := GetDifferenceBetweenDates(&date1, &date2)
	if result != "24h0m0s" {
		t.Error("Wrong answer! Expected 24h0m0s, got ", result)
	}
}

func TestUnixDateFromBaseDatePlusDuration(t *testing.T) {
	durS := "12 мин. 13 сек."
	result := UnixDateFromBaseDatePlusDuration(&durS)
	if result != "Fri May 15 19:28:18 UTC 2020" {
		t.Error("Wrong answer! Expected Fri May 15 19:28:18 UTC 2020, got ", result)
	}
}
