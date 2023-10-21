// Crear una interface, struct y new CamionRepository
package repositories

import (
	"context"
	"fmt"

	"github.com/juanperret/Directo-al-modelaje/model"
	"github.com/juanperret/Directo-al-modelaje/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type EnvioRepositoryInterface interface {
	ObtenerEnvios() ([]model.Envio, error)
	ObtenerEnvioPorId(id string) (model.Envio, error)
	InsertarEnvio(envio model.Envio) (*mongo.InsertOneResult, error)
	EliminarEnvio(id primitive.ObjectID) (*mongo.DeleteResult, error)
	ActualizarEnvio(envio model.Envio) (*mongo.UpdateResult, error)
}
type EnvioRepository struct {
	db DB
}

func NewEnvioRepository(db DB) *EnvioRepository {
	return &EnvioRepository{db: db}
}
func (repository EnvioRepository) ObtenerEnvios() ([]model.Envio, error) {
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
func (repository EnvioRepository) ObtenerEnvioPorId(id string) (model.Envio, error) {
	collection := repository.db.GetClient().Database("DirectoAlModelaje").Collection("Envios")
	objectID := utils.GetObjectIDFromStringID(id)
	filtro := bson.M{"_id": objectID}
	var envio model.Envio
	err := collection.FindOne(context.Background(), filtro).Decode(&envio)
	return envio, err
}
func (repository EnvioRepository) InsertarEnvio(envio model.Envio) (*mongo.InsertOneResult, error) {
	collection := repository.db.GetClient().Database("DirectoAlModelaje").Collection("Envios")
	resultado, err := collection.InsertOne(context.Background(), envio)
	return resultado, err
}

func (repository EnvioRepository) EliminarEnvio(id primitive.ObjectID) (*mongo.DeleteResult, error) {
	collection := repository.db.GetClient().Database("DirectoAlModelaje").Collection("Envios")
	filtro := bson.M{"_id": id}
	resultado, err := collection.DeleteOne(context.TODO(), filtro)
	return resultado, err
}
func (repository EnvioRepository) ActualizarEnvio(envio model.Envio) (*mongo.UpdateResult, error) {
	collection := repository.db.GetClient().Database("DirectoAlModelaje").Collection("Envios")
	filtro := bson.M{"_id": envio.Id}
	entidad := bson.M{"$set": bson.M{"estado": envio.Estado}}
	resultado, err := collection.UpdateOne(context.TODO(), filtro, entidad)
	return resultado, err
}
