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
    // Convertir el ID a ObjectID
    objectID:= utils.GetObjectIDFromStringID(id)
    collection := repository.db.GetClient().Database("DirectoAlModelaje").Collection("Productos")
    filtro := bson.M{"_id": objectID}

    var producto model.Producto
    result := collection.FindOne(context.Background(), filtro)
    // Decodificar el documento encontrado
    if err := result.Decode(&producto); err != nil {
        return model.Producto{}, fmt.Errorf("error al decodificar producto: %w", err)
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
	if collection == nil {
		return nil, fmt.Errorf("la colección de productos es nula")
	}

	if producto.Id.IsZero() {
		return nil, fmt.Errorf("el ID del producto es inválido o está vacío")
	}

	updates := bson.M{}
	if producto.Nombre != "" {
		updates["nombre"] = producto.Nombre
	}
	if producto.TipoProducto != "" {
		updates["tipo_producto"] = producto.TipoProducto
	}
	if producto.Peso_unitario > 0 {
		updates["peso"] = producto.Peso_unitario
	}
	if producto.Precio > 0 {
		updates["precio"] = producto.Precio
	}
	if producto.Stock >= 0 {
		updates["stock"] = producto.Stock
	}
	if producto.Stock_minimo >= 0 {
		updates["stock_minimo"] = producto.Stock_minimo
	}
	// Actualizar siempre la fecha de actualización
	updates["actualizacion"] = time.Now()

	if len(updates) == 0 {
		return nil, fmt.Errorf("no hay campos válidos para actualizar")
	}

	filtro := bson.M{"_id": producto.Id}
	entidad := bson.M{"$set": updates}

	resultado, err := collection.UpdateOne(context.TODO(), filtro, entidad)
	if err != nil {
		return nil, fmt.Errorf("error al actualizar producto: %w", err)
	}

	if resultado.MatchedCount == 0 {
		return nil, fmt.Errorf("no se encontró el producto con id %s para actualizar", producto.Id.Hex())
	}

	return resultado, nil
}


func (repository *ProductoRepository) EliminarProducto(id string) (*mongo.DeleteResult, error) {
    // Convertir el ID a ObjectID
    objectID := utils.GetObjectIDFromStringID(id)
    collection := repository.db.GetClient().Database("DirectoAlModelaje").Collection("Productos")
    filtro := bson.M{"_id": objectID}

    // Eliminar el documento
    resultado, err := collection.DeleteOne(context.Background(), filtro)
    if err != nil {
        return nil, fmt.Errorf("error al eliminar producto: %w", err)
    }
    return resultado, nil
}
