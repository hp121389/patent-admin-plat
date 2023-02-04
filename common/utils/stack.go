package utils

import (
	"fmt"
	"runtime"
)

func PrintStack() {
	var buf [10000]byte
	n := runtime.Stack(buf[:], false)
	fmt.Printf("==> %s\n", string(buf[:n]))
}
