package model

type PedidoProducto struct {
	id_pedidoProducto int     `bson:"id_pedidoProducto"`
	nombre            string  `bson:"nombre"`
	cantidad          int     `bson:"cantidad"`
	precio_unitario   float64 `bson:"precio_unitario"`
	stock             int     `bson:"stock"`
	tipo              string  `bson:"tipo"`
}
