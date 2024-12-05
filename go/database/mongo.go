package database

import (
	"context"
	"log" // Importar log para depuración

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDB struct {
	Client *mongo.Client
}


func NewMongoDB() *MongoDB {
	instancia := &MongoDB{}
	instancia.Connect()
	log.Println("Creando instancia de MongoDB...")
	err := instancia.Connect()

	if err != nil {
		log.Fatalf("Error al conectar a MongoDB: %v\n", err)
	}

	log.Println("Conexión a MongoDB exitosa.")
	return instancia
}

func (mongoDB *MongoDB) GetClient() *mongo.Client {
	if mongoDB.Client == nil {
		log.Println("Advertencia: El cliente de MongoDB no está inicializado.")
	}
	return mongoDB.Client
}

func (mongoDB *MongoDB) Connect() error {
	log.Println("Intentando conectar a MongoDB...")
	clientOptions := options.Client().ApplyURI("mongodb://mongodb-container:27017")

	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Printf("Error al crear el cliente de MongoDB: %v\n", err)
		return err
	}

	log.Println("Cliente de MongoDB creado, verificando la conexión con Ping...")
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Printf("Error al hacer Ping a MongoDB: %v\n", err)
		return err
	}

	log.Println("Ping exitoso. Conexión establecida.")
	mongoDB.Client = client

	return nil
}

func (mongoDB *MongoDB) Disconnect() error {
	log.Println("Desconectando cliente de MongoDB...")
	err := mongoDB.Client.Disconnect(context.Background())
	if err != nil {
		log.Printf("Error al desconectar el cliente de MongoDB: %v\n", err)
		return err
	}
	log.Println("Cliente de MongoDB desconectado exitosamente.")
	return nil
}