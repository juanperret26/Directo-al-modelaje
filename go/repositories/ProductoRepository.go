package repositories

import (
	"context"
	"fmt"

	"github.com/juanperret/Directo-al-modelaje/go/model"
	"github.com/juanperret/Directo-al-modelaje/go/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProductoRepositoryInterface interface {
	ObtenerProductos() ([]model.Producto, error)
	ObtenerProductosStockMinimo(tipoProducto string) ([]model.Producto, error)
	ObtenerProductoPorId(id string) (model.Producto, error)
	InsertarProducto(producto model.Producto) (*mongo.InsertOneResult, error)
	EliminarProducto(id string) (*mongo.DeleteResult, error)
	ActualizarProducto(Producto model.Producto) (*mongo.UpdateResult, error)
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

func (repo ProductoRepository) ActualizarProducto(Producto model.Producto) (*mongo.UpdateResult, error) {
	lista := repo.db.GetClient().Database("DirectoAlModelaje").Collection("Productos")
	filtro := bson.M{"_id": Producto.Id}
	entity := bson.M{"$set": Producto}
	resultado, err := lista.UpdateOne(context.TODO(), filtro, entity)
	return resultado, err
}

func (repository *ProductoRepository) EliminarProducto(id string) (*mongo.DeleteResult, error) {
	objectID := utils.GetObjectIDFromStringID(id)
	collection := repository.db.GetClient().Database("DirectoAlModelaje").Collection("Productos")
	filtro := bson.M{"_id": objectID}
	resultado, err := collection.DeleteOne(context.Background(), filtro)
	return resultado, err
}
