package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("hello world!", runtime.Version())
}
