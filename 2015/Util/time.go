package util

import (
	"fmt"
	"time"
)

func Time(function func(), part string) {
	start := time.Now()
	function()
	elapsed := time.Since(start)
	fmt.Printf("%s took %s\n", part, elapsed)
}
