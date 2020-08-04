package Lesson3

import "testing"

func TestConvertTimeToUnixFormat(t *testing.T) {
	srcString := "1986-04-16T05:20:00+06:00"
	result := ConvertTimeToUnixFormat(srcString)
	if result != "Wed Apr 16 05:20:00 +0600 1986" {
		t.Error("Wrong answer! Expected Wed Apr 16 05:20:00 +0600 1986, got ", result)
	}
}
