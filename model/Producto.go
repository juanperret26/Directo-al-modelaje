package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Producto struct {
	Id             primitive.ObjectID `bson:"_id,omitempty"`
	CodigoProducto string             `bson:"codigo"`
	Nombre         string             `bson:"nombre"`
	Peso_unitario  float64            `bson:"peso"`
	Precio         float64            `bson:"precio"`
	Stock          float64            `bson:"stock"`
	Stock_minimo   int                `bson:"stock_minimo"`
	Tipo           string             `bson:"tipo"`
	Actualizacion  time.Time          `bson:"actualizacion"`
	Creacion       time.Time          `bson:"creacion"`
}
