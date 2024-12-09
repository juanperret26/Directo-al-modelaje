// Crear una interface, struct y new CamionRepository
package repositories

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/juanperret26/Directo-al-modelaje/go/model"
	"github.com/juanperret26/Directo-al-modelaje/go/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type CamionRepositoryInterface interface {
	// Métodos para implementar en el service
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

// Constructor que valida que el DB no sea nulo
func NewCamionRepository(db DB) *CamionRepository {
	if db == nil {
		log.Fatal("No se puede inicializar CamionRepository con un db nulo")
	}
	return &CamionRepository{db: db}
}

// Método para obtener todos los camiones
func (repository CamionRepository) OtenerCamiones() ([]model.Camion, error) {
	collection := repository.db.GetClient().Database("DirectoAlModelaje").Collection("Camiones")
	if collection == nil {
		log.Println("[repository:CamionRepository][method:OtenerCamiones][reason:NIL_COLLECTION][message:La colección es nula]")
		return nil, fmt.Errorf("la colección de camiones es nula")
	}

	filtro := bson.M{}
	cursor, err := collection.Find(context.Background(), filtro)
	if err != nil {
		log.Printf("[repository:CamionRepository][method:OtenerCamiones][reason:ERROR][error:%v]", err)
		return nil, err
	}
	if cursor == nil {
		log.Println("[repository:CamionRepository][method:OtenerCamiones][reason:NIL_CURSOR][message:Cursor es nulo]")
		return nil, fmt.Errorf("no se pudo obtener los camiones")
	}
	defer cursor.Close(context.Background())

	var camiones []model.Camion
	for cursor.Next(context.Background()) {
		var camion model.Camion
		err := cursor.Decode(&camion)
		if err != nil {
			log.Printf("[repository:CamionRepository][method:OtenerCamiones][reason:DECODE_ERROR][error:%v]", err)
			continue
		}
		camiones = append(camiones, camion)
	}

	return camiones, nil
}

// Método para obtener un camión por su patente
func (repository CamionRepository) ObtenerCamionPorPatente(patente string) (model.Camion, error) {
	collection := repository.db.GetClient().Database("DirectoAlModelaje").Collection("Camiones")
	if collection == nil {
		log.Println("[repository:CamionRepository][method:ObtenerCamionPorPatente][reason:NIL_COLLECTION][message:La colección es nula]")
		return model.Camion{}, fmt.Errorf("la colección de camiones es nula")
	}

	filtro := bson.M{"patente": patente}
	var camion model.Camion
	err := collection.FindOne(context.Background(), filtro).Decode(&camion)
	if err != nil {
		log.Printf("[repository:CamionRepository][method:ObtenerCamionPorPatente][reason:NOT_FOUND][patente:%s][error:%v]", patente, err)
		return model.Camion{}, err
	}

	return camion, nil
}

// Método para insertar un camión
func (repository CamionRepository) InsertarCamion(camion model.Camion) (*mongo.InsertOneResult, error) {
	collection := repository.db.GetClient().Database("DirectoAlModelaje").Collection("Camiones")
	if collection == nil {
		log.Println("[repository:CamionRepository][method:InsertarCamion][reason:NIL_COLLECTION][message:La colección es nula]")
		return nil, fmt.Errorf("la colección de camiones es nula")
	}

	if camion.Patente == "" || camion.Costo_km == 0 {
		return nil, fmt.Errorf("datos inválidos: patente o costo por km son incorrectos")
	}

	camion.Fecha_creacion = time.Now()
	camion.Actualizacion = time.Now()

	resultado, err := collection.InsertOne(context.TODO(), camion)
	if err != nil {
		log.Printf("[repository:CamionRepository][method:InsertarCamion][reason:INSERT_ERROR][error:%v]", err)
		return nil, err
	}

	return resultado, nil
}

// Método para eliminar un camión por ID
func (repository CamionRepository) EliminarCamion(id string) (*mongo.DeleteResult, error) {
	objectID := utils.GetObjectIDFromStringID(id)
	if objectID == primitive.NilObjectID {
		return nil, fmt.Errorf("id inválido: no se pudo convertir a ObjectID")
	}

	collection := repository.db.GetClient().Database("DirectoAlModelaje").Collection("Camiones")
	if collection == nil {
		log.Println("[repository:CamionRepository][method:EliminarCamion][reason:NIL_COLLECTION][message:La colección es nula]")
		return nil, fmt.Errorf("la colección de camiones es nula")
	}

	filtro := bson.M{"_id": objectID}
	resultado, err := collection.DeleteOne(context.TODO(), filtro)
	if err != nil {
		log.Printf("[repository:CamionRepository][method:EliminarCamion][reason:DELETE_ERROR][id:%s][error:%v]", id, err)
		return nil, err
	}
	if resultado == nil || resultado.DeletedCount == 0 {
		err := fmt.Errorf("no se pudo encontrar el camión con id %s", id)
		log.Printf("[repository:CamionRepository][method:EliminarCamion][reason:NOT_FOUND][id:%s]", id)
		return nil, err
	}

	return resultado, nil
}

// Método para actualizar un camión
func (repository CamionRepository) ActualizarCamion(camion model.Camion) (*mongo.UpdateResult, error) {
	collection := repository.db.GetClient().Database("DirectoAlModelaje").Collection("Camiones")

	if collection == nil {
		log.Println("[repository:CamionRepository][method:ActualizarCamion][reason:NIL_COLLECTION][message:La colección es nula]")
		return nil, fmt.Errorf("la colección de camiones es nula")
	}

	if camion.ID.IsZero() {
		return nil, fmt.Errorf("el ID del camión es inválido o está vacío")
	}
	filtro := bson.M{"_id": camion.ID}

	updates := bson.M{}

	if camion.Costo_km != 0 {
		updates["costo_km"] = camion.Costo_km
	}
	if camion.Patente != "" {
		updates["patente"] = camion.Patente
	}
	if camion.Peso_maximo != 0 {
		updates["peso_maximo"] = camion.Peso_maximo
	}
	if camion.Cantidad_paradas != 0 {
		updates["capacidad_paradas"] = camion.Cantidad_paradas
	}
	
	updates["actualizacion"] = time.Now()

	if len(updates) == 0 {
		return nil, fmt.Errorf("no se encontraron campos válidos para actualizar")
	}
	entidad := bson.M{"$set": updates}

	resultado, err := collection.UpdateOne(context.TODO(), filtro, entidad)
	if err != nil {
		log.Printf("[repository:CamionRepository][method:ActualizarCamion][reason:UPDATE_ERROR][id:%s][error:%v]", camion.ID.Hex(), err)
		return nil, err
	}

	if resultado == nil || resultado.MatchedCount == 0 {
		err := fmt.Errorf("no se encontró el camión con id %s para actualizar", camion.ID.Hex())
		log.Printf("[repository:CamionRepository][method:ActualizarCamion][reason:NOT_FOUND][id:%s]", camion.ID.Hex())
		return nil, err
	}
	
	return resultado, nil
}

