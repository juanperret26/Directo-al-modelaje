package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Pedido struct {
	id_Pedido primitive.ObjectID `bson:"_id,omitempty"`
	estado string `bson:"estado"`
	productos []Producto `bson:"productos"`
	actualizacion time.Time `bson:"actualizacion"`
	destino Ciudad `bson:"ciudad"`
}