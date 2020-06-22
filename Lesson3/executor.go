package Lesson3

import "fmt"

func Executor() {
	var str1 string
	fmt.Scan(&str1)
	fmt.Println(GetQuotient(&str1))
}
