package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Envio struct {
	Id            primitive.ObjectID `bson:"_id,omitempty"`
	CodPedido     string             `bson:"codigo_pedido"`
	Estado        string             `bson:"estado"`
	Paradas       []Paradas          `bson:"paradas"`
	Destino       string             `bson:"destino"`
	Creacion      time.Time          `bson:"fecha_creacion"`
	Pedido        []string           `bson:"productos"`
	Actualizacion time.Time          `bson:"actualizacion"`
	Costo         int                `bson:"costo_total"`
}
