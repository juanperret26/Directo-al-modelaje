package model

type PedidoProducto struct {
	CodigoProducto  string  `bson:"codigo_producto"`
	Nombre          string  `bson:"nombre"`
	Cantidad        int     `bson:"cantidad"`
	Precio_unitario float64 `bson:"precio_unitario"`
	Stock           int     `bson:"stock"`
}
