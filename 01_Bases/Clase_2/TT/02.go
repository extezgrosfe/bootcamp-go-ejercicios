package main

import "fmt"

/*Ejercicio 2 - Matrix
Una empresa de inteligencia artificial necesita tener una funcionalidad para crear una
estructura que represente una matriz de datos.
Para ello requieren una estructura Matrix que tenga los métodos:
Set:  Recibe una serie de valores de punto flotante e inicializa los valores en la estructura Matrix
Print: Imprime por pantalla la matriz de una formas más visible (Con los saltos de línea entre filas)
La estructura Matrix debe contener los valores de la matriz, la dimensión del alto, la dimensión del ancho,
si es cuadrática y cuál es el valor máximo.
*/

type Matrix struct {
	valores      []float64
	filas        int
	columnas     int
	esCuadratica bool
	max          float64
}

func (m *Matrix) Set(f, c int, nums ...float64) {

	if f <= 0 || c <= 0 {
		fmt.Println("Las dimensiones de la matriz no pueden ser menores o iguales a 0")
	}
	if len(nums) != f*c {
		fmt.Printf("La cantidad de elementos %d no coincide la cantidad definida: %d", len(nums), f*c)
	} else {

		m.columnas = c
		m.filas = f
		m.valores = nums
		m.esCuadratica = (f == c)
		m.max = nums[0]
		for _, v := range nums {
			if v > m.max {
				m.max = v
			}
		}
	}
}

func (m Matrix) Print1() {
	for i := 0; i < m.filas; i++ {
		fmt.Printf("\n")
		for j := 0; j < m.columnas; j++ {
			fmt.Printf("%5.2f ", m.valores[i*m.columnas+j])
		}
	}
}

func main() {
	var mat1 Matrix
	mat1.Set(4, 3, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12)
	fmt.Println(mat1)
	mat1.Print1()
	mat1.Set(3, 4, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12)
	fmt.Println(mat1)
	mat1.Print1()
	mat1.Set(3, 3, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Println(mat1)
	mat1.Print1()
	mat1.Set(3, 3, 1, 2, 3, 4, 5, 6, 7, 8)
	fmt.Println(mat1)

	mat1.Print1()
}
