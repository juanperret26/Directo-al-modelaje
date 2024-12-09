package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Camion struct {
	ID               primitive.ObjectID `bson:"_id,omitempty" json:"ID"`
	Patente          string             `bson:"patente" json:"patente"`
	Peso_maximo      float64            `bson:"peso_maximo" json:"peso_maximo"`
	Costo_km         int                `bson:"costo_km" json:"Costo_km"`
	Fecha_creacion   time.Time          `bson:"fecha_creacion" json:"fecha_creacion"`
	Cantidad_paradas int                `bson:"cantidad" json:"cantidad_paradas"`
	Actualizacion    time.Time          `bson:"actualizacion" json:"actualizacion"`
}
