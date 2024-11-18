package repositories

import (
	"context"
	"fmt"

	"github.com/juanperret26/Directo-al-modelaje/go/model"
	"github.com/juanperret26/Directo-al-modelaje/go/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProductoRepositoryInterface interface {
	ObtenerProductos(stockMinimo int) ([]model.Producto, error)
	ObtenerProductoPorId(id string) (model.Producto, error)
	InsertarProducto(producto model.Producto) (*mongo.InsertOneResult, error)
	EliminarProducto(id string) (*mongo.DeleteResult, error)
	ActualizarProducto(producto model.Producto) (*mongo.UpdateResult, error)
}

type ProductoRepository struct {
	db DB
}

func NewProductoRepository(db DB) *ProductoRepository {
	return &ProductoRepository{db: db}
}

func (repository *ProductoRepository) ObtenerProductos(stockMinimo int) ([]model.Producto, error) {
	collection := repository.db.GetClient().Database("DirectoAlModelaje").Collection("Productos")

	// Crear el filtro según el stockMinimo
	filtro := bson.M{}
	if stockMinimo > 0 {
		filtro = bson.M{"stock": bson.M{"$gte": stockMinimo}}
	}

	cursor, err := collection.Find(context.TODO(), filtro)
	if err != nil {
		return nil, fmt.Errorf("error al obtener productos: %w", err)
	}
	defer cursor.Close(context.Background())

	var productos []model.Producto
	for cursor.Next(context.Background()) {
		var producto model.Producto
		if err := cursor.Decode(&producto); err != nil {
			fmt.Printf("Error al decodificar producto: %v\n", err)
			continue
		}
		productos = append(productos, producto)
	}

	if err := cursor.Err(); err != nil {
		return nil, fmt.Errorf("error en el cursor al iterar productos: %w", err)
	}

	return productos, nil
}

func (repository *ProductoRepository) ObtenerProductoPorId(id string) (model.Producto, error) {
	collection := repository.db.GetClient().Database("DirectoAlModelaje").Collection("Productos")
	objectID := utils.GetObjectIDFromStringID(id)
	filtro := bson.M{"_id": objectID}

	var producto model.Producto
	err := collection.FindOne(context.Background(), filtro).Decode(&producto)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return model.Producto{}, fmt.Errorf("producto no encontrado con ID: %s", id)
		}
		return model.Producto{}, fmt.Errorf("error al obtener producto por ID: %w", err)
	}

	return producto, nil
}

func (repository *ProductoRepository) InsertarProducto(producto model.Producto) (*mongo.InsertOneResult, error) {
	collection := repository.db.GetClient().Database("DirectoAlModelaje").Collection("Productos")
	resultado, err := collection.InsertOne(context.TODO(), producto)
	if err != nil {
		return nil, fmt.Errorf("error al insertar producto: %w", err)
	}
	return resultado, nil
}

func (repository *ProductoRepository) ActualizarProducto(producto model.Producto) (*mongo.UpdateResult, error) {
	collection := repository.db.GetClient().Database("DirectoAlModelaje").Collection("Productos")
	filtro := bson.M{"_id": producto.Id}
	entity := bson.M{"$set": producto}
	resultado, err := collection.UpdateOne(context.TODO(), filtro, entity)
	if err != nil {
		return nil, fmt.Errorf("error al actualizar producto: %w", err)
	}
	return resultado, nil
}

func (repository *ProductoRepository) EliminarProducto(id string) (*mongo.DeleteResult, error) {
	objectID := utils.GetObjectIDFromStringID(id)
	collection := repository.db.GetClient().Database("DirectoAlModelaje").Collection("Productos")
	filtro := bson.M{"_id": objectID}

	resultado, err := collection.DeleteOne(context.Background(), filtro)
	if err != nil {
		return nil, fmt.Errorf("error al eliminar producto: %w", err)
	}
	return resultado, nil
}
