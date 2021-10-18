package main

import (
	"fmt"
	"math"
)

/*Ejercicio 3 - Calcular Precio
Una empresa nacional se encarga de realizar venta de productos, servicios y mantenimiento.
Para ello requieren realizar un programa que se encargue de calcular el precio total de Productos,
Servicios y Mantenimientos. Debido a la fuerte demanda y para optimizar la velocidad requieren
que el cálculo de la sumatoria se realice en paralelo mediante 3 go routines.

Se requieren 3 estructuras:
Productos: nombre, precio, cantidad.
Servicios: nombre, precio, minutos trabajados.
Mantenimiento: nombre, precio.

Se requieren 3 funciones:
Sumar Productos: recibe un array de producto y devuelve el precio total (precio * cantidad).
Sumar Servicios: recibe un array de servicio y devuelve el precio total (precio * media hora
trabajada, si no llega a trabajar 30 minutos se le cobra como si hubiese trabajado media hora).
Sumar Mantenimiento: recibe un array de mantenimiento y devuelve el precio total.

Los 3 se deben ejecutar en paralelo y al final se debe mostrar por pantalla el monto final
(sumando el total de los 3).
*/

type Producto struct {
	Nombre   string
	Precio   float64
	Cantidad float64
}

type Servicio struct {
	Nombre  string
	Precio  float64
	Minutos float64
}

type Mantenimiento struct {
	Nombre string
	Precio float64
}

func sumarProductos(prod []Producto) float64 {
	var suma float64
	for _, v := range prod {
		suma += v.Cantidad * v.Precio
	}
	return suma
}

func sumarServicios(serv []Servicio) float64 {
	var suma float64
	for _, v := range serv {
		suma += v.Precio * math.Round(v.Minutos/30.0)
	}
	return suma
}

func sumarMantenimiento(mant []Mantenimiento) float64 {
	var suma float64
	for _, v := range mant {
		suma += v.Precio
	}
	return suma
}

func sumarCostosProd(c chan float64, tipo []Producto) {
	c <- sumarProductos(tipo)
}
func sumarCostosSer(c chan float64, tipo []Servicio) {
	c <- sumarServicios(tipo)
}
func sumarCostosMan(c chan float64, tipo []Mantenimiento) {
	c <- sumarMantenimiento(tipo)
}

func main() {
	misProductos := []Producto{
		{"Heladera", 10000, 3},
		{"Cocina", 5000, 2},
	}
	misServicios := []Servicio{
		{"Lavado", 500, 300},
		{"Planchado", 100, 60},
	}
	misMantenimientos := []Mantenimiento{
		{"Mantenimiento 1", 1500},
		{"Mantenimiento 2", 1000},
		{"Mantenimiento 3", 4000},
	}

	var costoTotal float64

	canal := make(chan float64)
	go sumarCostosProd(canal, misProductos)
	go sumarCostosSer(canal, misServicios)
	go sumarCostosMan(canal, misMantenimientos)
	costoTotal += <-canal + <-canal + <-canal
	fmt.Printf("Costo total: %.2f", costoTotal)
}
