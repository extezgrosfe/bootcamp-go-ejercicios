package main

import "fmt"

/*Ejercicio 3 - Calcular salario
Una empresa marinera necesita calcular el salario de sus empleados basándose en la cantidad de horas trabajadas
 por mes y la categoría.

Si es categoría C, su salario es de $1.000 por hora
Si es categoría B, su salario es de $1.500 por hora más un %20 de su salario mensual
Si es de categoría A, su salario es de $3.000 por hora más un %50 de su salario mensual

Se solicita generar una función que reciba por parámetro la cantidad de minutos trabajados por mes y la categoría,
 y que devuelva su salario.

*/

func salario(minutos float64, cat string) float64 {
	horas := minutos / 60
	switch cat {
	case "A":
		return horas * 3000 * 1.5
	case "B":
		return horas * 1500 * 1.2
	default:
		return horas * 1000
	}
}

func main() {
	fmt.Println(salario(600, "A"))
	fmt.Println(salario(600, "B"))
	fmt.Println(salario(600, "C"))
}
