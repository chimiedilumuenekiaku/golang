package main

import (
	"fmt"
)

func updateA(number int) int {
	number = 5
	return number
}

func updateB(item map[string]int) {
	item["bonbon"] = 4
	item["oranges"] = 12
}

func main() {
	//A: string , int, bool, float, array
	number := 10
	number = updateA(number)
	fmt.Println(number)

	//B: maps, functions
	groceryList := map[string]int{
		"apples":  5,
		"bananas": 10,
		"oranges": 8,
	}
	updateB(groceryList)
	fmt.Println(groceryList)
}
