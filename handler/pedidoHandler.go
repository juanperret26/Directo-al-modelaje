// Crear struct, new objeto y metodos
package handler

import (
	"time"

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
func (handler *PedidoHandler) ObtenerPedidos() []*dto.Pedido {
	return handler.pedidoService.ObtenerPedidos()
}
func (handler *PedidoHandler) ObtenerPedidosFiltrados(codigoEnvio string, estado string, fecha time.Time) []*dto.Pedido {
	return handler.pedidoService.ObtenerPedidosFiltrados(codigoEnvio, estado, fecha)
}
func (handler *PedidoHandler) ObtenerPedidoPorId(id string) *dto.Pedido {
	return handler.pedidoService.ObtenerPedidoPorId(id)
}
func (handler *PedidoHandler) InsertarPedido(pedido *dto.Pedido) bool {
	return handler.pedidoService.InsertarPedido(pedido)
}
func (handler *PedidoHandler) EliminarPedido(id string) bool {
	return handler.pedidoService.EliminarPedido(id)
}
