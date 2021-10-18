package main

import (
	"errors"
	"fmt"
)

/*Ejercicio 5 - Calcular cantidad de alimento

Un refugio de animales necesita calcular cuánto alimento debe comprar para las mascotas.
Por el momento solo tienen tarántulas, hamsters, perros, y gatos, pero se espera que puedan
haber muchos más animales que refugiar.

Por perro necesitan 10 kg de alimento
Por gato 5 kg
Por cada Hamster 250 gramos.
Por Tarántula 150 gramos.


Se solicita:
Implementar una función Animal que reciba como parámetro un valor de tipo texto con el animal especificado y
que retorne una función y un error (en caso que no exista el animal)
Una función para cada animal que calcule la cantidad de alimento en base a la cantidad del tipo de animal especificado.


ejemplo:


const (
   perro = "perro"
   gato = "gato"
)

...

animalPerro, err := Animal(perro)
animalGato, err := Animal(gato)

...

var cantidad float64
cantidad += animalPerro(5)
cantidad += animalGato(8)
*/

const (
	perro     = "perro"
	gato      = "gato"
	hamster   = "hamster"
	tarantula = "tarantula"
)

func comidaPerro(cantidad float64) float64 {
	return cantidad * 10
}
func comidaGato(cantidad float64) float64 {
	return cantidad * 5
}
func comidaHamster(cantidad float64) float64 {
	return cantidad * 0.25
}
func comidaTarantula(cantidad float64) float64 {
	return cantidad * 0.15
}

func animal(tipo string) (func(cantidad float64) float64, error) {
	switch tipo {
	case perro:
		return comidaPerro, nil
	case gato:
		return comidaGato, nil
	case hamster:
		return comidaHamster, nil
	case tarantula:
		return comidaTarantula, nil
	default:
		return nil, errors.New("No existe el animal")
	}
}

func main() {
	var cantidad float64

	animalito := perro

	funcionAnimal, err := animal(animalito)
	if err == nil {
		cantidad += funcionAnimal(5)
	} else {
		fmt.Printf("El animal %s no existe\n", animalito)
	}

	animalito = gato

	funcionAnimal, err = animal(animalito)
	if err == nil {
		cantidad += funcionAnimal(1)
	} else {
		fmt.Printf("El animal %s no existe\n", animalito)
	}

	animalito = hamster
	funcionAnimal, err = animal(animalito)
	if err == nil {
		cantidad += funcionAnimal(2)
	} else {
		fmt.Printf("El animal %s no existe\n", animalito)
	}

	animalito = tarantula
	funcionAnimal, err = animal(animalito)
	if err == nil {
		cantidad += funcionAnimal(1)
	} else {
		fmt.Printf("El animal %s no existe\n", animalito)
	}

	animalito = "pepito"
	funcionAnimal, err = animal(animalito)
	if err == nil {
		cantidad += funcionAnimal(5)
	} else {
		fmt.Printf("El animal %s no existe\n", animalito)
	}

	fmt.Printf("\nLa cantidad de alimento total es de %.2f Kg", cantidad)

}
