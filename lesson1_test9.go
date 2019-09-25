package main

import "fmt"

func luckyTicket() bool {
	//TODO Определите является ли билет счастливым. Счастливым считается билет,
	// в шестизначном номере которого сумма первых трёх цифр совпадает с суммой трёх последних.
	// Выведите "YES", если билет счастливый, в противном случае - "NO".
	var a int
	fmt.Scan(&a)
	if a < 100000 {
		return false
	}
	if (a/100000 + a%100000/10000 + a%10000/1000) == (a%10 + a%100/10 + a%1000/100) {
		return true
	}
	return false
}
