package dto

import "time"

type Camion struct {
	patente        int
	peso_maximo    int
	costo_km       int
	fecha_creacion time.Time
	actualizacion  time.Time
}
