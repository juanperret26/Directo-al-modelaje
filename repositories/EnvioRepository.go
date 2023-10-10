// Crear una interface, struct y new CamionRepository
package repositories

import (
	"context"
	"fmt"

	"github.com/juanperret/Directo-al-modelaje/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type EnvioRepositoryInterface interface {
	GetEnvios() ([]model.Envio, error)
	GetEnvio(id string) (model.Envio, error)
	InsertarEnvio(envio model.Envio) (*mongo.InsertOneResult, error)
}
type EnvioRepository struct {
	db DB
}

func NewEnvioRepository(db DB) *EnvioRepository {
	return &EnvioRepository{db: db}
}
func (repository EnvioRepository) GetEnvios() ([]model.Envio, error) {
	collection := repository.db.GetClient().Database("DirectoAlModelaje").Collection("Envios")
	filtro := bson.M{}
	cursor, err := collection.Find(context.Background(), filtro)
	defer cursor.Close(context.Background())

	var envios []model.Envio
	for cursor.Next(context.Background()) {
		var envio model.Envio
		err := cursor.Decode(&envio)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		}
		envios = append(envios, envio)
	}
	return envios, err

}
func (repository EnvioRepository) GetEnvio(id string) (model.Envio, error) {
	collection := repository.db.GetClient().Database("DirectoAlModelaje").Collection("Envios")
	filtro := bson.M{"_id": id}
	var envio model.Envio
	err := collection.FindOne(context.Background(), filtro).Decode(&envio)
	return envio, err
}
func (repository EnvioRepository) InsertarEnvio(envio model.Envio) (*mongo.InsertOneResult, error) {
	collection := repository.db.GetClient().Database("DirectoAlModelaje").Collection("Envios")
	resultado, err := collection.InsertOne(context.Background(), envio)
	return resultado, err
}
