package Lesson3

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func Executor() {
	reader := bufio.NewReader(os.Stdin)
	str1, err := reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	fmt.Printf("%.4f", GetQuotient(&str1))
}
