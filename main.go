package main

import (
	"fmt"

	"github.com/luismaderas/variables"
)

func main() {
	//fmt.Println("Hola Mundo")//
	//variables.MuestraEnteros()//
	//variables.RestoVariables()//
	estado, texto := variables.ConviertoaTexto(1588)
	fmt.Println(estado)
	fmt.Println(texto)
}
