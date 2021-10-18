package main

import "fmt"

/*Ejercicio 1 - Impuestos de salario
Una empresa de chocolates necesita calcular el impuesto de sus empleados al momento de depositar el sueldo,
 para cumplir el objetivo es necesario crear una función que devuelva el impuesto de un salario.
Teniendo en cuenta que si la persona gana más de $50.000 se le descontará un 17% del sueldo y si gana más de $150.000
 se le descontará además un 10%.
*/

func descuento(salario float64) float64 {
	switch {
	case salario > 150000:
		return salario * 0.27
	case salario > 50000:
		return salario * 0.17
	default:
		return 0
	}
}

func main() {

	fmt.Println(descuento(50000))
	fmt.Println(descuento(60000))
	fmt.Println(descuento(150000))
	fmt.Println(descuento(15000))

}
