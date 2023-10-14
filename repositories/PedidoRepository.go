// Crear una interface, struct y new CamionRepository
package repositories

type PedidoRepositoryInterface interface {
	//Metodos para implementar en el service
}
type PedidoRepository struct {
	db DB
}

func NewPedidoRepository(db DB) *PedidoRepository {
	return &PedidoRepository{db: db}
}

func CrearPedido