package main

import (
	"encoding/json"
	"fmt"
	"os"
)

/*Ejercicio 2 - Leer archivo
La misma empresa necesita leer el archivo almacenado,
para ello requiere que: se imprima por pantalla mostrando los valores tabulados
, con un título (tabulado a la izquierda para el ID y a la derecha para el Precio y Cantidad),
 el precio, la cantidad y abajo del precio se debe visualizar el total (Sumando precio por cantidad)

Ejemplo:

ID                            Precio  Cantidad
111223                      30012.00         1
444321                    1000000.00         4
434321                         50.50         1
                          4030062.50

*/

type Producto struct {
	Id       int     `json:"id"`
	Precio   float64 `json:"precio"`
	Cantidad float64 `json:"cantidad"`
}

func detalle(prod []Producto) {
	var sum float64
	fmt.Printf("%-10s%20s%10s\n", "ID", "Precio", "Cantidad")
	for _, v := range prod {
		sum += v.Precio * v.Cantidad
		fmt.Printf("%-10d%20.2f%10.2f\n", v.Id, v.Precio, v.Cantidad)
	}
	fmt.Printf("%30.2f", sum)
}

func main() {
	data, err := os.ReadFile("./salida_2.txt")
	if err != nil {
		fmt.Printf("Error reading")
		os.Exit(1)
	}

	var productos []Producto
	// productos := &[]Producto{}
	json.Unmarshal(data, &productos)

	detalle(productos)

}
