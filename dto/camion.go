package dto

import (
	"time"

	"github.com/juanperret/Directo-al-modelaje/model"
	"github.com/juanperret/Directo-al-modelaje/utils"
)

type Camion struct {
	Patente        string
	Peso_maximo    int
	Costo_km       int
	Fecha_creacion time.Time
	Actualizacion  time.Time
}

func NewCamion(camion model.Camion) *Camion {
	return &Camion{
		Patente:        utils.GetStringIDFromObjectID(camion.Patente),
		Peso_maximo:    camion.Peso_maximo,
		Costo_km:       camion.Costo_km,
		Fecha_creacion: time.Now(),
		Actualizacion:  camion.Actualizacion,
	}
}
func (camion Camion) GetModel() model.Camion {
	return model.Camion{
		Patente:        utils.GetObjectIDFromStringID(camion.Patente),
		Peso_maximo:    camion.Peso_maximo,
		Costo_km:       camion.Costo_km,
		Fecha_creacion: camion.Fecha_creacion,
		Actualizacion:  camion.Actualizacion,
	}
}
