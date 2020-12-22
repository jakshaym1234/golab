package main

import (
	"fmt"
	"time"
)

func main() {
	var (
		a = 10
		b = true
	)
	fmt.Println(a, b)
	// safe, speed := true, 50
	// fmt.Println(safe, speed)

	var speed int
	fmt.Println(speed)
	speed = 50
	fmt.Println(speed)
	speed = speed - 25
	fmt.Println(speed)
	var timet time.Time
	timet = time.Now()
	fmt.Println(timet)
}
