// Crear interface, structura y new EnvioService
package services

import (
	"errors"
	"log"
	"time"

	"github.com/juanperret/Directo-al-modelaje/dto"
	"github.com/juanperret/Directo-al-modelaje/model"
	"github.com/juanperret/Directo-al-modelaje/repositories"
	"github.com/juanperret/Directo-al-modelaje/utils"
)

type EnvioInterface interface {
	//Metodos para implementar en el handler
	ObtenerEnvios() []*dto.Envio
	ObtenerEnvioPorId(id string) *dto.Envio
	ObtenerPedidosFiltrados(codigoEnvio string, estado string, fechaInicio time.Time, fechaFinal time.Time) ([]*dto.Pedido, error)
	ObtenerEnviosPorParametros(patente string, estado string, ultimaParada string, fechaCreacionDesde time.Time, fechaCreacionHasta time.Time) ([]*dto.Envio, error)
	InsertarEnvio(envio *dto.Envio) error
	EliminarEnvio(id string) error
	ActualizarEnvio(envio *dto.Envio) error
	IniciarViaje(envio *dto.Envio) error
	ObtenerCantidadEnviosPorEstado(estado string) ([]dto.Estado, error)
	AgregarParada(envio *dto.Envio) (bool, error)
	ObtnerBeneficiosEntreFecha(fechaInicio time.Time, fechaFinal time.Time) (int, error)
}
type envioService struct {
	envioRepository    repositories.EnvioRepositoryInterface
	pedidoRepository   repositories.PedidoRepositoryInterface
	productoRepository repositories.ProductoRepositoryInterface
	camionRepository   repositories.CamionRepositoryInterface
}

func NewEnvioService(envioRepository repositories.EnvioRepositoryInterface, camionRepository repositories.CamionRepositoryInterface, pedidoRepository repositories.PedidoRepositoryInterface, productoRepository repositories.ProductoRepositoryInterface) *envioService {
	return &envioService{
		envioRepository:    envioRepository,
		camionRepository:   camionRepository,
		pedidoRepository:   pedidoRepository,
		productoRepository: productoRepository,
	}
}
func (service *envioService) ObtenerEnvios() []*dto.Envio {
	envioDB, _ := service.envioRepository.ObtenerEnvios()
	var envios []*dto.Envio
	for _, envioDB := range envioDB {
		envio := dto.NewEnvio(envioDB)
		envios = append(envios, envio)
	}
	return envios
}
func (service *envioService) ObtenerEnviosPorParametros(patente string, estado string, ultimaParada string, fechaCreacionDesde time.Time, fechaCreacionHasta time.Time) ([]*dto.Envio, error) {

	enviosDB, err := service.envioRepository.ObtenerEnviosPorParametros(patente, estado, ultimaParada, fechaCreacionDesde, fechaCreacionHasta)
	if err != nil {
		return nil, err
	}

	var envios []*dto.Envio
	for _, envioDB := range enviosDB {
		envio := dto.NewEnvio(envioDB)
		envios = append(envios, envio)
	}
	return envios, nil
}
func (service *envioService) ObtenerEnvioPorId(id string) *dto.Envio {
	envioDB, _ := service.envioRepository.ObtenerEnvioPorId(id)
	envio := dto.NewEnvio(envioDB)
	return envio
}
func (service *envioService) ObtenerPedidosFiltrados(codigoEnvio string, estado string, fechaInicio time.Time, fechaFinal time.Time) ([]*dto.Pedido, error) {

	pedidoDB, _ := service.envioRepository.ObtenerPedidosFiltrados(codigoEnvio, estado, fechaInicio, fechaFinal)
	var pedidos []*dto.Pedido
	for _, pedidoDB := range pedidoDB {
		pedido := dto.NewPedido(pedidoDB)
		pedidos = append(pedidos, pedido)
	}
	return pedidos, nil
}

