package main

import (
	"encoding/json"
	"fmt"
	"os"
)

/*Ejercicio 1 - Guardar archivo
Una empresa que se encarga de vender productos de limpieza necesita:
Implementar una funcionalidad para guardar un archivo de texto,
con la información de productos comprados, separados por punto y coma.
Debe tener el id del producto, precio y la cantidad.
Estos valores pueden ser hardcodeados o escritos en duro en una variable.
*/

type Producto struct {
	Id       int     `json:"id"`
	Precio   float64 `json:"precio"`
	Cantidad float64 `json:"cantidad"`
}

func main() {
	var misProductos []Producto

	p1 := Producto{1, 25.3, 25}
	p2 := Producto{2, 35.3, 12}
	p3 := Producto{3, 45.3, 36}
	p4 := Producto{4, 55.3, 96}
	p5 := Producto{5, 65.3, 12}

	misProductos = append(misProductos, p1, p2, p3, p4, p5)

	var salida string

	for _, producto := range misProductos {
		str, _ := json.Marshal(producto)
		dato := string(str)
		salida += fmt.Sprint(dato, ";")
	}

	// eliminar ultimo punto y coma si se desea
	//salida = salida[:len(salida)-1]

	os.WriteFile("./salida_1.txt", []byte(salida), 0666)

	b, _ := json.Marshal(misProductos)

	os.WriteFile("./salida_2.txt", b, 0666)
}
