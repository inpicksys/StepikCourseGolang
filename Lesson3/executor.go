package Lesson3

import "fmt"

func Executor() {
	stringToConvert := ""
	fmt.Scan(&stringToConvert)
	fmt.Println(RegisterBeforeLunch(&stringToConvert))
}