func (service *envioService) InsertarEnvio(envio *dto.Envio) error {
	var pesoTotal float64 = 0.0
	var camion, err = service.camionRepository.ObtenerCamionPorPatente(envio.PatenteCamion)
	log.Printf("camion: %v", err)

	for _, idPedido := range envio.Pedido {

		// Buscar el pedido correspondiente al ID
		pedido, err := service.pedidoRepository.ObtenerPedidoPorId(idPedido)
		if err != nil {
			// Manejar el error
			log.Printf("[service:PedidoService][method:ObtenerPedidosPorId][reason:NOT_FOUND][id:%d]", idPedido)
		}
		if pedido.Estado == "Aceptado" {
			for _, productoPedido := range pedido.PedidoProductos {
				// Buscar el producto correspondiente al codigo
				producto, err := service.productoRepository.ObtenerProductoPorId(productoPedido.CodigoProducto)
				PesoPedido := producto.Peso_unitario * float64(productoPedido.Cantidad)
				pesoTotal += PesoPedido

				if err != nil {
					log.Printf("[service:ProductoService][method:ObtenerProductoPorId][reason:NOT_FOUND][id:%d]", productoPedido.CodigoProducto)
				}
				if pesoTotal <= camion.Peso_maximo {
					envio.Estado = "A despachar"
					envio.PatenteCamion = camion.Patente
					service.envioRepository.InsertarEnvio(envio.GetModel())
					pedido.Estado = "Para enviar"
					service.pedidoRepository.ActualizarPedido(pedido)
					service.envioRepository.ActualizarEnvio(envio.GetModel())
				}
				if pesoTotal > camion.Peso_maximo {
					log.Printf("No se puede cargar el envio en el camion")
				}

			}
		} else {
			log.Printf("El pedido no esta aceptado")
		}
	}
	return err
}
func (service *envioService) EliminarEnvio(id string) error {

	_, err := service.envioRepository.EliminarEnvio(utils.GetObjectIDFromStringID(id))
	return err
}
func (service *envioService) ActualizarEnvio(envio *dto.Envio) error {

	err := service.envioRepository.ActualizarEnvio(envio.GetModel())
	return err
}
func (service *envioService) IniciarViaje(envio *dto.Envio) error {
	envioABuscar, err := service.envioRepository.ObtenerEnvioPorId(envio.Id)
	if err != nil {
		return err
	}
	envioABuscar.Estado = "En ruta"

	//Calcular el stock de los productos y actualizarlo
	for _, idPedido := range envioABuscar.Pedido {
		var pedido model.Pedido

		pedido.Estado = "Para enviar"
		service.pedidoRepository.ActualizarPedido(pedido)
		pedido, err := service.pedidoRepository.ObtenerPedidoPorId(idPedido)

		if err != nil {
			log.Printf("[service:PedidoService][method:ObtenerPedidosPorId][reason:NOT_FOUND][id:%d]", idPedido)
		}
		service.DescontarStock(pedido)
		service.pedidoRepository.ActualizarPedido(pedido)
	}
	service.envioRepository.ActualizarEnvio(envioABuscar)
	return err
}

func (service *envioService) DescontarStock(pedido model.Pedido) {
	for _, productoPedido := range pedido.PedidoProductos {
		// Buscar el producto correspondiente al codigo
		producto, err := service.productoRepository.ObtenerProductoPorId(productoPedido.CodigoProducto)
		if err != nil {
			log.Printf("[service:ProductoService][method:ObtenerProductoPorId][reason:NOT_FOUND][id:%d]", productoPedido.CodigoProducto)
		}
		producto.Stock -= productoPedido.Cantidad

		service.productoRepository.ActualizarProducto(producto)
	}
}

