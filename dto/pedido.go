// Crear struct y un new pedido con los metodos que sean necesario
package dto

import (
	"time"

	"github.com/juanperret/Directo-al-modelaje/model"
	"github.com/juanperret/Directo-al-modelaje/utils"
)

type Pedido struct {
	Id_Pedido      string
	Estado         string
	Fecha_creacion time.Time
	Productos      []PedidoProducto
	Actualizacion  time.Time
	Destino        string
}

func NewPedido(pedido model.Pedido) *Pedido {
	return &Pedido{
		Id_Pedido:      utils.GetStringIDFromObjectID(pedido.Id_Pedido),
		Estado:         pedido.Estado,
		Fecha_creacion: time.Now(),
		Productos:      []PedidoProducto{},
		Actualizacion:  pedido.Actualizacion,
		Destino:        pedido.Destino,
	}
}
func (pedido Pedido) GetModel() model.Pedido {
	return model.Pedido{
		Id_Pedido:      utils.GetObjectIDFromStringID(pedido.Id_Pedido),
		Estado:         pedido.Estado,
		Fecha_creacion: pedido.Fecha_creacion,
		Productos:      []model.PedidoProducto{},
		Actualizacion:  pedido.Actualizacion,
		Destino:        pedido.Destino,
	}
}
