// Crear interface, structura y new EnvioService
package services

import "github.com/juanperret/Directo-al-modelaje/repositories"

type EnvioInterface interface {
	//Metodos para implementar en el handler
}
type envioService struct {
	envioRepository repositories.EnvioRepositoryInterface
}

func NewEnvioService(envioRepository repositories.EnvioRepositoryInterface) *envioService {
	return &envioService{
		envioRepository: envioRepository,
	}
}
