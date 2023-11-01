// Crear interface, structura y new EnvioService
package services

import (
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
	InsertarEnvio(envio *dto.Envio) bool
	EliminarEnvio(id string) bool
	ActualizarEnvio(envio *dto.Envio) bool
	ValidacionViaje(envio *dto.Envio, inicio bool, parada dto.Paradas)
}
type envioService struct {
	envioRepository    repositories.EnvioRepositoryInterface
	pedidoRepository   repositories.PedidoRepository
	productoRepository repositories.ProductoRepository
	camionRepository   repositories.CamionRepository
}

func NewEnvioService(envioRepository repositories.EnvioRepositoryInterface) *envioService {
	return &envioService{
		envioRepository: envioRepository,
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

func (service *envioService) InsertarEnvio(envio *dto.Envio) bool {
	var pesoTotal float64 = 0.0
	var resultado = false
	var camion, _ = service.camionRepository.ObtenerCamionPorPatente(envio.PatenteCamion)
	var pedidos = envio.Pedido

	for _, idPedido := range pedidos {

		// Buscar el pedido correspondiente al ID
		var pedido model.Pedido
		pedido, err := service.pedidoRepository.ObtenerPedidoPorId(idPedido)
		if pedido.Estado == "Aceptado" {
			for _, productoPedido := range pedido.PedidoProductos {
				// Buscar el producto correspondiente al codigo
				producto, err := service.productoRepository.ObtenerProductoPorId(productoPedido.CodigoProducto)
				PesoPedido := producto.Peso_unitario * float64(productoPedido.Cantidad)
				pesoTotal += PesoPedido

				if err != nil {
					log.Printf("[service:ProductoService][method:ObtenerProductoPorId][reason:NOT_FOUND][id:%d]", productoPedido.CodigoProducto)
				}

			}
		}

		if err != nil {
			// Manejar el error
			log.Printf("[service:PedidoService][method:ObtenerPedidosPorId][reason:NOT_FOUND][id:%d]", idPedido)
		}

		if pesoTotal <= camion.Peso_maximo {
			envio.Estado = "A DESPACHAR"
			envio.PatenteCamion = camion.Patente
			service.envioRepository.InsertarEnvio(envio.GetModel())
			pedido.Estado = "PARA ENVIAR"
			resultado = true
			service.pedidoRepository.ActualizarPedido(pedido)
			service.envioRepository.ActualizarEnvio(envio.GetModel())
		}
		if pesoTotal > camion.Peso_maximo {
			log.Printf("No se puede cargar el envio en el camion")
		}
	}
	return resultado
}
func (service *envioService) EliminarEnvio(id string) bool {
	service.envioRepository.EliminarEnvio(utils.GetObjectIDFromStringID(id))
	return true
}
func (service *envioService) ActualizarEnvio(envio *dto.Envio) bool {
	service.envioRepository.ActualizarEnvio(envio.GetModel())
	return true
}
func (service *envioService) ValidacionViaje(envio *dto.Envio, inicio bool, parada dto.Paradas) {
	if inicio {
		envio.Estado = "En ruta"
		service.envioRepository.ActualizarEnvio(envio.GetModel())
		envio.Paradas = append(envio.Paradas, parada)
		if parada.Ciudad == envio.Destino {
			envio.Estado = "Despachado"
			service.envioRepository.ActualizarEnvio(envio.GetModel())
			for _, idPedido := range envio.Pedido {
				var pedido model.Pedido
				pedido, err := service.pedidoRepository.ObtenerPedidoPorId(idPedido)
				if err != nil {
					log.Printf("[service:PedidoService][method:ObtenerPedidosPorId][reason:NOT_FOUND][id:%d]", idPedido)
				}

				for _, productoPedido := range pedido.PedidoProductos {
					// Buscar el producto correspondiente al codigo
					producto, err := service.productoRepository.ObtenerProductoPorId(productoPedido.CodigoProducto)
					if err != nil {
						log.Printf("[service:ProductoService][method:ObtenerProductoPorId][reason:NOT_FOUND][id:%d]", productoPedido.CodigoProducto)
					}
					producto.Stock -= productoPedido.Cantidad
					service.productoRepository.ActualizarProducto(producto)

				}
				pedido.Estado = "Enviado"
				service.pedidoRepository.ActualizarPedido(pedido)
			}

		}
	}
}
