// Crear una interface, struct y new CamionRepository
package repositories

import (
	"context"
	"fmt"
	"time"

	"github.com/juanperret/Directo-al-modelaje/go/model"
	"github.com/juanperret/Directo-al-modelaje/go/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type CamionRepositoryInterface interface {
	//Metodos para implementar en el service
	OtenerCamiones() ([]model.Camion, error)
	ObtenerCamionPorPatente(patente string) (model.Camion, error)
	InsertarCamion(camion model.Camion) (*mongo.InsertOneResult, error)
	EliminarCamion(id string) (*mongo.DeleteResult, error)
	ActualizarCamion(camion model.Camion) (*mongo.UpdateResult, error)
}
type CamionRepository struct {
	// DB represents a database connection.
	db DB
}

func NewCamionRepository(db DB) *CamionRepository {
	return &CamionRepository{db: db}
}

func (repository CamionRepository) OtenerCamiones() ([]model.Camion, error) {
	collection := repository.db.GetClient().Database("DirectoAlModelaje").Collection("Camiones")
	filtro := bson.M{}
	cursor, err := collection.Find(context.Background(), filtro)
	defer cursor.Close(context.Background())

	var camiones []model.Camion
	for cursor.Next(context.Background()) {
		var camion model.Camion
		err := cursor.Decode(&camion)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		}
		camiones = append(camiones, camion)
	}
	return camiones, err
}

func (repository CamionRepository) ObtenerCamionPorPatente(patente string) (model.Camion, error) {
	collection := repository.db.GetClient().Database("DirectoAlModelaje").Collection("Camiones")
	filtro := bson.M{"patente": patente}
	var camion model.Camion
	err := collection.FindOne(context.Background(), filtro).Decode(&camion)
	return camion, err
}

// func (repository *CamionRepository) ObtenerCamionPorPatente(patente string) (model.Camion, error) {
// 	if repository == nil || repository.db == nil || repository.db.GetClient() == nil {
// 		return model.Camion{}, errors.New("El repositorio es nulo")
// 	}

// 	collection := repository.db.GetClient().Database("DirectoAlModelaje").Collection("Camiones")
// 	if collection == nil {
// 		return model.Camion{}, errors.New("La colección de la base de datos es nula")
// 	}

// 	filtro := bson.M{"patente": patente}
// 	var camion model.Camion
// 	err := collection.FindOne(context.Background(), filtro).Decode(&camion)
// 	if err != nil {
// 		return model.Camion{}, err
// 	}

// 	return camion, nil
// }

func (repository CamionRepository) InsertarCamion(camion model.Camion) (*mongo.InsertOneResult, error) {
	collection := repository.db.GetClient().Database("DirectoAlModelaje").Collection("Camiones")
	if camion.Patente != "" && camion.Costo_km != 0 {
		camion.Fecha_creacion = time.Now()
		camion.Actualizacion = time.Now()

	}
	resultado, err := collection.InsertOne(context.TODO(), camion)
	return resultado, err
}

func (repository CamionRepository) EliminarCamion(id string) (*mongo.DeleteResult, error) {
	objectID := utils.GetObjectIDFromStringID(id)

	collection := repository.db.GetClient().Database("DirectoAlModelaje").Collection("Camiones")
	filtro := bson.M{"_id": objectID}
	resultado, err := collection.DeleteOne(context.TODO(), filtro)
	if resultado == nil {
		err = fmt.Errorf("No se pudo enxontrar el camion con id %s", id)
	}
	return resultado, err
}

func (repository CamionRepository) ActualizarCamion(camion model.Camion) (*mongo.UpdateResult, error) {
	collection := repository.db.GetClient().Database("DirectoAlModelaje").Collection("Camiones")
	filtro := bson.M{"_id": camion.ID}
	entidad := bson.M{"$set": bson.M{"costo_km": camion.Costo_km, "actualizacion": time.Now()}}
	resultado, err := collection.UpdateOne(context.TODO(), filtro, entidad)
	return resultado, err
}
