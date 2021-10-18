package main

import (
	"encoding/json"
	"fmt"
)

/*Ejercicio 1 - Registro de estudiantes
Una universidad necesita registrar a los/as estudiantes y generar una funcionalidad para imprimir el detalle de los datos de cada uno de ellos/as, de la siguiente manera:

Nombre: [Nombre del alumno]
Apellido: [Apellido del alumno]
DNI: [DNI del alumno]
Fecha: [Fecha ingreso alumno]

Los valores que están en corchetes deben ser reemplazados por los datos brindados por los alumnos/as.
Para ello es necesario generar una estructura Alumnos con las variables Nombre, Apellido, DNI, Fecha y que tenga un método detalle
*/

type Alumno struct {
	Nombre   string `json:"Nombre"`
	Apellido string `json:"Apellido"`
	DNI      string `json:"DNI"`
	Fecha    string `json:"Fecha"`
}

func (a Alumno) Detalle() {
	alumnoJson, _ := json.Marshal(a)
	fmt.Println(string(alumnoJson))
}

func (a Alumno) Detalle2() {

	fmt.Printf("\nNombre: %s", a.Nombre)
	fmt.Printf("\nApellido: %s", a.Apellido)
	fmt.Printf("\nDNI: %s", a.DNI)
	fmt.Printf("\nFecha: %s", a.Fecha)

}

func main() {
	ma := Alumno{"Homero", "Simpson", "4043243", "24/04/1956"}
	ma.Detalle()
	ma.Detalle2()
}
