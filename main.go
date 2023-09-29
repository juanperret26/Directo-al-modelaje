package main

import (
	"log"

	"github.com/gin-gonic/gin"
	//Agregar imports de todas las clases, handlers, middlewares, etc

)

var (
	//Agregar handlers
	router      *gin.Engine
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
	//var authClient clients.AuthClientInterface
	//authClient = clients.NewAuthClient()
	//creacion de middleware de autenticacion
	//authMiddleware := middlewares.NewAuthMiddleware(authClient)

	//Listado de rutas
	group := router.Group("/aulas")
	//Uso del middleware para todas las rutas del grupo de rutas y hago todos los POST, GET y DELETE
	
}

// Generacion de los objetos que se van a usar en la api
func dependencies() {
	//Definicion de variables de interface
	//var database repositories.DB
	//var aulaRepository repositories.AulaRepositoryInterface
	//var aulaService services.AulaInterface

	//Creamos los objetos reales y los pasamos como parametro
	//database = repositories.NewMongoDB()
	//aulaRepository = repositories.NewAulaRepository(database)
	//aulaService = services.NewAulaService(aulaRepository)
	//aulaHandler = handlers.NewAulaHandler(aulaService)

}