package main

import "fmt"

var students [3]string
var matrix [3][3]int

func main() {
	grades := [3]string{"a", "b", "c"}
	fmt.Println(grades)

	seq := [...]int{1, 2, 3}
	fmt.Println(seq)

	students[0] = "akshay"
	students[1] = "joseph"
	fmt.Println(students)
	fmt.Println(len(students))

	//array of arrays
	matrix[0] = [3]int{1, 2, 3}
	matrix[1] = [3]int{4, 5, 6}
	matrix[2] = [3]int{7, 8, 9}
	fmt.Println(matrix)
}
