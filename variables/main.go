package main

import "fmt"

func main() {
	var speed int
	var a float64
	var b bool
	var c string
	fmt.Println(
		10, -101, 0,
	)
	fmt.Println(
		1., 1., 0., -1.5,
	)
	fmt.Println(speed)
	fmt.Println(a, b, c)
	fmt.Printf("%q\n", c)
	var (
		d int
		f float64
	)
	var g,h int
	fmt.Println(d,f,g,h)


	var i bool = true
	fmt.Println(i)
	var j = true
	fmt.Println(j)
	k := false
	fmt.Println(k)
}
