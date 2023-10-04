// Crear interface, structura y new AulaService
package services

import "github.com/juanperret/Directo-al-modelaje/repositories"
type CamionInterface interface {
	//Metodos para implementar en el handler
}
type camionService struct {
	camionRepository repositories.CamionRepositoryInterface
}
func NewCamionService(camionRepository repositories.CamionRepositoryInterface) *camionService {
	return &camionService{camionRepository: camionRepository}
}