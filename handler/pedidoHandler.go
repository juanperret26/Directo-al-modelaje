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
	if err := c.ShouldBindJSON(&pedido); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	resultado := handler.pedidoService.InsertarPedido(&pedido)
	c.JSON(http.StatusCreated, resultado)
}
func (handler *PedidoHandler) EliminarPedido(c *gin.Context) {
	id := c.Param("id")
	if handler.pedidoService.EliminarPedido(id) {
		c.JSON(http.StatusOK, gin.H{"mensaje": "Pedido eliminado correctamente"})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No se pudo eliminar el pedido"})
	}
}
