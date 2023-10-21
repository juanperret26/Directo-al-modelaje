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
	envioHandler    *handler.EnvioHandler
	camionHandler   *handler.CamionHandler
	productoHandler *handler.ProductoHandler
	pedidoHandler   *handler.PedidoHandler
	// Agregar router
	router *gin.Engine
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
	groupCamion := router.Group("/camiones")
	groupProducto := router.Group("/productos")
	groupPedido := router.Group("/pedidos")
	//Uso del middleware para todas las rutas del grupo de rutas y hago todos los POST, GET y DELETE
	//groupEnvio.Use(authMiddleware.ValidateToken)
	groupEnvio.GET("/", envioHandler.ObtenerEnvios)
	groupEnvio.GET("/:id", envioHandler.ObtenerEnvioPorId)
	groupEnvio.POST("/", envioHandler.InsertarEnvio)
	groupEnvio.DELETE("/:id", envioHandler.EliminarEnvio)
	groupEnvio.PUT("/", envioHandler.ActualizarEnvio)

	//Camiones
	//grupoCamion.Use(authMiddleware.ValidateToken)
	groupCamion.GET("/", camionHandler.ObtenerCamiones)
	groupCamion.GET("/:id", camionHandler.ObtenerCamionPorId)
	groupCamion.POST("/", camionHandler.InsertarCamion)
	groupCamion.DELETE("/:id", camionHandler.EliminarCamion)
	groupCamion.PUT("/", camionHandler.ActualizarCamion)
	//Productos
	//grupoProducto.Use(authMiddleware.ValidateToken)
	groupProducto.GET("/", productoHandler.ObtenerProductos)
	groupProducto.GET("/:id", productoHandler.ObtenerProductoPorId)
	groupProducto.POST("/", productoHandler.InsertarProducto)
	groupProducto.DELETE("/:id", productoHandler.EliminarProducto)
	groupProducto.PUT("/", productoHandler.ActualizarProducto)
	//Pedidos
	//grupoPedido.Use(authMiddleware.ValidateToken)
	groupPedido.GET("/", pedidoHandler.ObtenerPedidos)
	groupPedido.GET("/:id", pedidoHandler.ObtenerPedidoPorId)
	groupPedido.POST("/", pedidoHandler.InsertarPedido)
	groupPedido.DELETE("/:id", pedidoHandler.EliminarPedido)

}

// Generacion de los objetos que se van a usar en la api
func dependencies() {

	//Definicion de variables de interface
	//Envios
	var database repositories.DB
	database = repositories.NewMongoDB()
	var envioRepository repositories.EnvioRepositoryInterface
	var envioService services.EnvioInterface
	envioRepository = repositories.NewEnvioRepository(database)
	envioService = services.NewEnvioService(envioRepository)
	envioHandler = handler.NewEnvioHandler(envioService)

	//Camiones
	var camionRepository repositories.CamionRepositoryInterface
	var camionService services.CamionInterface
	camionRepository = repositories.NewCamionRepository(database)
	camionService = services.NewCamionService(camionRepository)
	camionHandler = handler.NewCamionHandler(camionService)

	//Productos
	var productoRepository repositories.ProductoRepositoryInterface
	var productoService services.ProductoInterface
	productoRepository = repositories.NewProductoRepository(database)
	productoService = services.NewProductoService(productoRepository)
	productoHandler = handler.NewProductoHandler(productoService)

	// //Pedidos
	var pedidoRepository repositories.PedidoRepositoryInterface
	var pedidoService services.PedidoInterface
	pedidoRepository = repositories.NewPedidoRepository(database)
	pedidoService = services.NewPedidoService(pedidoRepository)
	pedidoHandler = handler.NewPedidoHandler(pedidoService)

}
