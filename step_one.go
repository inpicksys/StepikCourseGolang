package main

import "fmt"

func main() {
	var x, p, y uint
	fmt.Scan(&x, &p, &y)
	result := bankPercents(x, p, y)
	fmt.Println(result)
}
