package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Camion struct {
	ID               primitive.ObjectID `bson:"_id,omitempty" json:"ID"`
	Patente          string             `bson:"patente" json:"Patente"`
	Peso_maximo      float64            `bson:"peso_maximo" json:"Peso_maximo"`
	Costo_km         int                `bson:"costo_km" json:"Costo_km"`
	Fecha_creacion   time.Time          `bson:"fecha_creacion" json:"Fecha_creacion"`
	CapacidadParadas int                `bson:"capacidad_paradas" json:"CapacidadParadas"`
	Actualizacion    time.Time          `bson:"actualizacion" json:"Actualizacion"`
}
