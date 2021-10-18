package main

import "fmt"

/*Ejercicio 2 - Ecommerce
Una importante empresa de ventas web necesita agregar una funcionalidad para agregar
productos a los usuarios. Para ello requieren que tanto los usuarios como los
productos tengan la misma dirección de memoria en el main del programa como en las funciones.
Se necesitan las estructuras:
Usuario: Nombre, Apellido, Correo, Productos (array de productos).
Producto: Nombre, precio, cantidad.
Se requieren las funciones:
Nuevo producto: recibe nombre y precio, y retorna un producto.
Agregar producto: recibe usuario, producto y cantidad, no retorna nada, agrega el producto al usuario.
Borrar productos: recibe un usuario, borra los productos del usuario.
*/

type Producto struct {
	Nombre   string  `json:"nombre"`
	Precio   float64 `json:"precio"`
	Cantidad float64 `json:"cantidad"`
}

type Usuario struct {
	Nombre       string     `json:"nombre"`
	Apellido     string     `json:"apellido"`
	Correo       string     `json:"correo"`
	MisProductos []Producto `json:"misproductos"`
}

func nuevoProducto(nombre string, precio float64) Producto {
	return Producto{nombre, precio, 0}
}

func (u *Usuario) agregarProducto(producto Producto, cant float64) {
	producto.Cantidad = cant
	u.MisProductos = append(u.MisProductos, producto)
}

func (u *Usuario) borrarProductos() {
	u.MisProductos = []Producto{}
}

func main() {
	usuario := &Usuario{"Bart", "Simpson", "elbarto@digitalhouse.com", []Producto{}}
	usuario.agregarProducto(nuevoProducto("Manzana", 15), 3)
	usuario.agregarProducto(nuevoProducto("Peras", 12), 2)
	usuario.agregarProducto(nuevoProducto("Limones", 25), 5)

	fmt.Println()
	usuario.borrarProductos()
	fmt.Println(usuario)
}
