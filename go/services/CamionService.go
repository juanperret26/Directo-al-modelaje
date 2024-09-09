// Crear interface, structura y new CamionService
package services

import (
	"log"

	"github.com/juanperret26/Directo-al-modelaje/go/dto"
	"github.com/juanperret26/Directo-al-modelaje/go/repositories"
)

type CamionInterface interface {
	//Metodos para implementar en el handler
	ObtenerCamiones() []*dto.Camion
	ObtenerCamionPorPatente(patente string) *dto.Camion
	InsertarCamion(camion *dto.Camion) error
	EliminarCamion(id string) error
	ActualizarCamion(camion *dto.Camion) error
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

func (service *camionService) ObtenerCamionPorPatente(id string) *dto.Camion {
	camionDB, err := service.camionRepository.ObtenerCamionPorPatente(id)
	var camion *dto.Camion
	if err == nil {
		camion = dto.NewCamion(camionDB)
	} else {
		log.Printf("[service:CamionService][method:ObtenerCamionPorPatente][reason:NOT_FOUND][patente:%d]", camion.Patente)
	}
	return camion
}
func (service *camionService) InsertarCamion(camion *dto.Camion) error {
	_, err := service.camionRepository.InsertarCamion(camion.GetModel())
	if err != nil {
		log.Printf("[service:CamionService][method:InsertarCamion][reason:ERROR][error:%v]", err)
	}
	return err
}

func (service *camionService) EliminarCamion(id string) error {
	_, err := service.camionRepository.EliminarCamion(id)
	if err != nil {
		log.Printf("[service:CamionService][method:EliminarCamion][reason:ERROR][id:%d]", id)
	}

	return err
}

func (service *camionService) ActualizarCamion(camion *dto.Camion) error {
	_, err := service.camionRepository.ActualizarCamion(camion.GetModel())
	if err != nil {
		log.Printf("[service:CamionService][method:ActualizarCamion][reason:ERROR][error:%v]", err)
	}
	return err
}
