package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type PedidoProducto struct {
	Id_pedidoProducto primitive.ObjectID `bson:"id_pedidoProducto"`
	CodigoProducto    string             `bson:"codigo_producto"`
	Nombre            string             `bson:"nombre"`
	Cantidad          float64            `bson:"cantidad"`
	Precio_unitario   float64            `bson:"precio_unitario"`
	Stock             int                `bson:"stock"`
	Tipo              string             `bson:"tipo"`
}
