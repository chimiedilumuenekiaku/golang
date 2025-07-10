package main

import "fmt"

func main() {
	myContact := newContact("Winslet", 12, map[string]string{"fixe": "01.54.14.25",
		"Portable": "06.56.24.15.12"})

	fmt.Println(myContact.displayContact())

	myNewContact := newContact("Otniel", 30, map[string]string{"fixe": "01.52.41.25",
		"Portable": "07.56.24.57.62"})

	fmt.Println(myNewContact.displayContact())

	myNewNewContact := newContact("Magali", 28, map[string]string{"fixe": "01.22.46.55",
		"Portable": "06.55.24.59.63"})

	fmt.Println(myNewNewContact.displayContact())
}
