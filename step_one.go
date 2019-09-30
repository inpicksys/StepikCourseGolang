package main

import "fmt"

func main() {
	var x, y int
	fmt.Scan(&x, &y)
	fmt.Println(dividedBy7MaxOnInterval(x, y))
}
