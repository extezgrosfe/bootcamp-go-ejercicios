package main

import (
	"fmt"
)

/*Ejercicio 3 - productos
Varias tiendas de ecommerce necesitan realizar una funcionalidad en Go para administrar productos y retornar el valor del precio total.
Las empresas tienen 3 tipos de productos:
Pequeño, Mediano y Grande. (Se espera que sean muchos más)
Existen costos adicionales por mantener el producto en el almacén de la tienda, y costos de envío.

Sus costos adicionales son:
Pequeño: El costo del producto (sin costo adicional)
Mediano: El costo del producto + un 3% por mantenerlo en existencia en el almacén de la tienda.
Grande: El costo del producto + un 6%  por mantenimiento, y un costo adicional  por envío de $2500.

Requerimientos:
Crear una estructura “tienda” que guarde una lista de productos.
Crear una estructura “producto” que guarde el tipo de producto, nombre y precio
Crear una interface “Ecommerce” que tenga los métodos “Precio” y “Agregar”.
Se requiere una función “nuevoproducto” que reciba el tipo de producto, su nombre y precio y devuelva la estructura del producto.
Se requiere una función “nuevaTienda” que devuelva la estructura de la tienda.
Interface Ecommerce:
 - El método “Precio” debe retornar el precio total en base al costo total de los productos y los adicionales si los hubiera.
 - El método “Agregar” debe recibir un producto y añadirlo a la lista de la tienda.

*/

const (
	pequenio = "PEQUEÑO"
	mediano  = "MEDIANO"
	grande   = "GRANDE"
)

type producto struct {
	Nombre string
	Precio float64
	Tipo   string
}

type Tienda struct {
	productos []Producto
}

type Producto interface {
	CalcularCosto() float64
}

type Ecommerce interface {
	Total() float64
	Agregar(p Producto)
}

func nuevoProducto(nombre string, precio float64, tipo string) Producto {
	return &producto{nombre, precio, tipo}
}

func nuevaTienda() Ecommerce {
	return &Tienda{}
}

func (p *producto) CalcularCosto() float64 {
	var costo float64

	switch p.Tipo {
	case pequenio:
		costo = p.Precio
	case mediano:
		costo = p.Precio * 1.3
	case grande:
		costo = (p.Precio * 1.6) + 2500
	default:
		costo = p.Precio
	}

	return costo
}

func (t *Tienda) Total() float64 {
	var total float64

	for _, p := range t.productos {
		total += p.CalcularCosto()
	}

	return total
}

func (t *Tienda) Agregar(p Producto) {
	t.productos = append(t.productos, p)
}

func main() {
	p1 := nuevoProducto("Televisor", 12000, grande)
	p2 := nuevoProducto("Celular", 3000, mediano)
	p3 := nuevoProducto("Parlante", 1200, mediano)
	p4 := nuevoProducto("Reloj", 2000, pequenio)

	tienda := nuevaTienda()

	tienda.Agregar(p1)
	tienda.Agregar(p2)
	tienda.Agregar(p3)
	tienda.Agregar(p4)

	total := tienda.Total()

	fmt.Println("El precio total es: $", total)
}
