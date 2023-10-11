// Crear struct, new objeto y metodos
package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/juanperret/Directo-al-modelaje/dto"
	"github.com/juanperret/Directo-al-modelaje/services"
)

type ProductoHandler struct {
	ProductoService services.ProductoInterface
}

func NewProductoHandler(productoService services.ProductoInterface) *ProductoHandler {
	return &ProductoHandler{ProductoService: productoService}
}

func (handler *ProductoHandler) ObtenerProductos(c *gin.Context) {
	productos := handler.ProductoService.ObtenerProductos()

	log.Printf("[handler:ProductoHandler] [method:ObtenerProductos] [productos:%v] [cantidad:%d]", productos, len(productos))
	c.JSON(http.StatusOK, productos)

}

func (handler *ProductoHandler) ObtenerProductoPorId(c *gin.Context) {
	id := c.Param("id")
	producto := handler.ProductoService.ObtenerProductoPorId(id)
	c.JSON(http.StatusOK, producto)

}

func (handler *ProductoHandler) InsertarProducto(c *gin.Context) {
	var producto dto.Producto
	if err := c.ShouldBindJSON(&producto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return

	}
}

func (handler *ProductoHandler) EliminarProducto(c *gin.Context) {
	id := c.Param("id")
	resultado := handler.ProductoService.EliminarProducto(id)
	c.JSON(http.StatusOK, resultado)
}

func (handler *ProductoHandler) ActualizarProducto(c *gin.Context) {
	var producto dto.Producto
	if err := c.ShouldBindJSON(&producto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	producto.Id = c.Param("id")
	resultado := handler.ProductoService.ActualizarProducto(&producto)
	c.JSON(http.StatusOK, resultado)

}
