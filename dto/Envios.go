package dto

import (
	"time"

	"github.com/juanperret/Directo-al-modelaje/model"
	"github.com/juanperret/Directo-al-modelaje/utils"
)

type Envio struct {
	Id_envio       string
	Estado         string
	Paradas        []Paradas
	Fecha_creacion time.Time
	Pedido         []string
	Actualizacion  time.Time
	Costo_total    int
}

func NewEnvio(envio model.Envio) model.Envio {
	return model.Envio{
		Id_envio:       utils.GetObjectIDFromStringID(envio.Id_envio.Hex()),
		Estado:         envio.Estado,
		Paradas:        envio.Paradas,
		Fecha_creacion: time.Now(),
		Pedido:         envio.Pedido,
		Actualizacion:  envio.Actualizacion,
		Costo_total:    envio.Costo_total,
	}
}
func (envio Envio) GetModel() model.Envio {
	return model.Envio{
		Id_envio:       utils.GetObjectIDFromStringID(envio.Id_envio),
		Estado:         envio.Estado,
		Paradas:        []model.Paradas{},
		Fecha_creacion: envio.Fecha_creacion,
		Pedido:         envio.Pedido,
		Actualizacion:  envio.Actualizacion,
		Costo_total:    envio.Costo_total,
	}
}
