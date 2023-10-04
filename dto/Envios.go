// Crear struct y un new envio con los metodos que sean necesario
//
package dto

import (
	"time"

	"github.com/juanperret/Directo-al-modelaje/model"
)

type Envio struct {
	id_Envio int
	patente int
	paradas[] Ciudad
	fecha_creacion time.Time
	pedidos[] Pedido
	actualizacion time.Time
	costo_total int
}
func NewEnvio(envio model.Envio) *Envio {
	return &Envio{
		id_Envio: model.GetIntIDFromObjectID(envio.id_Envio),
		patente: envio.patente,
		paradas: envio.paradas,
		fecha_creacion: envio.fecha_creacion,
		pedidos: envio.pedidos,
		actualizacion: envio.actualizacion,
		costo_total: envio.costo_total,
	}
}
