package Lesson3

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func Executor() {
	reader := bufio.NewReader(os.Stdin)
	dur, err := reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	dur = strings.Trim(dur, "\n")
	fmt.Println(UnixDateFromBaseDatePlusDuration(&dur))
}
