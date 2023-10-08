package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Camion struct {
	ID             primitive.ObjectID `bson:"_id,omitempty"`
	Patente        string             `bson:"patente"`
	Peso_maximo    int                `bson:"peso_maximo"`
	Costo_km       int                `bson:"costo_km"`
	Fecha_creacion time.Time          `bson:"fecha_creacion"`
	Actualizacion  time.Time          `bson:"actualizacion"`
}
