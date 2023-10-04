package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Camion struct {
	Patente        primitive.ObjectID `bson:"_patente,omitempty"`
	Peso_maximo    int                `bson:"peso_maximo"`
	Costo_km       int                `bson:"costo_km"`
	Fecha_creacion time.Time          `bson:"fecha_creacion"`
	Actualizacion  time.Time          `bson:"actualizacion"`
}
