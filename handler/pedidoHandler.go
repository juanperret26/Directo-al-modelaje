// Crear struct, new objeto y metodos
package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/juanperret/Directo-al-modelaje/dto"
	"github.com/juanperret/Directo-al-modelaje/services"
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
	log.Printf("[handler:PedidoHandler] [method:ObtenerPedidos] [pedidos:%v] [cantidad:%d]", pedidos, len(pedidos))
	c.JSON(http.StatusOK, pedidos)

}

func (handler *PedidoHandler) ObtenerPedidoPorId(c *gin.Context) {
	id := c.Param("id")
	pedido := handler.pedidoService.ObtenerPedidoPorId(id)
	c.JSON(http.StatusOK, pedido)
}
func (handler *PedidoHandler) InsertarPedido(c *gin.Context) {
	var pedido dto.Pedido
	err := c.ShouldBindJSON(&pedido)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		resultado := handler.pedidoService.InsertarPedido(&pedido)
		if resultado != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": resultado.Error()})
		} else {
			c.JSON(http.StatusCreated, gin.H{"error": resultado.Error()})
		}
	}
}
func (handler *PedidoHandler) EliminarPedido(c *gin.Context) {
	id := c.Param("id")
	err := handler.pedidoService.EliminarPedido(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"mensaje": "Pedido eliminado correctamente"})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}
func (handler *PedidoHandler) AceptarPedido(c *gin.Context) {
	id := c.Param("id")
	pedido := handler.pedidoService.ObtenerPedidoPorId(id)
	err := c.ShouldBindJSON(&pedido)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		resultado := handler.pedidoService.AceptarPedido(pedido)
		if resultado != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": resultado.Error()})
		} else {
			c.JSON(http.StatusCreated, gin.H{"error": resultado.Error()})
		}
	}
}
func (handler *PedidoHandler) ObtenerCantidadPedidosPorEstado(c *gin.Context) {
	estado := c.Param("estado")
	cantidad, err := handler.pedidoService.ObtenerCantidadPedidosPorEstado(estado)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, cantidad)
	}
}
