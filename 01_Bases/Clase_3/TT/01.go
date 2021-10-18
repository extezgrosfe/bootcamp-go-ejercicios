package main

import "fmt"

/*Ejercicio 1 - Red social
Una empresa de redes sociales requiere implementar una estructura usuario con
funciones que vayan agregando información a la estructura.
Para optimizar y ahorrar memoria requieren que la estructura usuarios ocupe el mismo
lugar en memoria para el main del programa y para las funciones:
La estructura debe tener los campos: Nombre, Apellido, edad, correo y contraseña
Y deben implementarse las funciones:
cambiar nombre: me permite cambiar el nombre y apellido.
cambiar edad: me permite cambiar la edad.
cambiar correo: me permite cambiar el correo.
cambiar contraseña: me permite cambiar la contraseña.
*/

type Usuario struct {
	Nombre   string `json:"nombre"`
	Apellido string `json:"Apellido"`
	Edad     int    `json:"edad"`
	Correo   string `json:"correo"`
	Clave    string `json:"clave"`
}

func (u *Usuario) cambiarNombre(newNombre, newApellido string) {
	u.Nombre = newNombre
	u.Apellido = newApellido
}
func (u *Usuario) cambiarEdad(newEdad int) {
	u.Edad = newEdad
}
func (u *Usuario) cambiarCorreo(newCorreo string) {
	u.Correo = newCorreo
}
func (u *Usuario) cambiarClave(newClave string) {
	u.Clave = newClave
}

func main() {
	miUsuario := &Usuario{"Homero", "Simpson", 39, "homero.simpson@digitalhouse.com", "marge123"}
	fmt.Println(miUsuario)
	miUsuario.cambiarNombre("Marge", "Bouvier")
	fmt.Println(miUsuario)
	miUsuario.cambiarCorreo("margen@digitalhouse.com")
	miUsuario.cambiarEdad(34)
	miUsuario.cambiarClave("lisa777")
	fmt.Println(miUsuario)
}
