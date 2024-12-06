package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Envio struct {
	Id            primitive.ObjectID `bson:"_id,omitempty" json:"ID"`
	PatenteCamion string             `bson:"patente_camion" json:"PatenteCamion"`
	Estado        string             `bson:"estado" json:"estado"`
	Paradas       []Parada           `bson:"paradas" json:"paradas"`
	Destino       Parada             `bson:"destino" json:"destino"`
	Creacion      time.Time          `bson:"fecha_creacion" json:"creacion"`
	Pedido        []string           `bson:"pedidos" json:"pedido"`
	Actualizacion time.Time          `bson:"actualizacion" json:"actualizacion"`
	Costo         int                `bson:"costo_total" json:"costo"`
}
