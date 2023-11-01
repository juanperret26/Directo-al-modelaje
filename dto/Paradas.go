package dto

import (
	"github.com/juanperret/Directo-al-modelaje/model"
)

type Paradas struct {
	Ciudad     string
	Kilometros int
}

func (parada Paradas) GetModel() model.Paradas {
	return model.Paradas{
		Nombre_ciudad:         parada.Ciudad,
		Kilometros_recorridos: parada.Kilometros,
	}
}

// Metodo para crear un dto a partir del modelo
func NewParada(parada *model.Paradas) *Paradas {
	return &Paradas{
		Ciudad:     parada.Nombre_ciudad,
		Kilometros: parada.Kilometros_recorridos,
	}
}
