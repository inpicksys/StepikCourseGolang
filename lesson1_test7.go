package main

//TODO По данному трехзначному числу определите, все ли его цифры различны
//Выведите "YES", если все цифры числа различны, в противном случае - "NO".
import "fmt"

func sameDigits() {
	var a int
	fmt.Scan(&a)
	if (a%10) != (a%100/10) && (a%10) != (a%1000/100/10) && (a%100/10) != (a%1000/100/10) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
