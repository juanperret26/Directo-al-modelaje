// Crear struct, new objeto y metodos
package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/juanperret26/Directo-al-modelaje/go/dto"
	"github.com/juanperret26/Directo-al-modelaje/go/services"
)

type ProductoHandler struct {
	ProductoService services.ProductoInterface
}

func NewProductoHandler(productoService services.ProductoInterface) *ProductoHandler {
	return &ProductoHandler{ProductoService: productoService}
}

func (handler *ProductoHandler) ObtenerProductos(c *gin.Context) {
	// Obtener el parámetro de query 'stockMinimo' y convertirlo a int
	stockMinimoQuery := c.DefaultQuery("stockMinimo", "0")
	stockMinimo, err := strconv.Atoi(stockMinimoQuery)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "stockMinimo debe ser un número entero"})
		return
	}
	// Llamar al servicio con el filtro de stock mínimo
	productos := handler.ProductoService.ObtenerProductos(stockMinimo)

	// Responder con los productos obtenidos
	c.JSON(http.StatusOK, productos)
}

//obtener producto por id
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

	if err := c.ShouldBindJSON(&producto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} else {
		resultado := handler.ProductoService.InsertarProducto(&producto)
		if resultado != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "No se pudo insertar el producto"})
		} else {
			c.JSON(http.StatusCreated, gin.H{"error": "Creado correctamente"})
		}
	}
}

func (handler *ProductoHandler) EliminarProducto(c *gin.Context) {
	id := c.Param("id")
	resultado := handler.ProductoService.EliminarProducto(id)
	if resultado != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": resultado.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"error": "Producto eliminado correctamente"})
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
			c.JSON(http.StatusOK, gin.H{"error": producto})
		}
	}
}
