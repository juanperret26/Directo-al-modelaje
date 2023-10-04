// Crear struct y un new producto con los metodos que sean necesario
package dto

import (
	"time"

	"github.com/juanperret/Directo-al-modelaje/model"
	"github.com/juanperret/Directo-al-modelaje/utils"
)

type Producto struct {
	Id_Producto          string
	Nombre               string
	Peso_unitario        float64
	Precio               float64
	Stock                int
	Stock_minimo         int
	Tipo                 string
	Ultima_actualizacion time.Time
	Fecha_creacion       time.Time
}

func NewProducto(producto model.Producto) *Producto {
	return &Producto{
		Id_Producto:          utils.GetStringIDFromObjectID(producto.Id_Producto),
		Nombre:               producto.Nombre,
		Peso_unitario:        producto.Peso_unitario,
		Precio:               producto.Precio,
		Stock:                producto.Stock,
		Stock_minimo:         producto.Stock_minimo,
		Tipo:                 producto.Tipo,
		Ultima_actualizacion: producto.Ultima_actualizacion,
		Fecha_creacion:       time.Now(),
	}
}
func (producto Producto) GetModel() model.Producto {
	return model.Producto{
		Id_Producto:          utils.GetObjectIDFromStringID(producto.Id_Producto),
		Nombre:               producto.Nombre,
		Peso_unitario:        producto.Peso_unitario,
		Precio:               producto.Precio,
		Stock:                producto.Stock,
		Stock_minimo:         producto.Stock_minimo,
		Tipo:                 producto.Tipo,
		Ultima_actualizacion: producto.Ultima_actualizacion,
		Fecha_creacion:       producto.Fecha_creacion,
	}
}
