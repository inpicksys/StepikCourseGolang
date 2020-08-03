package Lesson3

import (
	"fmt"
)

func Executor() {
	dataSet := "../data-20190514T0100"
	fmt.Println(EncodeFromFileDecodeToFileCountSumAllGlobalIDs(dataSet))
}
