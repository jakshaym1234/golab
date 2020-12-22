package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Printf("%#v", os.Args)
	fmt.Println("1st", os.Args[1])
	fmt.Println("length", len(os.Args))
}
