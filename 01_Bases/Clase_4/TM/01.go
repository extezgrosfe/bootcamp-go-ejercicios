package main

import "fmt"

/*Ejercicio 1 - Impuestos de salario #1
En tu función “main”, define una variable llamada “salary” y asignale un valor de tipo “int”.
Crea un error personalizado con un struct que implemente “Error()” con el mensaje “error:
el salario ingresado no alcanza el mínimo imponible" y lánzalo en caso de que “salary” sea menor a 150.000.
Caso contrario, imprime por consola el mensaje “Debe pagar impuesto”.
*/

type miErroSalary struct {
	mensaje string
}

func (e *miErroSalary) Error() string {
	return fmt.Sprintf("Mensaje: %s", e.mensaje)
}

func miErrorTest(salary int) error {
	if salary < 150000 {
		return &miErroSalary{"el salario ingresado no alcanza el mínimo imponible"}
	}
	return nil
}

func main() {
	salary := 25000
	err := miErrorTest(salary)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Debe pagar impuesto")
	}
}
