// Crear interface, structura y new PedidoService
package services

import (
	"errors"

	"github.com/juanperret/Directo-al-modelaje/dto"
	"github.com/juanperret/Directo-al-modelaje/model"
	"github.com/juanperret/Directo-al-modelaje/repositories"
)

type PedidoInterface interface {
	//Metodos para implementar en el handler
	ObtenerPedidos() []*dto.Pedido
	ObtenerPedidoPorId(id string) *dto.Pedido
	hayStockDisponiblePedido(pedido *dto.Pedido) bool
	InsertarPedido(pedido *dto.Pedido) bool
	EliminarPedido(id string) bool
	ActualizarPedido(pedido *dto.Pedido) bool
}
type pedidoService struct {
	pedidoRepository   repositories.PedidoRepositoryInterface
	productoRepository repositories.ProductoRepositoryInterface
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

func (service *pedidoService) InsertarPedido(pedido *dto.Pedido) bool {
	service.pedidoRepository.InsertarPedido(pedido.GetModel())

	return true
}
func (service *pedidoService) AceptarPedido(pedidoPorAceptar *dto.Pedido) error {
	//Primero buscamos el pedido a aceptar
	pedido, err := service.pedidoRepository.ObtenerPedidoPorId(pedidoPorAceptar.Id)

	if err != nil {
		return err
	}

	//Verifica que haya stock disponible para aceptar el pedido
	if !service.hayStockDisponiblePedido(pedido) {
		return errors.New("no hay stock disponible para aceptar el pedido")
	}

	//Cambia el estado del pedido a Aceptado, si es que no estaba ya en ese estado
	if pedido.Estado != model.Aceptado {
		pedido.Estado = model.Aceptado
	}

	//Actualiza el pedido en la base de datos
	return service.pedidoRepository.ActualizarPedido(*pedido)
}

func (service *pedidoService) hayStockDisponiblePedido(pedido *dto.Pedido) bool {
	//Busco los productos del pedido
	productosPedido := pedido.PedidoProductos

	//Recorro los productos del pedido
	for _, productoPedido := range productosPedido {
		//Armo un objeto producto con el ID para buscar en la base de datos
		productoParaBuscar := dto.Producto{CodigoProducto: productoPedido.CodigoProducto}

		//Busco el producto en la base de datos
		producto, err := service.productoRepository.ObtenerProductoPorId(productoParaBuscar.Id)

		if err != nil {
			return false
		}

		//Verifico que haya stock disponible para el producto
		if productoPedido.Cantidad > producto.Stock {
			return false
		}
	}

	//Si finalice el bucle, es porque hay stock de todos los productos
	return true
}
func (service *pedidoService) EliminarPedido(id string) bool {
	service.pedidoRepository.EliminarPedido(id)
	return true
}
func (service *pedidoService) ActualizarPedido(pedido *dto.Pedido) bool {
	service.pedidoRepository.ActualizarPedido(pedido.GetModel())
	return true
}
