package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type PedidoProducto struct {
	Id_pedidoProducto primitive.ObjectID `bson:"id_pedidoProducto"`
	Nombre            string             `bson:"nombre"`
	Cantidad          int                `bson:"cantidad"`
	Precio_unitario   float64            `bson:"precio_unitario"`
	Stock             int                `bson:"stock"`
	Tipo              string             `bson:"tipo"`
}
