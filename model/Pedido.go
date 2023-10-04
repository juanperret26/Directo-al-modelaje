package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Pedido struct {
	Id_Pedido      primitive.ObjectID `bson:"_id,omitempty"`
	Estado         string             `bson:"estado"`
	Fecha_creacion time.Time          `bson:"fecha_creacion"`
	Productos      []PedidoProducto   `bson:"productos"`
	Actualizacion  time.Time          `bson:"actualizacion"`
	Destino        string             `bson:"destino"`
}
