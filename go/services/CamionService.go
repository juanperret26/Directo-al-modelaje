// Crear interface, structura y new CamionService
package services

import (
	"errors"
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

func NewCamionService(
	camionRepository repositories.CamionRepositoryInterface,
	envioRepository repositories.EnvioRepositoryInterface,
) *camionService {
	return &camionService{
		camionRepository: camionRepository,
		envioRepository:  envioRepository,
	}
}

func (service *camionService) ObtenerCamiones() []*dto.Camion {
	camionDB, err := service.camionRepository.OtenerCamiones()
	if err != nil {
		log.Printf("[service:CamionService][method:ObtenerCamiones][reason:ERROR][error:%v]", err)
		return nil
	}

	var camiones []*dto.Camion
	for _, camionDB := range camionDB {
		camion := dto.NewCamion(camionDB)
		camiones = append(camiones, camion)
	}
	return camiones
}

func (service *camionService) ObtenerCamionPorPatente(id string) *dto.Camion {
	if id == "" {
		err := errors.New("la patente no puede estar vacía")
		log.Printf("[service:CamionService][method:ObtenerCamionPorPatente][reason:INVALID_INPUT][error:%v]", err)
		return nil
	}

	camionDB, err := service.camionRepository.ObtenerCamionPorPatente(id)
	if err != nil {
		log.Printf("[service:CamionService][method:ObtenerCamionPorPatente][reason:NOT_FOUND][patente:%s]", id)
		return nil
	}

	return dto.NewCamion(camionDB)
}

func (service *camionService) InsertarCamion(camion *dto.Camion) error {
	if camion == nil {
		err := errors.New("el objeto camion es nulo")
		log.Printf("[service:CamionService][method:InsertarCamion][reason:ERROR][error:%v]", err)
		return err
	}
	_, err := service.camionRepository.InsertarCamion(camion.GetModel())
	if err != nil {
		log.Printf("[service:CamionService][method:InsertarCamion][reason:ERROR][error:%v]", err)
	}

	return err
}

func (service *camionService) EliminarCamion(id string) error {
	if id == "" {
		err := errors.New("la patente del camión no puede estar vacía")
		log.Printf("[service:CamionService][method:EliminarCamion][reason:INVALID_INPUT][error:%v]", err)
		return err
	}

	_, err := service.camionRepository.EliminarCamion(id)
	if err != nil {
		log.Printf("[service:CamionService][method:EliminarCamion][reason:ERROR][id:%s]", id)
	}
	return err
}

func (service *camionService) ActualizarCamion(camion *dto.Camion) error {
	if camion == nil {
		err := errors.New("el objeto camion es nulo")
		log.Printf("[service:CamionService][method:ActualizarCamion][reason:INVALID_INPUT][error:%v]", err)
		return err
	}
	_, err := service.camionRepository.ActualizarCamion(camion.GetModel())
	if err != nil {
		log.Printf("[service:CamionService][method:ActualizarCamion][reason:ERROR][error:%v]", err)
	}
	return err
}
