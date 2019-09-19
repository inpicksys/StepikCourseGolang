//TODO Напишите программу, которая последовательно делает следующие операции с каждым числом:
//Число умножается на 2;
//Затем к числу прибавляется 100.
package main

import "fmt"

func stepOneExtendedWork() int {
	var a int
	fmt.Scan(&a)
	a *= 2
	a += 100
	return a
}
