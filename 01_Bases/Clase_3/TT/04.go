package main

/*Ejercicio 4 - Ordenamiento
Una empresa de sistemas requiere analizar qué algoritmos de ordenamiento utilizar para sus servicios.
Para ellos se requiere instanciar 3 arreglos con valores aleatorios desordenados
un arreglo de números enteros con 100 valores
un arreglo de números enteros con 1.000 valores
un arreglo de números enteros con 1.000.000 valores

Para instanciar las variables utilizar rand
package main

import (
   "math/rand"
)


func main() {
   variable1 := rand.Perm(100)
   variable2 := rand.Perm(1000)
   variable3 := rand.Perm(1000000)
}

Se debe realizar el ordenamiento de cada una por:
Ordenamiento por inserción
Ordenamiento por burbuja
Ordenamiento por selección

Una go routine por cada ejecución de ordenamiento
Debo esperar a que terminen los ordenamientos de 100 números para seguir el de 1.000 y después el de 1.000.000.
Por último debo medir el tiempo de cada uno y mostrar en pantalla el resultado, para saber qué ordenamiento fue mejor para cada arreglo
*/
import (
	"encoding/json"
	"fmt"
	"math/rand"
	"time"
)

func burbuja(num []int) {
	for i := 0; i < len(num)-1; i++ {
		for j := len(num) - 1; j >= i+1; j-- {
			if num[j-1] > num[j] {
				num[j-1], num[j] = num[j], num[j-1]
			}
		}
	}
}
func insercion(num []int) {
	var aux, j int
	for i := 1; i < len(num); i++ {
		aux = num[i]
		for j = i - 1; j >= 0 && aux < num[j]; j-- {
			num[j+1] = num[j]
		}
		num[j+1] = aux
	}
}

func seleccion(num []int) {
	var menor, i, j, indice int
	for i = 0; i < len(num)-1; i++ {
		menor = num[i]
		indice = i
		for j = i + 1; j < len(num); j++ {
			if menor > num[j] {
				menor = num[j]
				indice = j
			}
		}
		num[i], num[indice] = num[indice], num[i]
	}
}

type Tiempo struct {
	Tamanio int
	Bur     float64
	Inser   float64
	Selec   float64
}

func tiempoBur(num []int, t chan float64) {
	ini := time.Now()
	burbuja(num)
	fin := time.Now()
	t <- (fin.Sub(ini).Seconds())
}

func tiempoSel(num []int, t chan float64) {
	ini := time.Now()
	seleccion(num)
	fin := time.Now()
	t <- (fin.Sub(ini).Seconds())
}

func tiempoInser(num []int, t chan float64) {
	ini := time.Now()
	insercion(num)
	fin := time.Now()
	t <- (fin.Sub(ini).Seconds())
}

func main() {
	rand.Seed(int64(time.Now().UnixNano()))
	var arreglo0, arreglo1, arreglo2 []int
	var cant []int = []int{100, 1000, 10000}
	var tiem Tiempo
	var tiempos []Tiempo

	tiempos = []Tiempo{}
	iniGen := time.Now()
	for i := 0; i < 3; i++ {
		arreglo0 = rand.Perm(cant[i])
		arreglo1 = rand.Perm(cant[i])
		arreglo2 = rand.Perm(cant[i])
		copy(arreglo1, arreglo0)
		copy(arreglo2, arreglo0)
		tiem.Tamanio = cant[i]

		t1 := make(chan float64)
		go tiempoBur(arreglo0, t1)

		t2 := make(chan float64)
		go tiempoInser(arreglo1, t2)

		t3 := make(chan float64)
		go tiempoSel(arreglo2, t3)

		tiem.Bur = <-t1
		tiem.Inser = <-t2
		tiem.Selec = <-t3

		tiempos = append(tiempos, tiem)

		//	b, _ := json.Marshal(tiem)
		//fmt.Println(string(b))
	}

	finGen := time.Now()

	//fmt.Println(tiempos)
	for _, v := range tiempos {
		b, _ := json.Marshal(v)
		fmt.Println(string(b))
	}
	fmt.Println("Tardo: ", finGen.Sub(iniGen).Seconds())

}
