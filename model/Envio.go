package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Envio struct {
	Id            primitive.ObjectID `bson:"_id,omitempty"`
	PatenteCamion string             `bson:"patente_camion"`
	Estado        string             `bson:"estado"`
	Paradas       []Parada           `bson:"paradas"`
	Destino       Parada             `bson:"destino"`
	Creacion      time.Time          `bson:"fecha_creacion"`
	Pedido        []string           `bson:"pedidos"`
	Actualizacion time.Time          `bson:"actualizacion"`
	Costo         int                `bson:"costo_total"`
}
