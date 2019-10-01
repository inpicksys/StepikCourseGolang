package main

import "fmt"

func main() {
	var x, y string
	fmt.Scan(&x, &y)
	fmt.Println(deleteThatDigit(x, y))
}
