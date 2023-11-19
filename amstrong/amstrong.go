package main

import
	"fmt"


func demanderLimite() int {
	var valeur int
	fmt.Println("Donner une limilte: ")
	fmt.Scan(&valeur)
	if valeur < 1 || valeur > 1000 {
		fmt.Println("Erreur: donner une limite")
	}
	return valeur
}

func amstrong(n int) bool {
	var result bool
	if n == sommeCubes(n) {
		result = true
	} else {
		result = false
	}
	return result
}

func sommeCubes(n int) int {
	var (
		centaine, unite, dizaine, reste, result int
	)
	centaine = n / 100
	reste = n % 100
	dizaine = n / 10
	unite = reste % 10
	result = cube(centaine) + cube(dizaine) + cube(unite)
	return result
}
func cube(n int) int {
	var result int
	result = n * n * n

	return result
}
func main() {
	var (
		limite, cpteur int
	)
	limite = demanderLimite()
	cpteur = 1
	for cpteur < limite {
		if amstrong(cpteur) {
			fmt.Println(cpteur)
		}
	}
}
