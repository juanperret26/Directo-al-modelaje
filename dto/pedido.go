// Crear struct y un new pedido con los metodos que sean necesario
package dto

import (
	"time"

	"github.com/juanperret/Directo-al-modelaje/model"
)

type Pedido struct {
	id_Pedido      int
	estado         string
	fecha_creacion time.Time
	productos      []Producto
	actualizacion  time.Time
	destino        Ciudad
}

func NewPedido(pedido model.Pedido) *Pedido {
	return &Pedido{
		id_Pedido:      model.GetIntIDFromObjectID(pedido.id_Pedido),
		estado:         pedido.estado,
		fecha_creacion: pedido.fecha_creacion,
		productos:      pedido.productos,
		actualizacion:  pedido.actualizacion,
		destino:        pedido.destino,
	}
}
