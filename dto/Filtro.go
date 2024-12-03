package dto

import (
	"time"
)

type Filtro struct {
	CodigoEnvio                   string
	PatenteCamion                 string
	EstadoEnvio                   string
	EstadoPedido                  string
	UltimaParada                  string
	FechaCreacionDesde            time.Time
	FechaCreacionHasta            time.Time
	FechaUltimaActualizacionDesde time.Time
	FechaUltimaActualizacionHasta time.Time
}
