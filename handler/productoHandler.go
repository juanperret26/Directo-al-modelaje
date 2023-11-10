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
	if productos == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No se encontraron productos"})
	} else {
		log.Printf("[handler:ProductoHandler] [method:ObtenerProductos] [productos:%v] [cantidad:%d]", productos, len(productos))
		c.JSON(http.StatusOK, productos)
	}

}
func (handler *ProductoHandler) ObtenerProductosStockMinimo(c *gin.Context) {
	tipoProducto := c.Param("tipoProducto")
	productos := handler.ProductoService.ObtenerProductosStockMinimo(tipoProducto)
	if productos == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No se encontraron productos"})
	} else {
		c.JSON(http.StatusOK, productos)
	}
}
func (handler *ProductoHandler) ObtenerProductoPorId(c *gin.Context) {
	id := c.Param("id")
	producto := handler.ProductoService.ObtenerProductoPorId(id)
	if producto == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "producto no encontrado"})
	} else {
		c.JSON(http.StatusOK, producto)
	}
}

func (handler *ProductoHandler) InsertarProducto(c *gin.Context) {
	var producto dto.Producto
	err := c.ShouldBindJSON(&producto)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		resultado := handler.ProductoService.InsertarProducto(&producto)
		if resultado != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": resultado.Error()})
		} else {
			c.JSON(http.StatusOK, gin.H{"error": resultado.Error()})
		}
	}
}

func (handler *ProductoHandler) EliminarProducto(c *gin.Context) {
	id := c.Param("id")
	resultado := handler.ProductoService.EliminarProducto(id)
	if resultado != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": resultado.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"error": resultado.Error()})
	}
}

func (handler *ProductoHandler) ActualizarProducto(c *gin.Context) {
	var producto dto.Producto
	producto.Id = c.Param("id")
	err := c.ShouldBindJSON(&producto)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		resultado := handler.ProductoService.ActualizarProducto(&producto)
		if resultado != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": resultado.Error()})
		} else {
			c.JSON(http.StatusOK, gin.H{"error": resultado.Error()})
		}
	}
}
