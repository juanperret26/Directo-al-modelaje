package repositories

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/juanperret26/Directo-al-modelaje/go/model"
	"github.com/juanperret26/Directo-al-modelaje/go/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// EnvioRepositoryInterface define los métodos a implementar.
type EnvioRepositoryInterface interface {
	ObtenerEnvios() ([]model.Envio, error)
	ObtenerEnvioPorId(id string) (model.Envio, error)
	ObtenerPedidosFiltro(codigoEnvio string, estado string, fechaInicio time.Time, fechaFinal time.Time) ([]model.Pedido, error)
	ObtenerEnviosFiltro(patente string, estado string, ultimaParada string, fechaCreacionDesde time.Time, fechaCreacionHasta time.Time) ([]model.Envio, error)
	InsertarEnvio(envio model.Envio) (*mongo.InsertOneResult, error)
	EliminarEnvio(id primitive.ObjectID) (*mongo.DeleteResult, error)
	ActualizarEnvio(envio model.Envio) error
	ObtenerCantidadEnviosPorEstado(estado string) (int, error)
	IniciarViaje(envio model.Envio) error
}

// EnvioRepository es la implementación de EnvioRepositoryInterface.
type EnvioRepository struct {
	db DB
}

// NewEnvioRepository crea una nueva instancia de EnvioRepository.
func NewEnvioRepository(db DB) *EnvioRepository {
	return &EnvioRepository{db: db}
}

// IniciarViaje actualiza el estado de un envío para iniciar el viaje.
func (repository *EnvioRepository) IniciarViaje(envio model.Envio) error {
	if envio.Id.IsZero() {
		return errors.New("el ID del envío no puede estar vacío")
	}
	collection := repository.db.GetClient().Database("DirectoAlModelaje").Collection("Envios")
	filtro := bson.M{"_id": envio.Id}
	actualizacion := bson.M{"$set": envio}

	_, err := collection.UpdateOne(context.Background(), filtro, actualizacion)
	return err
}

// ObtenerEnvios devuelve todos los envíos.
func (repository *EnvioRepository) ObtenerEnvios() ([]model.Envio, error) {
	collection := repository.db.GetClient().Database("DirectoAlModelaje").Collection("Envios")
	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		log.Printf("Error al obtener envíos: %v", err)
		return nil, err
	}
	defer cursor.Close(context.Background())

	var envios []model.Envio
	for cursor.Next(context.Background()) {
		var envio model.Envio
		if err := cursor.Decode(&envio); err != nil {
			log.Printf("Error al decodificar envío: %v", err)
			continue
		}
		envios = append(envios, envio)
	}
	return envios, nil
}

// ObtenerEnvioPorId devuelve un envío dado su ID.
func (repository *EnvioRepository) ObtenerEnvioPorId(id string) (model.Envio, error) {
	if id == "" {
		return model.Envio{}, errors.New("el ID no puede estar vacío")
	}
	objectID := utils.GetObjectIDFromStringID(id)

	collection := repository.db.GetClient().Database("DirectoAlModelaje").Collection("Envios")
	var envio model.Envio
	if err := collection.FindOne(context.Background(), bson.M{"_id": objectID}).Decode(&envio); err != nil {
		return model.Envio{}, fmt.Errorf("error al buscar envío: %w", err)
	}
	return envio, nil
}

// InsertarEnvio agrega un nuevo envío a la base de datos.
func (repository *EnvioRepository) InsertarEnvio(envio model.Envio) (*mongo.InsertOneResult, error) {
	if envio.Id.IsZero() {
		envio.Id = primitive.NewObjectID()
	}
	envio.Creacion = time.Now()
	envio.Actualizacion = time.Now()

	collection := repository.db.GetClient().Database("DirectoAlModelaje").Collection("Envios")
	return collection.InsertOne(context.Background(), envio)
}

// EliminarEnvio elimina un envío dado su ID.
func (repository *EnvioRepository) EliminarEnvio(id primitive.ObjectID) (*mongo.DeleteResult, error) {
	if id.IsZero() {
		return nil, errors.New("el ID del envío no puede estar vacío")
	}

	collection := repository.db.GetClient().Database("DirectoAlModelaje").Collection("Envios")
	return collection.DeleteOne(context.Background(), bson.M{"_id": id})
}

// ActualizarEnvio actualiza un envío en la base de datos.
func (repository *EnvioRepository) ActualizarEnvio(envio model.Envio) error {
	if envio.Id.IsZero() {
		return errors.New("el ID del envío no puede estar vacío")
	}
	envio.Actualizacion = time.Now()

	collection := repository.db.GetClient().Database("DirectoAlModelaje").Collection("Envios")
	_, err := collection.UpdateOne(context.Background(), bson.M{"_id": envio.Id}, bson.M{"$set": envio})
	return err
}

// ObtenerCantidadEnviosPorEstado devuelve la cantidad de envíos en un estado específico.
func (repository *EnvioRepository) ObtenerCantidadEnviosPorEstado(estado string) (int, error) {
	collection := repository.db.GetClient().Database("DirectoAlModelaje").Collection("Envios")
	count, err := collection.CountDocuments(context.Background(), bson.M{"estado": estado})
	if err != nil {
		return 0, err
	}
	return int(count), nil
}
func (repository *EnvioRepository) ObtenerEnviosFiltro(patente string, estado string, ultimaParada string, fechaCreacionDesde time.Time, fechaCreacionHasta time.Time) ([]model.Envio, error) {
	collection := repository.db.GetClient().Database("DirectoAlModelaje").Collection("Envios")
	filtro := bson.M{}
	if patente != "" {
		filtro["patente"] = patente
	}
	if estado != "" {
		filtro["estado"] = estado
	}
	if ultimaParada != "" {
		filtro["ultima_parada"] = ultimaParada
	}
	if !fechaCreacionDesde.IsZero() {
		filtro["creacion"] = bson.M{"$gte": fechaCreacionDesde}
	}
	if !fechaCreacionHasta.IsZero() {
		filtro["creacion"] = bson.M{"$lte": fechaCreacionHasta}
	}
	cursor, err := collection.Find(context.Background(), filtro)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())
	var envios []model.Envio
	for cursor.Next(context.Background()) {
		var envio model.Envio
		if err := cursor.Decode(&envio); err != nil {
			log.Printf("Error al decodificar envío: %v", err)
			continue
		}
		envios = append(envios, envio)
	}
	return envios, nil
}
func (repository *EnvioRepository) ObtenerPedidosFiltro(codigoEnvio string, estado string, fechaInicio time.Time, fechaFinal time.Time) ([]model.Pedido, error) {
	collection := repository.db.GetClient().Database("DirectoAlModelaje").Collection("Pedidos")
	filtro := bson.M{}
	if codigoEnvio != "" {
		filtro["codigo_envio"] = codigoEnvio
	}
	if estado != "" {
		filtro["estado"] = estado
	}
	if !fechaInicio.IsZero() {
		filtro["fecha_creacion"] = bson.M{"$gte": fechaInicio}
	}
	if !fechaFinal.IsZero() {
		filtro["fecha_creacion"] = bson.M{"$lte": fechaFinal}
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
			log.Printf("Error al decodificar pedido: %v", err)
			continue
		}
		pedidos = append(pedidos, pedido)
	}
	return pedidos, nil
}
