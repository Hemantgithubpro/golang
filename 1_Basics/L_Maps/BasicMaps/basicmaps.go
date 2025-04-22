package main

import "fmt"


func main() {
	var m = make(map[string]int)
	m["Bell Labs"] = 40
	m["Google"] = 50
	m["Microsoft"] = 60
	m["Apple"] = 70

	// fmt.Println(m["Bell Labs"])
	
	for i, v := range m {
		fmt.Println(i, v)
	}
	
	m["Bell Labs"]++
	fmt.Println(m["Bell Labs"])

	exp()
	fmt.Println(exp()["english"])
}

func exp() map[string]int {
	var m map[string]int = map[string]int {"english":99, "math":98, "science":94, "history":95, "Apple": 100}
	fmt.Println(m["english"])
	return m
}