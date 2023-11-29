package dto

import (
	"time"

	"github.com/juanperret/Directo-al-modelaje/go/model"
	"github.com/juanperret/Directo-al-modelaje/go/utils"
)

type Envio struct {
	Id            string    `json:"id"`
	PatenteCamion string    `json:"patente_camion"`
	Estado        string    `json:"estado"`
	Paradas       []Parada  `json:"paradas"`
	Destino       Parada    `json:"destino"`
	Creacion      time.Time `json:"fecha_creacion"`
	Pedido        []string  `json:"pedidos"`
	Actualizacion time.Time `json:"actualizacion"`
	Costo         int       `json:"costo_total"`
}

func NewEnvio(envio model.Envio) *Envio {
	return &Envio{
		Id:            utils.GetStringIDFromObjectID(envio.Id),
		PatenteCamion: envio.PatenteCamion,
		Estado:        envio.Estado,
		Paradas:       NewParadas(envio.Paradas),
		Destino:       NewDestino(envio.Destino),
		Creacion:      time.Now(),
		Pedido:        envio.Pedido,
		Actualizacion: time.Now(),
		Costo:         envio.Costo,
	}
}
func (envio Envio) GetModel() model.Envio {
	return model.Envio{
		Id:            utils.GetObjectIDFromStringID(envio.Id),
		PatenteCamion: envio.PatenteCamion,
		Estado:        envio.Estado,
		Paradas:       envio.getParadas(),
		Destino:       envio.GetDestino(),
		Creacion:      envio.Creacion,
		Pedido:        envio.Pedido,
		Actualizacion: envio.Actualizacion,
		Costo:         envio.Costo,
	}
}

// Metodo para convertir una lista de Paradas del dto a una lista de Paradas del modelo
func (envio Envio) getParadas() []model.Parada {
	var paradasEnvio []model.Parada
	for _, parada := range envio.Paradas {
		paradasEnvio = append(paradasEnvio, parada.GetModel())
	}
	return paradasEnvio
}

// Metodo para convertir una lista de Paradas del modelo a una lista de Paradas del dto
func NewParadas(paradas []model.Parada) []Parada {
	var paradasEnvio []Parada
	for _, parada := range paradas {
		paradasEnvio = append(paradasEnvio, *NewParada(&parada))
	}
	return paradasEnvio
}

func NewDestino(parada model.Parada) Parada {
	return Parada{
		Ciudad:     parada.Nombre_ciudad,
		Kilometros: parada.Kilometros_recorridos,
	}
}

func (envio Envio) GetDestino() model.Parada {
	var destino model.Parada
	{
		destino.Nombre_ciudad = envio.Destino.Ciudad
		destino.Kilometros_recorridos = envio.Destino.Kilometros
	}
	return destino
}
