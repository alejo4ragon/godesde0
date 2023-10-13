package ejercicios

import (
	"fmt"
	"strconv"
)

func StringToInt(theString string) (int, string) {

	number, err := strconv.Atoi(theString)

	var responseKindNumber string

	if err != nil {
		fmt.Println("Error al convertir la cadena a entero:", err)
	} else {
		if number > 100 {
			responseKindNumber = "Es mayor a 100"
		} else {
			responseKindNumber = "Es menor a 100"
		}
	}

	return number, responseKindNumber
}
