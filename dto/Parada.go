package dto

import (
	"github.com/juanperret/Directo-al-modelaje/model"
)

type Parada struct {
	Ciudad     string `json:"ciudad"`
	Kilometros int    `json:"kilometros"`
}

func (parada Parada) GetModel() model.Paradas {
	return model.Paradas{
		Nombre_ciudad:         parada.Ciudad,
		Kilometros_recorridos: parada.Kilometros,
	}
}

// Metodo para crear un dto a partir del modelo
func NewParada(parada *model.Paradas) *Parada {
	return &Parada{
		Ciudad:     parada.Nombre_ciudad,
		Kilometros: parada.Kilometros_recorridos,
	}
}
