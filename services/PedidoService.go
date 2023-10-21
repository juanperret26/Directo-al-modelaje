// Crear interface, structura y new PedidoService
package services

import (
	"time"

	"github.com/juanperret/Directo-al-modelaje/dto"
	"github.com/juanperret/Directo-al-modelaje/repositories"
	"github.com/juanperret/Directo-al-modelaje/utils"
	
)

type PedidoInterface interface {
	//Metodos para implementar en el handler
	ObtenerPedidos() []*dto.Pedido
	ObtenerPedidoPorId(id string) *dto.Pedido
	ObtenerPedidosFiltrados(codigoEnvio string, estado string, fecha time.Time) []*dto.Pedido
	InsertarPedido(pedido *dto.Pedido) bool
	EliminarPedido(id string) bool
	ActualizarPedido(pedido *dto.Pedido) bool
}
type pedidoService struct {
	pedidoRepository repositories.PedidoRepositoryInterface
}

func NewPedidoService(pedidoRepository repositories.PedidoRepositoryInterface) *pedidoService {
	return &pedidoService{
		pedidoRepository: pedidoRepository,
	}
}
func (service *pedidoService) ObtenerPedidos() []*dto.Pedido {
	pedidoDB, _ := service.pedidoRepository.ObtenerPedidos()
	var pedidos []*dto.Pedido
	for _, pedidoDB := range pedidoDB {
		pedido := dto.NewPedido(pedidoDB)
		pedidos = append(pedidos, pedido)
	}
	return pedidos
}
func (service *pedidoService) ObtenerPedidoPorId(id string) *dto.Pedido {
	pedidoDB, _ := service.pedidoRepository.ObtenerPedidoPorId(id)
	pedido := dto.NewPedido(pedidoDB)
	return pedido
}
func (service *pedidoService)ObtenerPedidosFiltrados(codigoEnvio string, estado string, fecha time.Time) []*dto.Pedido {
	pedidoDB, _ := service.pedidoRepository.ObtenerPedidosFiltrados(codigoEnvio, estado, fecha)

	var pedidos []*dto.Pedido
	for _, pedidoDB := range pedidoDB {
		pedido := dto.NewPedido(pedidoDB)
		pedidos = append(pedidos, pedido)
	}
	return pedidos
}
func (service *pedidoService) InsertarPedido(pedido *dto.Pedido, producto *dto.Producto, cantidad float64) bool {
	service.pedidoRepository.InsertarPedido(pedido.GetModel())
	pedidoProducto := dto.NewPedidoProducto(producto.GetModel())
	pedidoProducto.Cantidad = cantidad
	if pedidoProducto.Cantidad <= producto.Stock {
		pedido.PedidoProductos = append(pedido.PedidoProductos, *pedidoProducto)
		pedido.Estado = "Aceptado"
		service.pedidoRepository.ActualizarPedido(pedido.GetModel())
	}
	return true
}
func (service *pedidoService) EliminarPedido(pedido *dto.Pedido, id string)bool{
	service.pedidoRepository.EliminarPedido(utils.GetObjectIDFromStringID(id))
	return true
}
func (service *pedidoService) ActualizarPedido(pedido *dto.Pedido) bool {
	service.pedidoRepository.ActualizarPedido(pedido.GetModel())
	return true
}
