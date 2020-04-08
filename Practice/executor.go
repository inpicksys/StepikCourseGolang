package Practice

import "fmt"

func Executor() {
	/*
		var a, b, c float64
		fmt.Scan(&a, &b)
		pifagorus(&a, &b, &c)
		fmt.Println(c)
		var text string
		fmt.Scan(&text)
		//stringSplitter(&text, &result)
		//fmt.Println(result)
		fmt.Println(findMaxInDigitRow(&text))
	*/
	var source int
	fmt.Scan(&source)
	fmt.Println(*whatTheFunc(source))
}
