// Crear struct y un new producto con los metodos que sean necesario
package dto

import (
	"time"

	"github.com/juanperret/Directo-al-modelaje/model"
	"github.com/juanperret/Directo-al-modelaje/utils"
)

type Producto struct {
	Id string

	Nombre        string
	TipoProducto  string
	Peso_unitario float64
	Precio        float64
	Stock         int
	Stock_minimo  int
	Actualizacion time.Time
	Creacion      time.Time
}

func NewProducto(producto model.Producto) *Producto {
	return &Producto{
		Id:            utils.GetStringIDFromObjectID(producto.Id),
		Nombre:        producto.Nombre,
		TipoProducto:  producto.TipoProducto,
		Peso_unitario: producto.Peso_unitario,
		Precio:        producto.Precio,
		Stock:         producto.Stock,
		Stock_minimo:  producto.Stock_minimo,
		Actualizacion: time.Now(),
		Creacion:      time.Now(),
	}
}
func (producto Producto) GetModel() model.Producto {
	return model.Producto{
		Id:            utils.GetObjectIDFromStringID(producto.Id),
		Nombre:        producto.Nombre,
		TipoProducto:  producto.TipoProducto,
		Peso_unitario: producto.Peso_unitario,
		Precio:        producto.Precio,
		Stock:         producto.Stock,
		Stock_minimo:  producto.Stock_minimo,
		Actualizacion: producto.Actualizacion,
		Creacion:      producto.Creacion,
	}
}
