package main

import (
	"errors"
	"fmt"
)

/*Ejercicio 2 - Calcular promedio

Un colegio de Buenos Aires necesita calcular el promedio (por alumno) de sus calificaciones.
Se solicita generar una función en la cual se le pueda pasar N cantidad de enteros y devuelva
 el promedio y un error en caso que uno de los números ingresados sea negativo
*/

func promedio(notas ...float64) (float64, error) {
	var suma float64
	for _, n := range notas {
		if n < 0 {
			return 0, errors.New("Se ingreso un numero negativo")
		}
		suma += n
	}
	return suma / float64(len(notas)), nil
}

func main() {

	fmt.Println(promedio(9, 9, 9, 10))
	fmt.Println(promedio(9, 9, -4, 10))
	fmt.Println(promedio(9, 9, 4, 10))

}
