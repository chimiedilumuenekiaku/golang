package main

import (
	"fmt"
)
func demanderNombre() int{
  var reponse int
  fmt.Print("Quel est le nombre?: ")
  fmt.Scan(&reponse)
  
  return reponse
}
func afficherTableau(n int){
  var compteur int
  compteur = 1
  for compteur <= 10 {
    fmt.Println(compteur," x ",n," = ",compteur*n)
    compteur = compteur + 1
  }
}
func main(){
	var nombre int
  nombre = demanderNombre()
  afficherTableau(nombre)
}
