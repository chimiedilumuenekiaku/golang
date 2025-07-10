package main

import "fmt"

type Example struct {
	a string
	b int
	c bool
}

func main() {
	myExample := Example{
		a: "Alex",
		b: 30,
		c: true,
	}

	footExample := Example{"Chimie", 40, true} // This will not

	// Accessing fields of the struct
	fmt.Println(myExample)
	fmt.Println(footExample)

	myContact := newContact("Winslet", 12)
	fmt.Println(myContact)
}
