package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Pedido struct {
	Id              primitive.ObjectID `bson:"_id,omitempty" json:"ID"`
	Estado          string             `bson:"estado" json:"estado"`
	Fecha_creacion  time.Time          `bson:"fecha_creacion" json:"fecha_creacion"`
	PedidoProductos []PedidoProducto   `bson:"productos" json:"pedidoproductos"`
	Actualizacion   time.Time          `bson:"actualizacion" json:"actualizacion"`
	Destino         string             `bson:"destino" json:"destino"`
}
 