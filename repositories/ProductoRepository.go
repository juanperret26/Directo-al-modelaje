package repositories

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/juanperret/Directo-al-modelaje/model"
	"github.com/juanperret/Directo-al-modelaje/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProductoRepositoryInterface interface {
	ObtenerProductos() ([]model.Producto, error)
	ObtenerProductosStockMinimo(tipoProducto string) ([]model.Producto, error)
	ObtenerProductoPorId(id string) (model.Producto, error)
	InsertarProducto(producto model.Producto) (*mongo.InsertOneResult, error)
	EliminarProducto(id string) (*mongo.DeleteResult, error)
	ActualizarProducto(id string) error
}

type ProductoRepository struct {
	db DB
}

func NewProductoRepository(db DB) *ProductoRepository {
	return &ProductoRepository{db: db}
}

func (repository *ProductoRepository) ObtenerProductos() ([]model.Producto, error) {
	collection := repository.db.GetClient().Database("DirectoAlModelaje").Collection("Productos")
	filtro := bson.M{}
	cursor, err := collection.Find(context.TODO(), filtro)

	defer cursor.Close(context.Background())

	var productos []model.Producto
	for cursor.Next(context.Background()) {
		var producto model.Producto
		err := cursor.Decode(&producto)
		if err != nil {
			fmt.Printf("Error: %v\n", err)

		}
		productos = append(productos, producto)
	}
	return productos, err
}

func (repository *ProductoRepository) ObtenerProductosStockMinimo(tipoProducto string) ([]model.Producto, error) {
	collection := repository.db.GetClient().Database("DirectoAlModelaje").Collection("Productos")
	filtro := bson.M{"tipo_producto": tipoProducto}

	cursor, err := collection.Find(context.Background(), filtro)
	defer cursor.Close(context.Background())
	var productos []model.Producto
	for cursor.Next(context.Background()) {
		var producto model.Producto
		err := cursor.Decode(&producto)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		}
		if producto.Stock < producto.Stock_minimo {
			productos = append(productos, producto)
		}
	}
	return productos, err
}

func (repository *ProductoRepository) ObtenerProductoPorId(id string) (model.Producto, error) {
	collection := repository.db.GetClient().Database("DirectoAlModelaje").Collection("Productos")
	objectID := utils.GetObjectIDFromStringID(id)
	filtro := bson.M{"_id": objectID}

	cursor, err := collection.Find(context.Background(), filtro)
	defer cursor.Close(context.Background())

	var producto model.Producto
	for cursor.Next(context.Background()) {
		err := cursor.Decode(&producto)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		}
	}
	return producto, err
}

func (repository *ProductoRepository) InsertarProducto(producto model.Producto) (*mongo.InsertOneResult, error) {

	collection := repository.db.GetClient().Database("DirectoAlModelaje").Collection("Productos")
	resultado, err := collection.InsertOne(context.TODO(), producto)
	return resultado, err
}

func (repository *ProductoRepository) ActualizarProducto(id string) error {
	//Actualizamos la fecha de actualizacion del producto
	objectID := utils.GetObjectIDFromStringID(id)
	collection := repository.db.GetClient().Database("DirectoAlModelaje").Collection("Productos")

	filtro := bson.M{"_id": objectID}
	producto, err := repository.ObtenerProductoPorId(id)
	if err != nil {
		return err
	}
	//Creo una operacion personalizada, para que no actualice nunca la fecha de creacion o el id del creador
	actualizacion := bson.M{
		"$set": bson.M{
			"nombre":        producto.Nombre,
			"tipoProducto":  producto.TipoProducto,
			"Peso_unitario": producto.Peso_unitario,
			"precio":        producto.Precio,
			"stock":         producto.Stock,
			"stock_minimo":  producto.Stock_minimo,
			"actualizacion": time.Now(),
			"creacion":      producto.Creacion,
		},
	}

	operacion, err := collection.UpdateOne(context.Background(), filtro, actualizacion)

	if operacion.MatchedCount == 0 {
		return errors.New("no se encontró el producto a actualizar")
	}

	return err
}

// // func (repository *PedidoRepository) EliminarPedido(id string) (*mongo.UpdateResult, error) {
// 	objectID := utils.GetObjectIDFromStringID(id)
// 	collection := repository.db.GetClient().Database("DirectoAlModelaje").Collection("Pedidos")
// 	filtro := bson.M{"_id": objectID}
// 	entidad := bson.M{"$set": bson.M{"estado": "Cancelado", "actualizacion": time.Now()}}
// 	resultado, err := collection.UpdateOne(context.TODO(), filtro, entidad)
// 	return resultado, err
// }

func (repository *ProductoRepository) EliminarProducto(id string) (*mongo.DeleteResult, error) {
	objectID := utils.GetObjectIDFromStringID(id)
	collection := repository.db.GetClient().Database("DirectoAlModelaje").Collection("Productos")
	filtro := bson.M{"_id": objectID}
	resultado, err := collection.DeleteOne(context.Background(), filtro)
	return resultado, err
}
