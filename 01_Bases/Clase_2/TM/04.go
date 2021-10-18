package main

import (
	"errors"
	"fmt"
)

/*Ejercicio 4 - Calcular estadísticas

Los profesores de una universidad de Colombia necesitan calcular algunas estadísticas de
calificaciones de los alumnos de un curso, requiriendo calcular los valores mínimo,
máximo y promedio de sus calificaciones.

Se solicita generar una función que indique qué tipo de cálculo se quiere realizar (mínimo, máximo o promedio)
y que devuelva otra función ( y un error en caso que el cálculo no esté definido)
que se le puede pasar una cantidad N de enteros y devuelva el cálculo que se indicó en la función anterior
Ejemplo:

const (
   minimo = "minimo"
   promedio = "promedio"
   maximo = "maximo"
)

...

minFunc, err := operacion(minimo)
promFunc, err := operacion(promedio)
maxFunc, err := operacion(maximo)

...

valorMinimo := minFunc(2,3,3,4,1,2,4,5)
valorPromedio := promFunc(2,3,3,4,1,2,4,5)
valorMaximo := maxFunc(2,3,3,4,1,2,4,5)

*/

const (
	minimo   = "minimo"
	promedio = "promedio"
	maximo   = "maximo"
)

func min(nums ...int) float64 {
	mini := nums[0]
	for _, v := range nums {
		if v < mini {
			mini = v
		}
	}
	return float64(mini)
}

func max(nums ...int) float64 {
	maximo := nums[0]
	for _, v := range nums {
		if v > maximo {
			maximo = v
		}
	}
	return float64(maximo)
}

func prom(nums ...int) float64 {
	suma := 0.0
	for _, v := range nums {
		suma += float64(v)
	}
	return suma / float64(len(nums))
}

func operacion(oper string) (func(num ...int) float64, error) {
	switch oper {
	case minimo:
		return min, nil
	case maximo:
		return max, nil
	case promedio:
		return prom, nil
	default:
		return nil, errors.New("invalid operacion")
	}

}

func main() {
	minFunc, err := operacion(minimo)
	if err == nil {
		valor := minFunc(2, 3, 3, 4, 1, 2, 4, 5)
		fmt.Printf("El minimo es: %f", valor)
	}

	promFunc, err := operacion(promedio)
	if err == nil {
		valor := promFunc(2, 3, 3, 4, 1, 2, 4, 5)
		fmt.Printf("\nEl promedio es: %f", valor)
	}

	maxFunc, err := operacion(maximo)
	if err == nil {
		valor := maxFunc(2, 3, 3, 4, 1, 2, 4, 5)
		fmt.Printf("\nEl maximo es: %f", valor)
	}

}
