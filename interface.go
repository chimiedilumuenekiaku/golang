package main

import "fmt"

type animal interface {
	Noise() string     // bruit
	NumberOfLegs() int // nombre de pattes
}

type Cat struct {
	Name  string // nom
	Breed string // espèce
}

type Spider struct {
	Name     string // nom
	Breed    string // espèce
	Venomous bool   // venimeux
}

func PrintAnimalInfo(a animal) {
	fmt.Println("Le bruit de cet animal est ", a.Noise(), "et il possède", a.NumberOfLegs(), " pattes!")
}

func (c *Cat) Noise() string {
	return "Miaou"
}

func (c *Cat) NumberOfLegs() int {
	return 4
}

func (s *Spider) Noise() string {
	return "Hiss"
}

func (s *Spider) NumberOfLegs() int {
	return 8
}

func main() {
	cat := Cat{
		Name:  "Kitty",
		Breed: "Siamois",
	}

	PrintAnimalInfo(&cat)

	spider := Spider{
		Name:     "Spiddy",
		Breed:    "Veuve Noire",
		Venomous: true,
	}
	PrintAnimalInfo(&spider)
}
