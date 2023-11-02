package main

import (
	//Agregar imports de todas las clases, handlers, middlewares, etc
	"html/template"
	"log"
	"net/http"

	"github.com/gin-contrib/cors"
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

	tmpl *template.Template
)

func main() {
	router = gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:8080"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	router.Use(cors.New(config))

	//Iniciar objetos de handler
	dependencies()
	//Iniciar rutas
	mappingRoutes()

	router.LoadHTMLGlob("html/*")

	router.Static("/static", "./static")

	log.Println("Iniciando el servidor...")
	router.Run(":8080")
}

func mappingRoutes() {
	//cliente para api externa
	//router.Use(middlewares.CORSMiddleware())

	// //Listado de rutas
	groupEnvio := router.Group("/envios")
	groupCamion := router.Group("/camiones")
	groupProducto := router.Group("/productos")
	groupPedido := router.Group("/pedidos")
	//Uso del middleware para todas las rutas del grupo de rutas y hago todos los POST, GET y DELETE
	//groupEnvio.Use(authMiddleware.ValidateToken)
	groupEnvio.GET("/", envioHandler.ObtenerEnvios)
	groupEnvio.GET("/:id", envioHandler.ObtenerEnvioPorId)
	groupEnvio.GET("/:estado", envioHandler.ObtenerCantidadEnviosPorEstado)
	groupEnvio.POST("/", envioHandler.InsertarEnvio)
	groupEnvio.DELETE("/:id", envioHandler.EliminarEnvio)
	groupEnvio.PUT("/", envioHandler.ActualizarEnvio)

	//Camiones
	//grupoCamion.Use(authMiddleware.ValidateToken)
	groupCamion.GET("/", camionHandler.ObtenerCamiones)
	groupCamion.GET("/:patente", camionHandler.ObtenerCamionPorPatente)
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
	groupPedido.GET("/:estado", pedidoHandler.ObtenerCantidadPedidosPorEstado)
	groupPedido.POST("/", pedidoHandler.InsertarPedido)
	groupPedido.DELETE("/:id", pedidoHandler.EliminarPedido)
	groupPedido.PUT("/:id", pedidoHandler.AceptarPedido)

	//rutas html
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	router.GET("/htmlproductos", func(c *gin.Context) {
		c.HTML(http.StatusOK, "productos.html", nil)
	})
	router.GET("/htmlpedidos", func(c *gin.Context) {
		c.HTML(http.StatusOK, "pedidos.html", nil)
	})
	router.GET("/htmlinformes", func(c *gin.Context) {
		c.HTML(http.StatusOK, "informes.html", nil)
	})
	router.GET("/htmlenvios", func(c *gin.Context) {
		c.HTML(http.StatusOK, "envios.html", nil)
	})
	router.GET("/htmlcamiones", func(c *gin.Context) {
		c.HTML(http.StatusOK, "camiones.html", nil)
	})
	router.GET("/formProductos", func(c *gin.Context) {
		c.HTML(http.StatusOK, "formProductos.html", nil)
	})
	router.GET("/formPedidos", func(c *gin.Context) {
		c.HTML(http.StatusOK, "formPedidos.html", nil)
	})
	router.GET("/formEnvios", func(c *gin.Context) {
		c.HTML(http.StatusOK, "formEnvios.html", nil)
	})
	router.GET("/formCamiones", func(c *gin.Context) {
		c.HTML(http.StatusOK, "formCamiones.html", nil)
	})
}

// Generacion de los objetos que se van a usar en la api
func dependencies() {

	//Definicion de variables de interface
	//Envios
	var database repositories.DB
	database = repositories.NewMongoDB()

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
	pedidoService = services.NewPedidoService(pedidoRepository, productoRepository)
	pedidoHandler = handler.NewPedidoHandler(pedidoService)

	var envioRepository repositories.EnvioRepositoryInterface
	var envioService services.EnvioInterface
	envioRepository = repositories.NewEnvioRepository(database)
	envioService = services.NewEnvioService(envioRepository, camionRepository, pedidoRepository, productoRepository)
	envioHandler = handler.NewEnvioHandler(envioService)
}
