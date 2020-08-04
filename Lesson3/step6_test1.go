package Lesson3

import (
	"time"
)

func ConvertTimeToUnixFormat(dateAndTime string) string {
	srcTime, err := time.Parse(time.RFC3339, dateAndTime)
	if err != nil {
		panic(err)
	}
	return srcTime.Format(time.UnixDate)
}
