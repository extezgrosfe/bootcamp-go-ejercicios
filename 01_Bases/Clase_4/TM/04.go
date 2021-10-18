package main

import (
	"errors"
	"fmt"
)

/*Ejercicio 4 -  Impuestos de salario #4
Vamos a hacer que nuestro programa sea un poco más complejo y útil.
Desarrolla las funciones necesarias para permitir a la empresa calcular:
Salario mensual de un trabajador según la cantidad de horas trabajadas.
En caso de que el salario mensual sea igual o superior a $150.000, se le deberá descontar
el 10% en concepto de impuesto. La función que se ocupe de realizar este cálculo deberá retornar más de un valor,
incluyendo un error en caso de que la cantidad de horas mensuales ingresadas sea menor a 80 o un número negativo.
El error deberá indicar “error: el trabajador no puede haber trabajado menos de 80 hs mensuales”.
Calcular el medio aguinaldo correspondiente al trabajador (fórmula de cálculo de aguinaldo:
[mejor salario del semestre] dividido 12 y multiplicar el [resultado obtenido] por la
[cantidad de meses trabajados en el semestre]). La función que realice el cálculo deberá retornar más de un valor,
incluyendo un error en caso de que se ingrese un número negativo.

Desarrolla el código necesario para cumplir con las funcionalidades requeridas, utilizando “errors.New()”,
“fmt.Errorf()” y “errors.Unwrap()”. No olvides realizar las validaciones de los retornos de error en tu función
“main()”.

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

func salarioMensual(salary float64, horas int) (float64, error) {

	if horas < 80 {
		return 0, &miErroSalary{"el trabajador no puede haber trabajado menos de 80 hs mensuales"}
	} else if salary >= 150000 {
		return salary * 0.9, nil
	} else {
		return salary, nil
	}
}

func main() {
	salario, err := salarioMensual(125000, 70)
	e1 := fmt.Errorf("SalaryError: %w", err)
	if errors.Unwrap(e1) != nil {
		fmt.Println(e1)
		fmt.Println(errors.Unwrap(e1))
		fmt.Println(errors.New("Revise nuevamente su salario por si tiene que cambiar"))
	} else {
		fmt.Printf("Salario final es $%.2f", salario)
	}
}
