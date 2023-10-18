package dto

import (
	"time"

	"github.com/juanperret/Directo-al-modelaje/model"
	"github.com/juanperret/Directo-al-modelaje/utils"
)

type Envio struct {
	Id            string
	Estado        string
	Paradas       []Paradas
	Destino       string
	Creacion      time.Time
	Pedido        []string
	Actualizacion time.Time
	Costo         int
}

func NewEnvio(envio model.Envio) *Envio {
	return &Envio{
		Id:            utils.GetStringIDFromObjectID(envio.Id),
		Estado:        envio.Estado,
		Paradas:       []Paradas{},
		Creacion:      time.Now(),
		Pedido:        envio.Pedido,
		Actualizacion: time.Now(),
		Costo:         envio.Costo,
	}
}
func (envio Envio) GetModel() model.Envio {
	return model.Envio{
		Id:            utils.GetObjectIDFromStringID(envio.Id),
		Estado:        envio.Estado,
		Paradas:       []model.Paradas{},
		Creacion:      envio.Creacion,
		Pedido:        envio.Pedido,
		Actualizacion: envio.Actualizacion,
		Costo:         envio.Costo,
	}
}
