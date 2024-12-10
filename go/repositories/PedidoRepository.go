package repositories

import (
	"context"
	"fmt"
	"time"

	"github.com/juanperret26/Directo-al-modelaje/go/model"
	"github.com/juanperret26/Directo-al-modelaje/go/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type PedidoRepositoryInterface interface {
	// Métodos para implementar en el service
	ObtenerPedidos() ([]model.Pedido, error)
	ObtenerPedidoPorId(id string) (model.Pedido, error)
	InsertarPedido(pedido model.Pedido) (*mongo.InsertOneResult, error)
	EliminarPedido(id string) (*mongo.UpdateResult, error)
	ActualizarPedido(pedido model.Pedido) (*mongo.UpdateResult, error)
	ObtenerCantidadPedidosPorEstado(estado string) (int, error)
	ObtenerPedidosPorEstado(estado string) ([]model.Pedido, error)
}

type PedidoRepository struct {
	db DB
}

func NewPedidoRepository(db DB) *PedidoRepository {
	return &PedidoRepository{db: db}
}

func (repository *PedidoRepository) ObtenerPedidos() ([]model.Pedido, error) {
	collection := repository.db.GetClient().Database("DirectoAlModelaje").Collection("Pedidos")
	filtro := bson.M{}
	cursor, err := collection.Find(context.TODO(), filtro)
	if err != nil {
		return nil, err // Asegura manejo de errores temprano
	}
	defer cursor.Close(context.Background()) // Siempre cerrar el cursor

	var pedidos []model.Pedido
	for cursor.Next(context.Background()) {
		var pedido model.Pedido
		if err := cursor.Decode(&pedido); err != nil {
			fmt.Printf("Error decodificando pedido: %v\n", err)
			continue // Ignora errores individuales pero sigue procesando
		}
		pedidos = append(pedidos, pedido)
	}
	return pedidos, cursor.Err() // Retorna error si ocurrió al recorrer el cursor
}

func (repository *PedidoRepository) ObtenerPedidoPorId(id string) (model.Pedido, error) {
	collection := repository.db.GetClient().Database("DirectoAlModelaje").Collection("Pedidos")
	objectID := utils.GetObjectIDFromStringID(id)
	filtro := bson.M{"_id": objectID}

	var pedido model.Pedido
	err := collection.FindOne(context.Background(), filtro).Decode(&pedido) // Usa FindOne para optimizar
	if err != nil {
		return model.Pedido{}, err
	}
	return pedido, nil
}

func (repository *PedidoRepository) InsertarPedido(pedido model.Pedido) (*mongo.InsertOneResult, error) {
	collection := repository.db.GetClient().Database("DirectoAlModelaje").Collection("Pedidos")
	pedido.Fecha_creacion = time.Now()
	pedido.Actualizacion = time.Now()
	pedido.Estado = "Pendiente" // Estado predeterminado
	return collection.InsertOne(context.Background(), pedido)
}

func (repository *PedidoRepository) EliminarPedido(id string) (*mongo.UpdateResult, error) {
	objectID := utils.GetObjectIDFromStringID(id)
	collection := repository.db.GetClient().Database("DirectoAlModelaje").Collection("Pedidos")
	filtro := bson.M{"_id": objectID} // Solo permite cancelar pedidos pendientes
	entidad := bson.M{"$set": bson.M{"estado": "Cancelado", "actualizacion": time.Now()}}
	return collection.UpdateOne(context.TODO(), filtro, entidad)
}

func (repository *PedidoRepository) ActualizarPedido(pedido model.Pedido) (*mongo.UpdateResult, error) {
	collection := repository.db.GetClient().Database("DirectoAlModelaje").Collection("Pedidos")
	filtro := bson.M{"_id": pedido.Id}
	entidad := bson.M{
		"$set": bson.M{
			"estado":       pedido.Estado,
			"actualizacion": time.Now(),
		},
	}

	result, err := collection.UpdateOne(context.TODO(), filtro, entidad)
	if err != nil {
		return nil, fmt.Errorf("error al actualizar el pedido: %w", err)
	}
	if result.MatchedCount == 0 {
		return nil, fmt.Errorf("no se encontró un pedido con el ID: %s", pedido.Id)
	}

	return result, nil
}


func (repository *PedidoRepository) ObtenerCantidadPedidosPorEstado(estado string) (int, error) {
	collection := repository.db.GetClient().Database("DirectoAlModelaje").Collection("Pedidos")
	filtro := bson.M{"estado": estado}

	cantidad, err := collection.CountDocuments(context.Background(), filtro)
	if err != nil {
		return 0, err
	}
	return int(cantidad), nil
}

func (repository *PedidoRepository) ObtenerPedidosPorEstado(estado string) ([]model.Pedido, error) {
	collection := repository.db.GetClient().Database("DirectoAlModelaje").Collection("Pedidos")
	filtro := bson.M{"estado": estado}

	cursor, err := collection.Find(context.Background(), filtro)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var pedidos []model.Pedido
	for cursor.Next(context.Background()) {
		var pedido model.Pedido
		if err := cursor.Decode(&pedido); err != nil {
			fmt.Printf("Error decodificando pedido: %v\n", err)
			continue
		}
		pedidos = append(pedidos, pedido)
	}
	return pedidos, cursor.Err()
}
