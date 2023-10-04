// Crear una interface, struct y new CamionRepository
package repositories

type CamionRepositoryInterface interface {
	//Metodos para implementar en el service
}
type CamionRepository struct {
	// DB represents a database connection.
	db DB
}

func NewCamionRepository(db DB) *CamionRepository {
	return &CamionRepository{db: db}
}