func (service *envioService) AgregarParada(envio *dto.Envio) (bool, error) {
	//En teoria, recibimos un envio que tiene solamente id y la nueva parada
	//Primero buscamos el envio por id
	envioDB, err := service.envioRepository.ObtenerEnvioPorId(envio.Id)
	camion, err := service.camionRepository.ObtenerCamionPorPatente(envioDB.PatenteCamion)
	if err != nil {
		return false, err
	}

	//Validamos que el envio esté en estado EnRuta
	if envioDB.Estado != "En ruta" {
		return false, errors.New("el envio no esta en ruta")
	}

	//Agregamos la nueva parada al envio
	envioDB.Paradas = append(envioDB.Paradas, envio.Paradas[0].GetModel())
	envioDB.Costo = envioDB.Costo + envio.Paradas[0].Kilometros*camion.Costo_km
	if envioDB.Destino.Nombre_ciudad == envio.Paradas[0].Ciudad {
		envioDB.Estado = "Despachado"
		for _, idPedido := range envioDB.Pedido {
			pedido, err := service.pedidoRepository.ObtenerPedidoPorId(idPedido)
			if err != nil {
				log.Printf("[service:PedidoService][method:ObtenerPedidosPorId][reason:NOT_FOUND][id:%d]", idPedido)
			}
			pedido.Estado = "Enviado"
			service.pedidoRepository.ActualizarPedido(pedido)
		}

	}
	//Actualizamos el envio en la base de datos, que ahora tiene la nueva parada
	return true, service.envioRepository.ActualizarEnvio(envioDB)
}

func (service *envioService) ObtenerCantidadEnviosPorEstado(estado string) ([]dto.Estado, error) {
	var cantidadEnvios []int
	var listaEstados []dto.Estado
	switch estado {
	case "A despachar":
		cantidadEnviosADespachar, err := service.envioRepository.ObtenerCantidadEnviosPorEstado(estado)
		if err != nil {
			return nil, err
		}
		cantidadEnvios = append(cantidadEnvios, cantidadEnviosADespachar)
		listaEstados = append(listaEstados, dto.Estado{Estado: "A despachar", Cantidad: cantidadEnviosADespachar})
	case "En ruta":
		cantidadEnviosEnRuta, err := service.envioRepository.ObtenerCantidadEnviosPorEstado(estado)
		if err != nil {
			return nil, err
		}
		cantidadEnvios = append(cantidadEnvios, cantidadEnviosEnRuta)
		listaEstados = append(listaEstados, dto.Estado{Estado: "En ruta", Cantidad: cantidadEnviosEnRuta})
	case "Despachado":
		cantidadEnviosDespachados, err := service.envioRepository.ObtenerCantidadEnviosPorEstado(estado)
		if err != nil {
			return nil, err
		}
		cantidadEnvios = append(cantidadEnvios, cantidadEnviosDespachados)
		listaEstados = append(listaEstados, dto.Estado{Estado: "Despachado", Cantidad: cantidadEnviosDespachados})
	default:
		log.Printf("El estado ingresado no es valido")
	}
	return listaEstados, nil
}
func (service *envioService) ObtnerBeneficiosEntreFecha(fechaInicio time.Time, fechaFinal time.Time) (int, error) {
	var beneficio int = 0
	envios, err := service.envioRepository.ObtenerEnviosPorParametros("", "", "", fechaInicio, fechaFinal)
	if err != nil {
		return 0, err
	}
	for _, envio := range envios {
		for _, pedido := range envio.Pedido {
			pedidoDB, err := service.pedidoRepository.ObtenerPedidoPorId(pedido)
			if err != nil {
				return 0, err
			}
			for _, productoPedido := range pedidoDB.PedidoProductos {
				// Buscar el producto correspondiente al codigo
				producto, err := service.productoRepository.ObtenerProductoPorId(productoPedido.CodigoProducto)
				if err != nil {
					log.Printf("[service:ProductoService][method:ObtenerProductoPorId][reason:NOT_FOUND][id:%d]", productoPedido.CodigoProducto)
				}
				beneficio += int(producto.Precio) * productoPedido.Cantidad
			}
		}
	}
	return beneficio, nil
}
