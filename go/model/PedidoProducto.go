package model

type PedidoProducto struct {
	CodigoProducto  string  `bson:"codigo_producto" json:"codigoproducto"`
	Nombre          string  `bson:"nombre" json:"nombre"`
	Cantidad        int     `bson:"cantidad" json:"cantidad"`
	Precio_unitario float64 `bson:"precio_unitario" json:"precio_unitario"`
	Stock           int     `bson:"stock" json:"stock"`
}
