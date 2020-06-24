package Lesson3

import "fmt"

func Executor() {
	/*
		reader := bufio.NewReader(os.Stdin)
		str1, err := reader.ReadString('\n')
		if err != nil && err != io.EOF {
			panic(err)
		}
		fmt.Printf("%.4f", GetQuotient(&str1))
	*/
	var num uint
	fmt.Scan(&num)
	fmt.Println(Lambda(num))
}
