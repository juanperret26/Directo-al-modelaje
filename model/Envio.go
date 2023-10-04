package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Envio struct {
	id_Envio       primitive.ObjectID `bson:"_id,omitempty"`
	estado         string             `bson:"estado"`
	paradas        []Ciudad           `bson:"paradas"`
	fecha_creacion time.Time          `bson:"fecha_creacion"`
	productos      []Producto         `bson:"productos"`
	actualizacion  time.Time          `bson:"actualizacion"`
	costo_total    int                `bson:"costo_total"`
}
