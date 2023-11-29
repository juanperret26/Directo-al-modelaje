package dto

import (
	"time"

	"github.com/juanperret/Directo-al-modelaje/go/model"
	"github.com/juanperret/Directo-al-modelaje/go/utils"
)

type Camion struct {
	ID             string    `json:"id"`
	Patente        string    `json:"patente"`
	Peso_maximo    float64   `json:"peso_maximo"`
	Costo_km       int       `json:"costo_km"`
	Fecha_creacion time.Time `json:"fecha_creacion"`
	Actualizacion  time.Time `json:"actualizacion"`
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
