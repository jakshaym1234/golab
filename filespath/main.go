package main

import (
	"fmt"
	"path"
)

func main() {
	var dir, file string
	dir, file = path.Split("/sample/sample.pdf")
	fmt.Println(dir, file)
	var file1 string
	_, file1 = path.Split("how/this.pdf")
	fmt.Println(file1)
	fmt.Println("this is the path\n", file1)

	var speeed float64
	var force int
	speeed = 25.5
	force = 7
	speeed = speeed * float64(force)
	fmt.Println(speeed)
}