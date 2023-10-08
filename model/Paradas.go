package model

type Paradas struct {
	Id             int    `bson:"id_parada"`
	Nombre_ciudad         string `bson:"nombre_ciudad"`
	Kilometros_recorridos int    `bson:"kilometros_recorridos"`
}
