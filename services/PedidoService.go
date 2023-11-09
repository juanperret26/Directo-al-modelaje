// Crear interface, structura y new PedidoService
package services

import (
	"errors"

	"github.com/juanperret/Directo-al-modelaje/dto"
	"github.com/juanperret/Directo-al-modelaje/repositories"
)

type PedidoInterface interface {
	//Metodos para implementar en el handler
	ObtenerPedidos() []*dto.Pedido
	ObtenerPedidoPorId(id string) *dto.Pedido
	hayStockDisponiblePedido(pedido *dto.Pedido) bool
	InsertarPedido(pedido *dto.Pedido) error
	EliminarPedido(id string) error
	AceptarPedido(pedido *dto.Pedido) error
	ActualizarPedido(pedido *dto.Pedido) error
	ObtenerCantidadPedidosPorEstado(estado string) ([]dto.Estado, error)
}
type pedidoService struct {
	pedidoRepository   repositories.PedidoRepositoryInterface
	productoRepository repositories.ProductoRepositoryInterface
}

func NewPedidoService(pedidoRepository repositories.PedidoRepositoryInterface, productoRepository repositories.ProductoRepositoryInterface) *pedidoService {
	return &pedidoService{
		pedidoRepository:   pedidoRepository,
		productoRepository: productoRepository,
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

func (service *pedidoService) InsertarPedido(pedido *dto.Pedido) error {

	_, err := service.pedidoRepository.InsertarPedido(pedido.GetModel())
	return err
}
func (service *pedidoService) AceptarPedido(pedidoPorAceptar *dto.Pedido) error {
	//Primero buscamos el pedido a aceptar
	pedido, err := service.pedidoRepository.ObtenerPedidoPorId(pedidoPorAceptar.Id)

	if err != nil {
		return err
	} else {
		//Verifica que haya stock disponible para aceptar el pedido
		if !service.hayStockDisponiblePedido(pedidoPorAceptar) {
			return errors.New("no hay stock disponible para aceptar el pedido")
		} else {

			//Cambia el estado del pedido a Aceptado, si es que no estaba ya en ese estado
			if pedido.Estado != "Aceptado" {
				pedido.Estado = "Aceptado"
			}

			//Actualiza el pedido en la base de datos
			_, err := service.pedidoRepository.ActualizarPedido(pedido)
			return err
		}
	}
}

func (service *pedidoService) hayStockDisponiblePedido(pedido *dto.Pedido) bool {
	//Busco los productos del pedido
	productosPedido := pedido.PedidoProductos

	//Recorro los productos del pedido
	for _, productoPedido := range productosPedido {
		//Armo un objeto producto con el ID para buscar en la base de datos

		//Busco el producto en la base de datos
		producto, err := service.productoRepository.ObtenerProductoPorId(productoPedido.CodigoProducto)

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
func (service *pedidoService) EliminarPedido(id string) error {
	_, err := service.pedidoRepository.EliminarPedido(id)
	return err
}
func (service *pedidoService) ActualizarPedido(pedido *dto.Pedido) error {
	_, err := service.pedidoRepository.ActualizarPedido(pedido.GetModel())
	return err
}

// Obtener la cantidad de pedidos por estado
func (service *pedidoService) ObtenerCantidadPedidosPorEstado(estado string) ([]dto.Estado, error) {
	//Por cada estado posible de pedidos, obtengo la cantidad de pedidos en ese estado
	var cantidadPedidos []int
	var listaEstados []dto.Estado
	switch estado {
	case "Pendiente":
		cantidadPedidosPendientes, err := service.pedidoRepository.ObtenerCantidadPedidosPorEstado(estado)
		if err != nil {
			return nil, err
		}
		cantidadPedidos = append(cantidadPedidos, cantidadPedidosPendientes)
		listaEstados = append(listaEstados, dto.Estado{Estado: "Pendiente", Cantidad: cantidadPedidosPendientes})
	case "Aceptado":
		cantidadPedidosAceptados, err := service.pedidoRepository.ObtenerCantidadPedidosPorEstado(estado)
		if err != nil {
			return nil, err
		}
		cantidadPedidos = append(cantidadPedidos, cantidadPedidosAceptados)
		listaEstados = append(listaEstados, dto.Estado{Estado: "Aceptado", Cantidad: cantidadPedidosAceptados})
	case "Cancelado":
		cantidadPedidosCancelados, err := service.pedidoRepository.ObtenerCantidadPedidosPorEstado(estado)
		if err != nil {
			return nil, err
		}
		cantidadPedidos = append(cantidadPedidos, cantidadPedidosCancelados)
		listaEstados = append(listaEstados, dto.Estado{Estado: "Cancelado", Cantidad: cantidadPedidosCancelados})
	case "ParaEnviar":
		cantidadPedidosParaEnviar, err := service.pedidoRepository.ObtenerCantidadPedidosPorEstado(estado)
		if err != nil {
			return nil, err
		}
		cantidadPedidos = append(cantidadPedidos, cantidadPedidosParaEnviar)
		listaEstados = append(listaEstados, dto.Estado{Estado: "ParaEnviar", Cantidad: cantidadPedidosParaEnviar})
	case "Enviado":
		cantidadPedidosEnviados, err := service.pedidoRepository.ObtenerCantidadPedidosPorEstado(estado)
		if err != nil {
			return nil, err
		}
		cantidadPedidos = append(cantidadPedidos, cantidadPedidosEnviados)
		listaEstados = append(listaEstados, dto.Estado{Estado: "Enviado", Cantidad: cantidadPedidosEnviados})
	default:
		return nil, errors.New("El estado ingresado no es valido")
	}

	return listaEstados, nil
}
