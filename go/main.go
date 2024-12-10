package main

import (
	//Agregar imports de todas las clases, handlers, middlewares, etc

	"log"

	"github.com/gin-gonic/gin"

	"github.com/juanperret26/Directo-al-modelaje/go/clients"
	"github.com/juanperret26/Directo-al-modelaje/go/handler"
	"github.com/juanperret26/Directo-al-modelaje/go/middlewares"

	"github.com/juanperret26/Directo-al-modelaje/go/repositories"
	"github.com/juanperret26/Directo-al-modelaje/go/services"
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
	authClient := &clients.AuthClient{}
	authMiddleware := middlewares.NewAuthMiddleware(authClient)

	router.Use(middlewares.CORSMiddleware())
	router.Use(authMiddleware.ValidateToken)
	// //Listado de rutas
	groupEnvio := router.Group("/envios")
	groupCamion := router.Group("/camiones")
	groupProducto := router.Group("/productos")
	groupPedido := router.Group("/pedidos")
	//Uso del middleware para todas las rutas del grupo de rutas y hago todos los POST, GET y DELETE
	groupEnvio.Use(authMiddleware.ValidateToken)
	groupEnvio.GET("/", envioHandler.ObtenerEnvios)
	groupEnvio.GET("/:id", envioHandler.ObtenerEnvioPorId)
	groupEnvio.GET("/estado/:estado", envioHandler.ObtenerCantidadEnviosPorEstado)
	groupEnvio.GET("/envios/fechas", envioHandler.ObtenerBeneficiosEntreFechas)
	groupEnvio.POST("/", envioHandler.InsertarEnvio)
	groupEnvio.PUT("/:id/parada", envioHandler.AgregarParada)
	groupEnvio.DELETE("/:id", envioHandler.EliminarEnvio)
	groupEnvio.PUT("/", envioHandler.ActualizarEnvio)
	groupEnvio.PUT("/:id/iniciar", envioHandler.IniciarViaje)

	//Camiones
	groupCamion.Use(authMiddleware.ValidateToken)
	groupCamion.GET("/", camionHandler.ObtenerCamiones)
	groupCamion.GET("/:patente", camionHandler.ObtenerCamionPorPatente)
	groupCamion.POST("/", camionHandler.InsertarCamion)
	groupCamion.DELETE("/:id", camionHandler.EliminarCamion)
	groupCamion.PUT("/", camionHandler.ActualizarCamion)
	//Productos
	groupProducto.Use(authMiddleware.ValidateToken)
	groupProducto.GET("/", productoHandler.ObtenerProductos)
	groupProducto.GET("/:id", productoHandler.ObtenerProductoPorId)
	// groupProducto.GET("/:tipoProducto", productoHandler.ObtenerProductosStockMinimo)
	groupProducto.POST("/", productoHandler.InsertarProducto)
	groupProducto.DELETE("/:id", productoHandler.EliminarProducto)
	groupProducto.PUT("/", productoHandler.ActualizarProducto)
	//Pedidos
	groupPedido.Use(authMiddleware.ValidateToken)
	groupPedido.GET("/", pedidoHandler.ObtenerPedidos)
	groupPedido.GET("/:id", pedidoHandler.ObtenerPedidoPorId)
	groupPedido.GET("/estado/:estado", pedidoHandler.ObtenerCantidadPedidosPorEstado)
	groupPedido.POST("/", pedidoHandler.InsertarPedido)
	groupPedido.DELETE("/:id", pedidoHandler.EliminarPedido)
	groupPedido.PUT("/:id", pedidoHandler.AceptarPedido)

}

// Generacion de los objetos que se van a usar en la api
func dependencies() {

	//Definicion de variables de interface
	database := repositories.NewMongoDB()
	var camionRepository repositories.CamionRepositoryInterface
	var camionService services.CamionInterface
	var envioRepository repositories.EnvioRepositoryInterface
	var envioService services.EnvioInterface
	var pedidoRepository repositories.PedidoRepositoryInterface
	var pedidoService services.PedidoInterface
	var productoRepository repositories.ProductoRepositoryInterface
	var productoService services.ProductoInterface

	//Productos
	productoRepository = repositories.NewProductoRepository(database)
	productoService = services.NewProductoService(productoRepository)
	productoHandler = handler.NewProductoHandler(productoService)

	// //Pedidos
	pedidoRepository = repositories.NewPedidoRepository(database)
	pedidoService = services.NewPedidoService(pedidoRepository, productoRepository)
	pedidoHandler = handler.NewPedidoHandler(pedidoService)

	//Envio
	envioRepository = repositories.NewEnvioRepository(database)
	envioService = services.NewEnvioService(envioRepository, camionRepository, pedidoRepository, productoRepository)
	envioHandler = handler.NewEnvioHandler(envioService)

	//Camiones
	camionRepository = repositories.NewCamionRepository(database)
	camionService = services.NewCamionService(camionRepository, envioRepository)
	camionHandler = handler.NewCamionHandler(camionService)
}
