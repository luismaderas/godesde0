package ejercicios

import (
	"strconv"
)

func ConviertoaNunero(texto string) (int, string) {
	numero, err := strconv.Atoi(texto)
	if err != nil {
		return 0, "Hubo un error" + err.Error()
	}
	if numero > 100 {
		return numero, "Es mayor que 100"
	} else {
		return numero, "Es menor que 100"
	}

}
