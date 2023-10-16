package dto

import (
	"time"

	"github.com/juanperret/Directo-al-modelaje/model"
	"github.com/juanperret/Directo-al-modelaje/utils"
)

type Camion struct {
	ID             string
	Patente        string
	Peso_maximo    int
	Costo_km       int
	Fecha_creacion time.Time
	Actualizacion  time.Time
}

func NewCamion(camion model.Camion) *Camion {
	return &Camion{
		ID:             utils.GetStringIDFromObjectID(camion.ID),
		Patente:        camion.Patente,
		Peso_maximo:    camion.Peso_maximo,
		Costo_km:       camion.Costo_km,
		Fecha_creacion: time.Now(),
		Actualizacion:  time.Now(),
	}
}
func (camion Camion) GetModel() model.Camion {
	return model.Camion{
		ID:             utils.GetObjectIDFromStringID(camion.ID),
		Patente:        camion.Patente,
		Peso_maximo:    camion.Peso_maximo,
		Costo_km:       camion.Costo_km,
		Fecha_creacion: camion.Fecha_creacion,
		Actualizacion:  camion.Actualizacion,
	}
}
