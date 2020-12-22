package main

import "fmt"

var sliceb []int
var slicec []int

func main() {
	slicea := []int{3, 4, 5}
	fmt.Println(slicea)
	fmt.Println(len(slicea))
	fmt.Println(cap(slicea))

	sliceb = []int{10, 11, 12}
	fmt.Println(sliceb)

	slicec = make([]int, 5, 10)
	fmt.Println(slicec)
	fmt.Println(cap(slicec))
	fmt.Println(len(slicec))

	slicec = append(slicec, 1, 2, 3, 4, 5)
	fmt.Println(slicec)
}
