package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Envio struct {
	Id_envio       primitive.ObjectID `bson:"_id,omitempty"`
	Estado         string             `bson:"estado"`
	Paradas        []Paradas          `bson:"paradas"`
	Fecha_creacion time.Time          `bson:"fecha_creacion"`
	Pedido         []string           `bson:"productos"`
	Actualizacion  time.Time          `bson:"actualizacion"`
	Costo_total    int                `bson:"costo_total"`
}
