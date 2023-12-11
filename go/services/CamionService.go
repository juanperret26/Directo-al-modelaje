// Crear interface, structura y new CamionService
package services

import (
	"errors"
	"log"

	"github.com/juanperret/Directo-al-modelaje/go/dto"
	"github.com/juanperret/Directo-al-modelaje/go/repositories"
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

	if camion.Patente != "" && camion.Costo_km != 0 && camion.Peso_maximo != 0 {
		_, err := service.camionRepository.InsertarCamion(camion.GetModel())
		return err
	} else {
		err := errors.New("No se pasaron bien los datos")
		return err
	}

}

func (service *camionService) EliminarCamion(id string) error {
	_, err := service.camionRepository.EliminarCamion(id)
	return err
}

func (service *camionService) ActualizarCamion(camion *dto.Camion) error {
	_, err := service.camionRepository.ActualizarCamion(camion.GetModel())
	return err
}
