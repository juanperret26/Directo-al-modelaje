// Crear interface, structura y new EnvioService
package services

import (
	"github.com/juanperret/Directo-al-modelaje/dto"
	"github.com/juanperret/Directo-al-modelaje/repositories"
	"github.com/juanperret/Directo-al-modelaje/utils"
)

type EnvioInterface interface {
	//Metodos para implementar en el handler
	GetEnvios() []*dto.Envio
	GetEnvio(id string) *dto.Envio
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
func (service *envioService) GetEnvios() []*dto.Envio {
	envioDB, _ := service.envioRepository.GetEnvios()
	var envios []*dto.Envio
	for _, envioDB := range envioDB {
		envio := dto.NewEnvio(envioDB)
		envios = append(envios, envio)
	}
	return envios
}
func (service *envioService) GetEnvio(id string) *dto.Envio {
	envioDB, _ := service.envioRepository.GetEnvio(id)
	envio := dto.NewEnvio(envioDB)
	return envio
}
func (service *envioService) InsertarEnvio(envio *dto.Envio) bool {
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
