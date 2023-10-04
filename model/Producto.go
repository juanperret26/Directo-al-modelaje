package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Producto struct {
	id_Producto primitive.ObjectID `bson:"_id,omitempty"`
	nombre string `bson:"nombre"`
	peso float64 `bson:"peso"`
	precio float64 `bson:"precio"`
	stock int `bson:"stock"`
	stock_minimo int `bson:"stock_minimo"`
	tipo string `bson:"tipo"`
	ultima_actualizacion time.Time`bson:"ultima_actualizacion"`
	fecha_creacion time.Time `bson:"fecha_creacion"`
}