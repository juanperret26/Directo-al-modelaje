package dto

import (
	"github.com/juanperret/Directo-al-modelaje/model"
)

type PedidoProducto struct {
	CodigoProducto  string
	Nombre          string
	Cantidad        int
	Precio_unitario float64
}

func NewPedidoProductoFromPedido(producto model.Producto, cantidad int) *PedidoProducto {
	return &PedidoProducto{
		CodigoProducto:  producto.Id.String(),
		Nombre:          producto.Nombre,
		Cantidad:        cantidad,
		Precio_unitario: producto.Precio,
	}
}
func NewPedidoProducto(producto model.PedidoProducto) *PedidoProducto {
	return &PedidoProducto{
		CodigoProducto:  producto.CodigoProducto,
		Nombre:          producto.Nombre,
		Cantidad:        producto.Cantidad,
		Precio_unitario: producto.Precio_unitario,
	}
}

func (pedidoProducto PedidoProducto) GetModel() model.PedidoProducto {
	return model.PedidoProducto{
		CodigoProducto:  pedidoProducto.CodigoProducto,
		Nombre:          pedidoProducto.Nombre,
		Cantidad:        pedidoProducto.Cantidad,
		Precio_unitario: pedidoProducto.Precio_unitario,
	}
}
