// Crear una interface, struct y new CamionRepository
package repositories

import (
	"context"
	"fmt"
	"time"

	"github.com/juanperret/Directo-al-modelaje/model"
	"github.com/juanperret/Directo-al-modelaje/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type EnvioRepositoryInterface interface {
	ObtenerEnvios() ([]model.Envio, error)
	ObtenerEnvioPorId(id string) (model.Envio, error)
	ObtenerPedidosFiltrados(codigoEnvio string, estado string, fechaInicio time.Time, fechaFinal time.Time) ([]model.Pedido, error)
	ObtenerEnviosPorParametros(patente string, estado string, ultimaParada string, fechaCreacionDesde time.Time, fechaCreacionHasta time.Time) ([]model.Envio, error)
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
func (repository *EnvioRepository) ObtenerPedidosFiltrados(codigoEnvio string, estado string, fechaInicio time.Time, fechaFinal time.Time) ([]model.Pedido, error) {
	collection := repository.db.GetClient().Database("DirectoAlModelaje").Collection("Envios")

	filtro := bson.M{"codigo_envio": codigoEnvio, "estado": estado}

	if !fechaInicio.IsZero() && !fechaFinal.IsZero() {
		filtro["fecha_creacion"] = bson.M{"$gte": fechaInicio, "$lte": fechaFinal}
	}

	cursor, err := collection.Find(context.Background(), filtro)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var pedidos []model.Pedido
	for cursor.Next(context.Background()) {
		var pedido model.Pedido
		if err := cursor.Decode(&pedido); err != nil {
			fmt.Printf("Error decoding pedido: %v\n", err)
			continue
		}
		pedidos = append(pedidos, pedido)
	}

	return pedidos, nil
}

func (repository EnvioRepository) ObtenerEnviosPorParametros(patente string, estado string, ultimaParada string, fechaCreacionDesde time.Time, fechaCreacionHasta time.Time) ([]model.Envio, error) {
	collection := repository.db.GetClient().Database("DirectoAlModelaje").Collection("Envios")
	filtro := bson.M{}

	//Solo filtra por patente si le pasamos un valor distinto de ""
	if patente != "" {
		filtro["patente_camion"] = patente
	}

	//Solo filtra por estado si le pasamos un estado positivo
	if estado != "" {
		filtro["estado"] = estado
	}

	//Tomo la fecha de creacion en 0001-01-01 como la ausencia de filtro
	if !fechaCreacionDesde.IsZero() || !fechaCreacionHasta.IsZero() {
		filtroFecha := bson.M{}
		if !fechaCreacionDesde.IsZero() {
			filtroFecha["$gte"] = fechaCreacionDesde
		}
		if !fechaCreacionHasta.IsZero() {
			filtroFecha["$lte"] = fechaCreacionHasta
		}
		filtro["fecha_creacion"] = filtroFecha
	}

	//TODO: hay que probar que este filtro ande bien

	if ultimaParada != "" {
		filtro["paradas"] = bson.M{}
		filtro["paradas.$slice"] = -1
	}

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
