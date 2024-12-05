package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Camion struct {
	ID               primitive.ObjectID `bson:"_id,omitempty"`
	Patente          string             `bson:"patente"`
	Peso_maximo      float64            `bson:"peso_maximo"`
	Costo_km         int                `bson:"costo_km"`
	Fecha_creacion   time.Time          `bson:"fecha_creacion"`
	CapacidadParadas int                `bson:"capacidad_paradas"`
	Actualizacion    time.Time          `bson:"actualizacion"`
}
