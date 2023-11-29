package dto

import (
<<<<<<< HEAD:go/dto/Parada.go
	"github.com/juanperret26/Directo-al-modelaje/go/model"
=======
	"github.com/juanperret/Directo-al-modelaje/model"
>>>>>>> parent of ad42c9c (prueba cambio de directorio):dto/Parada.go
)

type Parada struct {
	Ciudad     string `json:"ciudad"`
	Kilometros int    `json:"kilometros"`
}

func (parada Parada) GetModel() model.Parada {
	return model.Parada{
		Nombre_ciudad:         parada.Ciudad,
		Kilometros_recorridos: parada.Kilometros,
	}
}

// Metodo para crear un dto a partir del modelo
func NewParada(parada *model.Parada) *Parada {
	return &Parada{
		Ciudad:     parada.Nombre_ciudad,
		Kilometros: parada.Kilometros_recorridos,
	}
}
