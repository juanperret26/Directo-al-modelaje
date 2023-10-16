// Crear interface, structura y new EnvioService
package services

import (
	"github.com/juanperret/Directo-al-modelaje/dto"
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
	envioRepository repositories.EnvioRepositoryInterface
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
func (service *envioService) InsertarEnvio(envio *dto.Envio, pedidos []*dto.Pedido) bool {
	pesoTotal := 0
	for i, pedido := range pedidos {
		pesoPorPedido := pedido.PedidoProductos[i].Producto.Peso * pedido.PedidoProductos[i].Cantidad
		pesoTotal += pesoPorPedido
		if pedido.Estado == "Aceptado" {

		}
	}

	service.envioRepository.InsertarEnvio(envio.GetModel())
	return true
}
func (service *envioService) EliminarEnvio(id string) bool {
	service.envioRepository.EliminarEnvio(utils.GetObjectIDFromStringID(id))
	return true
}
func (service *envioService) ActualizarEnvio(envio *dto.Envio) bool {
	service.envioRepository.ActualizarEnvio(envio.GetModel())
	return true
}
