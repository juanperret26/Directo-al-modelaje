// Crear interface, structura y new PedidoService
package services

import "github.com/juanperret/Directo-al-modelaje/repositories"

type PedidoInterface interface {
	//Metodos para implementar en el handler
}
type pedidoService struct {
	pedidoRepository repositories.PedidoRepositoryInterface
}

func NewPedidoService(pedidoRepository repositories.PedidoRepositoryInterface) *pedidoService {
	return &pedidoService{
		pedidoRepository: pedidoRepository,
	}
}
