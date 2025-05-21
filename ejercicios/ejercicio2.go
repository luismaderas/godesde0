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

func TabladeMultiplicar() {
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
	fmt.Println("LA TABLA DEL ", numero, " ES: ")
	for i := 1; i <= 10; i++ {
		//fmt.Printf(" %d X %d  = %d \n", numero, i, i*numero)
		fmt.Println(numero, " x ", i, " = ", numero*i)
	}

}
