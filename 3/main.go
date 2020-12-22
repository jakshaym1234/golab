package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Hello" + "1st")
	fmt.Println("2nd")
	fmt.Println(runtime.NumCPU())
	fmt.Println(runtime.NumCPU() + 2)
}
