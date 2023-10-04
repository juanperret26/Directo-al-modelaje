package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Producto struct {
	Id_Producto          primitive.ObjectID `bson:"_id,omitempty"`
	Nombre               string             `bson:"nombre"`
	Peso_unitario        float64            `bson:"peso"`
	Precio               float64            `bson:"precio"`
	Stock                int                `bson:"stock"`
	Stock_minimo         int                `bson:"stock_minimo"`
	Tipo                 string             `bson:"tipo"`
	Ultima_actualizacion time.Time          `bson:"ultima_actualizacion"`
	Fecha_creacion       time.Time          `bson:"fecha_creacion"`
}
