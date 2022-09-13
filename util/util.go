package util

import (
	"fmt"
	"strconv"
	"time"
)

func Time(function func(), part string) {
	start := time.Now()
	function()
	elapsed := time.Since(start)
	fmt.Printf("%s took %s\n", part, elapsed)
}

func Atoi(s string) int {
	num, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}

	return num
}
