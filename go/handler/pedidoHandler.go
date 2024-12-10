// Crear struct, new objeto y metodos
package handler

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/juanperret26/Directo-al-modelaje/go/dto"
	"github.com/juanperret26/Directo-al-modelaje/go/services"
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
		c.JSON(http.StatusOK, pedido)
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
			c.JSON(http.StatusCreated, gin.H{"mensaje": "Pedido creado correctamente"})
		}

	}

}
func (handler *PedidoHandler) EliminarPedido(c *gin.Context) {
	id := c.Param("id")
	if err := handler.pedidoService.EliminarPedido(id); err == nil {
		c.JSON(http.StatusOK, gin.H{"mensaje": "Pedido cancelado correctamente"})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No se pudo cancelar el pedido"})
	}
}


func (handler *PedidoHandler) AceptarPedido(c *gin.Context) {
    // Obtener el ID del parámetro de la URL
    id := c.Param("id")
    if id == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "El ID del pedido es requerido"})
        return
    }

    peditoDto :=dto.Pedido{Id : id}	 
    if err := handler.pedidoService.AceptarPedido(&peditoDto); err != nil {
        if err.Error() == "pedido no encontrado" {
            c.JSON(http.StatusNotFound, gin.H{"error": "Pedido no encontrado"})
        } else {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        }
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Pedido actualizado a 'Aceptado'"})
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
func (handler *EnvioHandler) ObtenerPedidosFiltrados(c *gin.Context) {
	codigoEnvio := c.Param("codigoEnvio")
	estado := c.Param("estado")
	fechaInicioStr := c.Param("fechaInicio")
	fechaFinalStr := c.Param("fechaFinal")

	// Convertir strings a time.Time
	fechaInicio, errInicio := time.Parse("2006-01-02", fechaInicioStr)
	if errInicio != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Formato de fechaInicio incorrecto"})
		return
	}

	fechaFinal, errFinal := time.Parse("2006-01-02", fechaFinalStr)
	if errFinal != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Formato de fechaFinal incorrecto"})
		return
	}
	filtro := dto.Filtro{
		CodigoEnvio:        codigoEnvio,
		EstadoPedido:       estado,
		FechaCreacionDesde: fechaInicio,
		FechaCreacionHasta: fechaFinal,
	}
	// Manejo de errores
	pedidos, err := handler.envioService.ObtenerPedidosFiltro(&filtro)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	} else {
		// Registro de información
		log.Printf("Se obtuvieron pedidos filtrados para codigoEnvio %s, estado %s, fechaInicio %s, fechaFinal %s", codigoEnvio, estado, fechaInicio, fechaFinal)
		c.JSON(http.StatusOK, pedidos)
	}
}
