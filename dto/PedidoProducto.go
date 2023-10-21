package dto

import (
	"github.com/juanperret/Directo-al-modelaje/model"
	"github.com/juanperret/Directo-al-modelaje/utils"
)

type PedidoProducto struct {
	Id              string
	CodigoProducto  string
	Nombre          string
	Cantidad        float64
	Precio_unitario float64
}

func NewPedidoProducto(producto model.Producto) *PedidoProducto {
	return &PedidoProducto{
		Id:              utils.GetStringIDFromObjectID(producto.Id),
		CodigoProducto:  producto.CodigoProducto,
		Nombre:          producto.Nombre,
		Cantidad:        0,
		Precio_unitario: producto.Precio,
	}
}

func GetModel(producto model.Producto) *PedidoProducto {
	return &PedidoProducto{
		Id: utils.GetStringIDFromObjectID(producto.Id),

		Nombre:          producto.Nombre,
		Cantidad:        0,
		Precio_unitario: producto.Precio,
	}
}
