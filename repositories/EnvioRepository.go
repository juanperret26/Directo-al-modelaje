// Crear una interface, struct y new CamionRepository
package repositories

type EnvioRepositoryInterface interface {
	//Metodos para implementar en el service
}
type EnvioRepository struct {
	db DB
}

func NewEnvioRepository(db DB) *EnvioRepository {
	return &EnvioRepository{db: db}
}
