package main

import (
	"Documents/estudos/golang/footypes"

	"fmt"
)

func main() {
	fmt.Println("test")
	var foovar footypes.Foo
	foovar.TypeInt = 10

	fmt.Println(foovar)
}
