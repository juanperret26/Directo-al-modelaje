package dto

import (
	"github.com/juanperret/Directo-al-modelaje/model"
	"github.com/juanperret/Directo-al-modelaje/utils"
)

type PedidoProducto struct {
	Id              string
	Nombre          string
	Cantidad        int
	Precio_unitario float64
	Stock           int
	Tipo            string
}

func NewPedidoProducto(producto model.Producto) *PedidoProducto {
	return &PedidoProducto{
		Id:              utils.GetStringIDFromObjectID(producto.Id),
		Nombre:          producto.Nombre,
		Cantidad:        0,
		Precio_unitario: producto.Precio,
		Stock:           producto.Stock,
		Tipo:            producto.Tipo,
	}
}

func GetModel(producto model.Producto) *PedidoProducto {
	return &PedidoProducto{
		Id:              utils.GetStringIDFromObjectID(producto.Id),
		Nombre:          producto.Nombre,
		Cantidad:        0,
		Precio_unitario: producto.Precio,
		Stock:           producto.Stock,
		Tipo:            producto.Tipo,
	}
}
