package main

import "fmt"

type Example struct {
	a string
	b int
	c bool
}

func main() {
	myContact := newContact("Winslet", 12, map[string]string{"fixe": "01.54.14.25",
		"Portable": "06.56.24.15.12"})

	myNewContact := newContact("Otniel", 30, map[string]string{"fixe": "01.52.41.25",
		"Portable": "07.56.24.57.62"})

	fmt.Println(myContact)
	fmt.Println(myNewContact.displayContact())
}
