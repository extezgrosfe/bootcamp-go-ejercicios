package main

import (
	"errors"
	"fmt"
)

/*Ejercicio 2 - Impuestos de salario #2

Haz lo mismo que en el ejercicio anterior pero reformulando el código para que, en reemplazo de “Error()”,
se implemente “errors.New()”
*/

func main() {
	salary := 25000

	if salary < 150000 {
		err := errors.New("Mensaje: el salario ingresado no alcanza el mínimo imponible")
		fmt.Println(err)
	} else {
		fmt.Println("Debe pagar impuesto")
	}
}
