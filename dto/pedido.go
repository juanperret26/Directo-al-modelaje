// Crear struct y un new pedido con los metodos que sean necesario
package dto

import (
	"time"

	"github.com/juanperret/Directo-al-modelaje/model"
	"github.com/juanperret/Directo-al-modelaje/utils"
)

type Pedido struct {
	Id             string
	Estado         string
	Fecha_creacion time.Time
	Productos      []PedidoProducto
	Actualizacion  time.Time
	Destino        string
}

func NewPedido(pedido model.Pedido) *Pedido {
	return &Pedido{
		Id:             utils.GetStringIDFromObjectID(pedido.Id),
		Estado:         pedido.Estado,
		Fecha_creacion: time.Now(),
		Productos:      []PedidoProducto{},
		Actualizacion:  pedido.Actualizacion,
		Destino:        pedido.Destino,
	}
}
func (pedido Pedido) GetModel() model.Pedido {
	return model.Pedido{
		Id:             utils.GetObjectIDFromStringID(pedido.Id),
		Estado:         pedido.Estado,
		Fecha_creacion: pedido.Fecha_creacion,
		Productos:      []model.PedidoProducto{},
		Actualizacion:  pedido.Actualizacion,
		Destino:        pedido.Destino,
	}
}
