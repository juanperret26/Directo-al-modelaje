// Crear struct y un new pedido con los metodos que sean necesario
package dto

import (
	"time"

	"github.com/juanperret/Directo-al-modelaje/model"
	"github.com/juanperret/Directo-al-modelaje/utils"
)

type Pedido struct {
	Id              string           `json:"id"`
	Estado          string           `json:"estado"`
	Fecha_creacion  time.Time        `json:"fecha_creacion"`
	PedidoProductos []PedidoProducto `json:"pedido_productos"`
	Actualizacion   time.Time        `json:"actualizacion"`
	Destino         string           `json:"destino"`
}

func NewPedido(pedido model.Pedido) *Pedido {
	return &Pedido{
		Id:              utils.GetStringIDFromObjectID(pedido.Id),
		Estado:          pedido.Estado,
		Fecha_creacion:  time.Now(),
		PedidoProductos: NewProductosPedido(pedido.PedidoProductos),
		Actualizacion:   time.Now(),
		Destino:         pedido.Destino,
	}
}
func (pedido Pedido) GetModel() model.Pedido {
	return model.Pedido{
		Id:              utils.GetObjectIDFromStringID(pedido.Id),
		Estado:          pedido.Estado,
		Fecha_creacion:  pedido.Fecha_creacion,
		PedidoProductos: pedido.getProductosElegidos(),
		Actualizacion:   pedido.Actualizacion,
		Destino:         pedido.Destino,
	}
}

// Hacer una funcion privada para transformar el model en dto.
func (pedido Pedido) getProductosElegidos() []model.PedidoProducto {
	var productosElegidos []model.PedidoProducto
	for _, producto := range pedido.PedidoProductos {
		productosElegidos = append(productosElegidos, producto.GetModel())
	}
	return productosElegidos
}

// Metodo para convertir una lista de ProductoPedido del modelo a una lista de ProductoPedido del dto
func NewProductosPedido(productosElegidos []model.PedidoProducto) []PedidoProducto {
	var productosElegidosDto []PedidoProducto
	for _, producto := range productosElegidos {
		productosElegidosDto = append(productosElegidosDto, *NewPedidoProducto(producto))
	}
	return productosElegidosDto
}
