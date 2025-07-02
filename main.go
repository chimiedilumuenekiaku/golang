package main

import (
	"fmt"
)

// Fonction Amstrong
func amstrong(nombre int) bool {
	var somme int
	var nombre_temp int
	nombre_temp = nombre

	for nombre_temp != 0 {
		chiffre := nombre_temp % 10
		nombre_temp = nombre_temp / 10
		somme += chiffre * chiffre * chiffre
	}
	if somme == nombre {
		return true
	} else {
		return false
	}

}

// Fonction demander limite
func demanderLimite() int {
	var valeur int
	fmt.Println("Donner une limilte: ")
	fmt.Scan(&valeur)
	if valeur < 1 || valeur > 1000 {
		fmt.Println("Erreur: donner une limite inferieur a 1000")
		fmt.Scan(&valeur)
	}
	return valeur
}

func main() {
	var cpteur int
	limite := demanderLimite()
	cpteur = 1
	for cpteur <= limite {
		if amstrong(cpteur) {
			fmt.Printf("%d est un nombre Amstrong\n", cpteur)
		}
		cpteur++
	}
}
