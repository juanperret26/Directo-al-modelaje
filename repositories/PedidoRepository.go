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

type PedidoRepositoryInterface interface {
	//Metodos para implementar en el service
	ObtenerPedidos() ([]model.Pedido, error)
	ObtenerPedidoPorId(id string) (model.Pedido, error)
	InsertarPedido(pedido model.Pedido) (*mongo.InsertOneResult, error)
	EliminarPedido(id primitive.ObjectID) (*mongo.UpdateResult, error)
	ActualizarPedido(pedido model.Pedido) (*mongo.UpdateResult, error)
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

	defer cursor.Close(context.Background())

	var pedidos []model.Pedido
	for cursor.Next(context.Background()) {
		var pedido model.Pedido
		err := cursor.Decode(&pedido)
		if err != nil {
			fmt.Printf("Error: %v\n", err)

		}
		pedidos = append(pedidos, pedido)
	}
	return pedidos, err
}
func (repository *PedidoRepository) ObtenerPedidoPorId(id string) (model.Pedido, error) {
	collection := repository.db.GetClient().Database("DirectoAlModelaje").Collection("Pedidos")
	objectID := utils.GetObjectIDFromStringID(id)
	filtro := bson.M{"_id": objectID}

	cursor, err := collection.Find(context.Background(), filtro)
	defer cursor.Close(context.Background())

	var pedido model.Pedido
	for cursor.Next(context.Background()) {
		err := cursor.Decode(&pedido)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		}
	}
	return pedido, err
}
func (repository *PedidoRepository) InsertarPedido(pedido model.Pedido) (*mongo.InsertOneResult, error) {
	collection := repository.db.GetClient().Database("DirectoAlModelaje").Collection("Pedidos")
	pedido.Fecha_creacion = time.Now()
	pedido.Actualizacion = time.Now()
	pedido.Estado = "Pendiente"
	resultado, err := collection.InsertOne(context.TODO(), pedido)
	return resultado, err
}

func (repository *PedidoRepository) EliminarPedido(id primitive.ObjectID) (*mongo.UpdateResult, error) {
	collection := repository.db.GetClient().Database("DirectoAlModelaje").Collection("Pedidos")
	filtro := bson.M{"_id": id}
	entidad := bson.M{"$set": bson.M{"estado": "Cancelado", "actualizacion": time.Now()}}
	resultado, err := collection.UpdateOne(context.TODO(), filtro, entidad)
	return resultado, err
}
func (repository *PedidoRepository) ActualizarPedido(pedido model.Pedido) (*mongo.UpdateResult, error) {
	collection := repository.db.GetClient().Database("DirectoAlModelaje").Collection("Pedidos")
	filtro := bson.M{"_id": pedido.Id}
	entidad := bson.M{"$set": bson.M{"estado": pedido.Estado, "actualizacion": time.Now()}}
	resultado, err := collection.UpdateOne(context.TODO(), filtro, entidad)
	return resultado, err
}
