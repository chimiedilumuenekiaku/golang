package main

import "fmt"

func update(pointeurOfNumber *int, value int) {
	*pointeurOfNumber = value
}

func main() {
	// A: string, int, bool, float,array
	// B: maps, functions
	number := 10
	fmt.Println(number)
	// Pour recupérer l'adresse mémoire de la variable number
	fmt.Println("Addresse memoire de number:", &number)

	myPointer := &number // myPointer est un pointeur vers number
	fmt.Println("Valeur de myPointer:", myPointer)

	fmt.Printf("Valeur de l'adresse mémoire %v est %d. \n", myPointer, *myPointer)

	update(myPointer, 100)
	fmt.Printf("Valeur de l'adresse mémoire %v a changé-> %d.\n", myPointer, number)

}
