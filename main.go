package main

import (
	"fmt"

	/*"github.com/luismaderas/variables"*/
	"github.com/luismaderas/ejercicios"
)

func main() {
	//fmt.Println("Hola Mundo")//
	//variables.MuestraEnteros()//
	//variables.RestoVariables()//
	/* estado, texto := variables.ConviertoaTexto(1588)
	fmt.Println(estado)
	fmt.Println(texto) */

	numero, texto := ejercicios.ConviertoaNunero("200")
	fmt.Println(numero)
	fmt.Println(texto)

}
