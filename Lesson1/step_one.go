package Lesson1

import (
	"fmt"
)

func Executor() {
	var x, y string
	fmt.Scan(&x, &y)
	fmt.Println(deleteThatDigit(x, y))
}
