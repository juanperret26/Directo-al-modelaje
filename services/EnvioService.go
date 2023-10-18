// Crear interface, structura y new EnvioService
package services

import (
	"log"

	"github.com/juanperret/Directo-al-modelaje/dto"
	"github.com/juanperret/Directo-al-modelaje/model"
	"github.com/juanperret/Directo-al-modelaje/repositories"
	"github.com/juanperret/Directo-al-modelaje/utils"
)

type EnvioInterface interface {
	//Metodos para implementar en el handler
	ObtenerEnvios() []*dto.Envio
	ObtenerEnvioPorId(id string) *dto.Envio
	InsertarEnvio(envio *dto.Envio) bool
	EliminarEnvio(id string) bool
	ActualizarEnvio(envio *dto.Envio) bool
}
type envioService struct {
	envioRepository    repositories.EnvioRepositoryInterface
	pedidoRepository   repositories.PedidoRepository
	productoRepository repositories.ProductoRepository
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
func (service *envioService) ObtenerEnvioPorId(id string) *dto.Envio {
	envioDB, _ := service.envioRepository.ObtenerEnvioPorId(id)
	envio := dto.NewEnvio(envioDB)
	return envio
}
func (service *envioService) InsertarEnvio(envio *dto.Envio, pedidos []string, camion *dto.Camion) bool {
	var pesoTotal float64 = 0.0
	var resultado = false
	for _, idPedido := range pedidos {

		// Buscar el pedido correspondiente al ID
		var pedido model.Pedido
		pedido, err := service.pedidoRepository.ObtenerPedidoPorId(idPedido)
		if pedido.Estado == "Aceptado" {
			for _, productoPedido := range pedido.PedidoProductos {
				// Buscar el producto correspondiente al codigo
				producto, err := service.productoRepository.ObtenerProductoPorId(productoPedido.CodigoProducto)
				PesoPedido := producto.Peso_unitario * productoPedido.Cantidad
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
			service.envioRepository.InsertarEnvio(envio.GetModel())
			pedido.Estado = "PARA ENVIAR"
			resultado = true
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
