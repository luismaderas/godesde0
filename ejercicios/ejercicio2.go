package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var numero int
var err error
var tabla string
var texto string

func TabladeMultiplicar() string {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("Ingrese el número de la tabla: ")
		if scanner.Scan() {
			numero, err = strconv.Atoi(scanner.Text())
			if err != nil {
				continue
			} else {
				break
			}
		}
	}
	//fmt.Println("LA TABLA DEL ", numero, " ES: ")
	for i := 1; i <= 10; i++ {
		texto += fmt.Sprintf(" %d X %d  = %d \n", numero, i, i*numero)
	}
	//fmt.Println(texto)
	return texto
}
