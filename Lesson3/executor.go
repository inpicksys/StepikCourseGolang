package Lesson3

import (
	"fmt"
	"io/ioutil"
	"os"
)

func Executor() {
	data, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", JSONWork(data))
}
