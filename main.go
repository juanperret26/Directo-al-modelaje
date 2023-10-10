package main

import (
	//Agregar imports de todas las clases, handlers, middlewares, etc
	"log"

	"github.com/gin-gonic/gin"
	//"github.com/juanperret/Directo-al-modelaje/clients"
	"github.com/juanperret/Directo-al-modelaje/handler"
	//"github.com/juanperret/Directo-al-modelaje/middlewares"
	"github.com/juanperret/Directo-al-modelaje/repositories"
	"github.com/juanperret/Directo-al-modelaje/services"
)

var (
	//Agregar handlers
	envioHandler *handler.EnvioHandler
	router       *gin.Engine
)

func main() {
	router = gin.Default()
	//Iniciar objetos de handler
	dependencies()
	//Iniciar rutas
	mappingRoutes()

	log.Println("Iniciando el servidor...")
	router.Run(":8080")
}

func mappingRoutes() {
	//cliente para api externa
	// var authClient clients.AuthClientInterface
	// authClient = clients.NewAuthClient()
	// //creacion de middleware de autenticacion
	// authMiddleware := middlewares.NewAuthMiddleware(authClient)

	// //Listado de rutas
	groupEnvio := router.Group("/envios")
	//Uso del middleware para todas las rutas del grupo de rutas y hago todos los POST, GET y DELETE
	// groupEnvio.Use(authMiddleware.ValidateToken)
	groupEnvio.GET("/", envioHandler.GetEnvios)
	groupEnvio.GET("/:id", envioHandler.GetEnvio)
	groupEnvio.POST("/", envioHandler.InsertarEnvio)
}

// Generacion de los objetos que se van a usar en la api
func dependencies() {
	//Definicion de variables de interface
	var database repositories.DB
	var envioRepository repositories.EnvioRepositoryInterface
	var envioService services.EnvioInterface

	//Creamos los objetos reales y los pasamos como parametro
	database = repositories.NewMongoDB()
	envioRepository = repositories.NewEnvioRepository(database)
	envioService = services.NewEnvioService(envioRepository)
	envioHandler = handler.NewEnvioHandler(envioService)

}
