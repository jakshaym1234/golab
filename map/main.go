package main

import "fmt"

var my2map map[string]int

func main() {
	mymap := make(map[string]int)
	mymap = map[string]int{
		"india": 1,
		"uk":    2,
		"us":    3,
	}
	fmt.Println(mymap)
	fmt.Println(mymap["india"])
	my2map = map[string]int{
		"I":  1,
		"UK": 2,
		"US": 3,
	}
	fmt.Println(my2map)

	//adding to map
	my2map["AUS"] = 4
	fmt.Println(my2map)

	//deleting from the map
	delete(my2map, "US")
	fmt.Println(my2map)

	presence, ok := my2map["UK"]
	fmt.Println(presence, ok)

	presence2, ok2 := my2map["AUSA"]
	fmt.Println(presence2, ok2)

	_, ok3 := my2map["AUSSA"]
	fmt.Println(ok3)

	var my3map map[string]int
	my3map = map[string]int{
		"MH": 1,
		"KL": 2,
	}
	fmt.Println(my3map)
}
