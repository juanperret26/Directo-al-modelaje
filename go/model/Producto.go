package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Producto struct {
	Id            primitive.ObjectID `bson:"_id,omitempty" json:"Id"`
	Nombre        string             `bson:"nombre" json:"nombre"`
	TipoProducto  string             `bson:"tipo_producto" json:"tipo_producto"`
	Peso_unitario float64            `bson:"peso" json:"peso"`
	Precio        float64            `bson:"precio" json:"precio"`
	Stock         int                `bson:"stock" json:"stock"`
	Stock_minimo  int                `bson:"stock_minimo" json:"stock_minimo"`
	Actualizacion time.Time          `bson:"actualizacion" json:"actualizacion"`
	Creacion      time.Time          `bson:"creacion" json:"creacion"`
}
 