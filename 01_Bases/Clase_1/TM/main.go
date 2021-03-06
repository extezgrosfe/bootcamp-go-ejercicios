package main

import "fmt"

func main() {
	/*
	   Ejercicio 1 - Imprimí tu nombre

	   Crea una aplicación donde tengas como variable tu nombre y dirección.
	   Imprime en consola el valor de cada una de las variables.
	*/
	var (
		nombre    = "Homero Simpson"
		direccion = "Avenida Siempreviva 742"
	)

	fmt.Printf("Mi nombre es: %s y vivo en: %s\n", nombre, direccion)

	/*Ejercicio 2 - Clima

	Una empresa de meteorología quiere tener una aplicación donde pueda tener la temperatura y humedad y presión atmosférica de distintos lugares.
	Declara 3 variables especificando el tipo de dato, como valor deben tener la temperatura, humedad y presión de donde te encuentres.
	Imprime los valores de las variables en consola.
	¿Qué tipo de dato le asignarías a las variables?
	*/

	var temperatura float32 = 16.5
	var humedad float32 = 78.55
	var presion float32 = 925.95

	fmt.Printf("La temperatura es %.2f con una humedad del %.2f %% a una presión atmosférica del %.2f", temperatura, humedad, presion)

	/*Ejercicio 3 - Declaración de variables

	Un profesor de programación está corrigiendo los exámenes de sus estudiantes de la materia Programación I para poder brindarles
	las correspondientes devoluciones. Uno de los puntos del examen consiste en declarar distintas variables.
	Necesita ayuda para:
	Detectar cuáles de estas variables que declaró el alumno son correctas.
	Corregir las incorrectas.
	  var 1nombre string
	  var apellido string
	  var int edad
	  1apellido := 6
	  var licencia_de_conducir = true
	  var estatura de la persona int
	  cantidadDeHijos := 2
	*/

	// var nombre string
	// var apellido string
	// var edad int
	// apellido2 := "Simpson"
	// var licenciaDeConducir = true
	// var estatura int
	// cantidadDeHijos := 2

	/*Ejercicio 4 - Tipos de datos
	Un estudiante de programación intentó realizar declaraciones de variables con sus respectivos tipos en Go pero tuvo varias dudas mientras lo hacía.
	A partir de esto, nos brindó su código y pidió la ayuda de un desarrollador experimentado que pueda:
	Verificar su código y realizar las correcciones necesarias.

	  var apellido string = "Gomez"
	  var edad int = "35"
	  boolean := "false";
	  var sueldo string = 45857.90
	  var nombre string = "Julián"
	*/

	// var apellido string = "Gomez"			(BIEN)
	// var edad int = "35"					var edad int = 35
	// boolean := "false";					boolean := "false"    o  boolean := false      interpretando que false puede ser una palabra o un booleano
	// var sueldo string = 45857.90			var sueldo float32 = 45857.90
	// var nombre string = "Julián"			(BIEN)
}
