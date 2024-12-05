package model

type Parada struct {
	Nombre_ciudad         string `bson:"nombre_ciudad"`
	Kilometros_recorridos int    `bson:"kilometros_recorridos"`
}
