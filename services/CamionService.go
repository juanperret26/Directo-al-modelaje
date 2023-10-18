// Crear interface, structura y new CamionService
package services

import (
	"log"

	"github.com/juanperret/Directo-al-modelaje/dto"
	"github.com/juanperret/Directo-al-modelaje/repositories"
	"github.com/juanperret/Directo-al-modelaje/utils"
)

type CamionInterface interface {
	//Metodos para implementar en el handler
	ObtenerCamiones() []*dto.Camion
	ObtenerCamionPorId(id string) *dto.Camion
	InsertarCamion(camion *dto.Camion) bool
	EliminarCamion(id string) bool
	ActualizarCamion(camion *dto.Camion) bool
}
type camionService struct {
	camionRepository repositories.CamionRepositoryInterface
	envioRepository  repositories.EnvioRepositoryInterface
}

func NewCamionService(camionRepository repositories.CamionRepositoryInterface) *camionService {
	return &camionService{camionRepository: camionRepository}
}

func (service *camionService) ObtenerCamiones() []*dto.Camion {
	camionDB, _ := service.camionRepository.OtenerCamiones()
	var camiones []*dto.Camion
	for _, camionDB := range camionDB {
		camion := dto.NewCamion(camionDB)
		camiones = append(camiones, camion)
	}
	return camiones
}

func (service *camionService) ObtenerCamionPorId(id string) *dto.Camion {
	camionDB, err := service.camionRepository.ObtenerCamionPorId(id)
	var camion *dto.Camion
	if err == nil {
		camion = dto.NewCamion(camionDB)
	} else {
		log.Printf("[service:CamionService][method:ObtenerCamionPorId][reason:NOT_FOUND][id:%d]", id)
	}
	return camion
}
func (service *camionService) InsertarCamion(camion *dto.Camion) bool {
	service.camionRepository.InsertarCamion(camion.GetModel())
	return true
}

func (service *camionService) EliminarCamion(id string) bool {
	service.camionRepository.EliminarCamion(utils.GetObjectIDFromStringID(id))
	return true
}

func (service *camionService) ActualizarCamion(camion *dto.Camion) bool {
	service.camionRepository.ActualizarCamion(camion.GetModel())
	return true
}
