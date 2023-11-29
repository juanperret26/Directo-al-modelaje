// Crear struct, new objeto y metodos
package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
<<<<<<< HEAD:go/handler/pedidoHandler.go
	"github.com/juanperret26/Directo-al-modelaje/go/dto"
	"github.com/juanperret26/Directo-al-modelaje/go/services"
=======
	"github.com/juanperret/Directo-al-modelaje/dto"
	"github.com/juanperret/Directo-al-modelaje/services"
>>>>>>> parent of ad42c9c (prueba cambio de directorio):handler/pedidoHandler.go
)

type PedidoHandler struct {
	pedidoService services.PedidoInterface
}

func NewPedidoHandler(pedidoService services.PedidoInterface) *PedidoHandler {
	return &PedidoHandler{
		pedidoService: pedidoService,
	}
}
func (handler *PedidoHandler) ObtenerPedidos(c *gin.Context) {
	pedidos := handler.pedidoService.ObtenerPedidos()
	if pedidos == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No se encontraron pedidos"})
	} else {
		log.Printf("[handler:PedidoHandler] [method:ObtenerPedidos] [pedidos:%v] [cantidad:%d]", pedidos, len(pedidos))
		c.JSON(http.StatusOK, pedidos)
	}
}

func (handler *PedidoHandler) ObtenerPedidoPorId(c *gin.Context) {
	id := c.Param("id")
	pedido := handler.pedidoService.ObtenerPedidoPorId(id)
	if pedido == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "pedido no encontrado"})
	} else {
		c.JSON(http.StatusOK, gin.H{"pedido": pedido})
	}
}
func (handler *PedidoHandler) InsertarPedido(c *gin.Context) {
	var pedido dto.Pedido
	if err := c.ShouldBindJSON(&pedido); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} else {
		resultado := handler.pedidoService.InsertarPedido(&pedido)
		if resultado != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": resultado.Error()})
		} else {
			c.JSON(http.StatusCreated, gin.H{"estado": "Creado correctamente"})
		}

	}

}
func (handler *PedidoHandler) EliminarPedido(c *gin.Context) {
	id := c.Param("id")
	if err := handler.pedidoService.EliminarPedido(id); err == nil {
		c.JSON(http.StatusOK, gin.H{"mensaje": "Pedido eliminado correctamente"})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No se pudo eliminar el pedido"})
	}
}
func (handler *PedidoHandler) AceptarPedido(c *gin.Context) {
	id := c.Param("id")
	pedido := handler.pedidoService.ObtenerPedidoPorId(id)

	if err := c.ShouldBindJSON(&pedido); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} else {
		resultado := handler.pedidoService.AceptarPedido(pedido)
		if resultado != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": resultado.Error()})
		} else {
			c.JSON(http.StatusAccepted, gin.H{"resultado": "Pedido aceptado correctamente"})
		}
	}

}
func (handler *PedidoHandler) ObtenerCantidadPedidosPorEstado(c *gin.Context) {
	estado := c.Param("estado")
	cantidad, err := handler.pedidoService.ObtenerCantidadPedidosPorEstado(estado)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	} else {
		c.JSON(http.StatusOK, cantidad)
	}
}
func (handler *PedidoHandler) ObtenerPedidosPorEstado(c *gin.Context) {
	estado := c.Param("estado")
	pedidos, err := handler.pedidoService.ObtenerPedidosPorEstado(estado)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return

	} else {
		c.JSON(http.StatusOK, pedidos)
	}
}
