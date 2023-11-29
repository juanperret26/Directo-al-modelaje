// Crear struct y un new producto con los metodos que sean necesario
package dto

import (
	"time"

	"github.com/juanperret26/Directo-al-modelaje/go/model"
	"github.com/juanperret26/Directo-al-modelaje/go/utils"
)

type Producto struct {
	Id string `json:"id"`

	Nombre        string    `json:"nombre"`
	TipoProducto  string    `json:"tipoProducto"`
	Peso_unitario float64   `json:"peso_unitario"`
	Precio        float64   `json:"precio"`
	Stock         int       `json:"stock"`
	Stock_minimo  int       `json:"stock_minimo"`
	Actualizacion time.Time `json:"actualizacion"`
	Creacion      time.Time `json:"creacion"`
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
