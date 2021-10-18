package main

import "fmt"

func main() {
	/*Ejercicio 4 - A qué mes corresponde

	Realizar una aplicación que contenga una variable con el número del mes.
	Según el número, imprimir el mes que corresponda en texto.
	¿Se te ocurre si se puede resolver de más de una manera? ¿Cuál elegirías y por qué?
	Ej: 7, Julio
	*/

	var mes = 7

	meses := map[int]string{1: "Enero", 2: "Febrero", 3: "Marzo", 4: "Abril", 5: "Mayo", 6: "Junio", 7: "Julio", 8: "Agosto", 9: "Septiembre", 10: "Octubre", 11: "Noviembre", 12: "Diciembre"}
	fmt.Printf("%d, %v\n", mes, meses[mes])
	//Otra forma
	meses2 := []string{"Enero", "Febrero", "Marzo", "Abril", "Mayo", "Junio", "Julio", "Agosto", "Septiembre", "Octubre", "Noviembre", "Diciembre"}
	fmt.Printf("%d, %v\n", mes, meses2[mes-1])

	/*Ejercicio 5 - Listado de nombres

	Una profesor de la universidad quiere tener un listado con todos sus estudiantes. Es necesario generar una aplicación que
	contenga dicha lista.

	Estudiantes:

	Benjamin, Nahuel, Brenda, Marcos, Pedro, Axel, Alez, Dolores, Federico, Hernan, Leandro, Eduardo, Duvraschka.

	Luego de 2 clases, se sumó un estudiante nuevo. Es necesario agregarlo al listado, sin modificar el código que escribiste
	inicialmente.

	Estudiante:
	Gabriela
	*/
	estudiantes := []string{"Benjamin", "Nahuel", "Brenda", "Marcos", "Pedro", "Axel", "Alez", "Dolores", "Federico", "Hernan", "Leandro", "Eduardo", "Duvraschka"}
	fmt.Println(estudiantes)
	fmt.Println("Luego de 2 clases...")
	estudiantes = append(estudiantes, "Gabriela")
	fmt.Println(estudiantes)

	/*Ejercicio 6 - Qué edad tiene...
	Un empleado de una empresa quiere saber el nombre y edad de uno de sus empleados. Según el siguiente mapa, ayuda  a
	imprimir la edad de Benjamin.

	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}

	Por otro lado también es necesario:
	Saber cuántos de sus empleados son mayores a 21 años.
	Agregar un empleado nuevo a la lista, llamado Federico que tiene 25 años.
	Eliminar a Pedro del mapa.
	*/

	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}
	fmt.Printf("La edad de Benjamin es: %d", employees["Benjamin"])

	cantidad := 0
	for _, v := range employees {
		if v > 21 {
			cantidad++
		}
	}
	fmt.Printf("\nCantidad de empleados con mas de 21 años: %d ", cantidad)

	//Agregar a Federico con 25 años

	employees["Federico"] = 25
	fmt.Println(employees)

	//Eliminar a Pedro

	delete(employees, "Pedro")
	fmt.Println(employees)
}
