package services

import (
	"errors"
	"log"

	"github.com/juanperret26/Directo-al-modelaje/go/dto"
	"github.com/juanperret26/Directo-al-modelaje/go/model"
	"github.com/juanperret26/Directo-al-modelaje/go/repositories"
	"github.com/juanperret26/Directo-al-modelaje/go/utils"
)

type EnvioInterface interface {
	ObtenerEnvios() []*dto.Envio
	ObtenerEnvioPorId(id string) *dto.Envio
	ObtenerPedidosFiltro(filtro *dto.Filtro) ([]*dto.Pedido, error)
	ObtenerEnviosFiltro(filtro *dto.Filtro) ([]*dto.Envio, error)
	InsertarEnvio(envio *dto.Envio) error
	EliminarEnvio(id string) error
	ActualizarEnvio(envio *dto.Envio) error
	IniciarViaje(envio *dto.Envio) error
	ObtenerCantidadEnviosPorEstado(estado string) (int, error)
	AgregarParada(envio *dto.Envio) (bool, error)
	ObtnerBeneficiosEntreFecha(fecha *dto.Filtro) (int, error)

}

type envioService struct {
	envioRepository    repositories.EnvioRepositoryInterface
	pedidoRepository   repositories.PedidoRepositoryInterface
	productoRepository repositories.ProductoRepositoryInterface
	camionRepository   repositories.CamionRepositoryInterface
}

func NewEnvioService(envioRepository repositories.EnvioRepositoryInterface, camionRepository repositories.CamionRepositoryInterface, pedidoRepository repositories.PedidoRepositoryInterface, productoRepository repositories.ProductoRepositoryInterface) *envioService {
	return &envioService{
		envioRepository:    envioRepository,
		camionRepository:   camionRepository,
		pedidoRepository:   pedidoRepository,
		productoRepository: productoRepository,
	}
}

func (service *envioService) ObtenerEnvios() []*dto.Envio {
	envioDB, err := service.envioRepository.ObtenerEnvios()
	if err != nil {
		log.Printf("[service:EnvioService][method:ObtenerEnvios][reason:ERROR][error:%v]", err)
		return nil
	}
	var envios []*dto.Envio
	for _, envioDB := range envioDB {
		envio := dto.NewEnvio(envioDB)
		envios = append(envios, envio)
	}
	return envios
}

func (service *envioService) ObtenerEnviosFiltro(filtro *dto.Filtro) ([]*dto.Envio, error) {

	enviosDB, err := service.envioRepository.ObtenerEnviosFiltro(filtro.PatenteCamion, filtro.EstadoEnvio, filtro.UltimaParada, filtro.FechaCreacionDesde, filtro.FechaCreacionHasta)
	if err != nil {
		log.Printf("[service:EnvioService][method:ObtenerEnviosFiltro][reason:ERROR][error:%v]", err)
		return nil, err
	}
	var envios []*dto.Envio
	for _, envioDB := range enviosDB {
		envio := dto.NewEnvio(envioDB)
		envios = append(envios, envio)
	}
	return envios, nil
}

func (service *envioService) ObtenerEnvioPorId(id string) *dto.Envio {
	if id == "" {
		log.Println("[service:EnvioService][method:ObtenerEnvioPorId][reason:INVALID_INPUT][message:ID vacío]")
		return nil
	}
	envioDB, err := service.envioRepository.ObtenerEnvioPorId(id)
	if err != nil {
		log.Printf("[service:EnvioService][method:ObtenerEnvioPorId][reason:NOT_FOUND][id:%s]", id)
		return nil
	}
	envio := dto.NewEnvio(envioDB)
	return envio
}

func (service *envioService) ObtenerPedidosFiltro(filtro *dto.Filtro) ([]*dto.Pedido, error) {
	if filtro == nil {
		err := errors.New("el filtro no puede ser nulo")
		log.Printf("[service:EnvioService][method:ObtenerPedidosFiltro][reason:INVALID_INPUT][error:%v]", err)
		return nil, err
	}
	pedidoDB, err := service.envioRepository.ObtenerPedidosFiltro(filtro.CodigoEnvio, filtro.EstadoPedido, filtro.FechaCreacionDesde, filtro.FechaCreacionHasta)
	if err != nil {
		log.Printf("[service:EnvioService][method:ObtenerPedidosFiltro][reason:ERROR][error:%v]", err)
		return nil, err
	}
	var pedidos []*dto.Pedido
	for _, pedidoDB := range pedidoDB {
		pedido := dto.NewPedido(pedidoDB)
		pedidos = append(pedidos, pedido)
	}
	return pedidos, nil
}

func (service *envioService) InsertarEnvio(envio *dto.Envio) error {
	if envio == nil {
		return errors.New("el envio no puede ser nulo")
	}

	if service.camionRepository == nil {
		return errors.New("el repositorio de camiones no está inicializado")
	}

	camion, err := service.camionRepository.ObtenerCamionPorPatente(envio.PatenteCamion)
	if err != nil {
		return err
	}

	pesoMaximo := camion.Peso_maximo
	pedidos := envio.Pedido
	var pedidosEnvio []model.Pedido
	for _, pedido := range pedidos {
		pedidoBuscado, _ := service.pedidoRepository.ObtenerPedidoPorId(pedido)
		pedidosEnvio = append(pedidosEnvio, pedidoBuscado)
	}
	pesoTotalPedidos := 0.0
	for _, pedido := range pedidosEnvio {
		for _, productoPedido := range pedido.PedidoProductos {
			pesoTotalPedidos += productoPedido.Precio_unitario * float64(productoPedido.Cantidad)
		}
	}
	if pesoTotalPedidos > pesoMaximo {
		return errors.New("El peso total de los pedidos supera el peso máximo del camión")
	}
	for _, pedido := range pedidosEnvio {
		service.pedidoRepository.ActualizarPedido(pedido)
	}
	err = service.envioRepository.InsertarEnvio(envio.GetModel())
	if err != nil {
		return err
	}

	return nil
}

func (service *envioService) EliminarEnvio(id string) error {
	if id == "" {
		err := errors.New("ID vacío")
		log.Printf("[service:EnvioService][method:EliminarEnvio][reason:INVALID_INPUT][error:%v]", err)
		return err
	}
	_, err := service.envioRepository.EliminarEnvio(utils.GetObjectIDFromStringID(id))
	if err != nil {
		log.Printf("[service:EnvioService][method:EliminarEnvio][reason:ERROR][id:%s][error:%v]", id, err)
	}
	return err
}

func (service *envioService) ActualizarEnvio(envio *dto.Envio) error {
    if envio == nil {
        err := errors.New("el envio no puede ser nulo")
        log.Printf("[service:EnvioService][method:ActualizarEnvio][reason:INVALID_INPUT][error:%v]", err)
        return err
    }
    
    if err := service.envioRepository.ActualizarEnvio(envio.GetModel()); err != nil {
        log.Printf("[service:EnvioService][method:ActualizarEnvio][reason:ERROR][error:%v]", err)
        return err
    }
    
    return nil
}

func (service envioService) AgregarParada(envio *dto.Envio) (bool, error) {
	//En teoria, recibimos un envio que tiene solamente id y la nueva parada
	//Primero buscamos el envio por id
	envioDB, err := service.envioRepository.ObtenerEnvioPorId(envio.Id)
	camion, err := service.camionRepository.ObtenerCamionPorPatente(envioDB.PatenteCamion)
	if err != nil {
		log.Printf("[service:EnvioService][method:AgregarParada][reason:NOTFOUND][id:%d]", envio.Id)
		return false, err
	}

	//Validamos que el envio esté en estado EnRuta
	// if envioDB.Estado !="En Viaje" {
	// 	return false, errors.New("el envio no esta en viaje")
	// }

	//Agregamos la nueva parada al envio
	envioDB.Paradas = append(envioDB.Paradas, envio.Paradas[0].GetModel())
	envioDB.Costo = envioDB.Costo + envio.Paradas[0].Kilometros*camion.Costo_km
	if envioDB.Destino.Nombre_ciudad == envio.Paradas[0].Ciudad {
		envioDB.Estado = "Despachado"
		for _, idPedido := range envioDB.Pedido {
			pedido, err := service.pedidoRepository.ObtenerPedidoPorId(idPedido)
			if err != nil {
				log.Printf("[service:PedidoService][method:ObtenerPedidosPorId][reason:NOT_FOUND][id:%d]", idPedido)
			}
			pedido.Estado = "Enviado"
			service.pedidoRepository.ActualizarPedido(pedido)
		}

	}
	//Actualizamos el envio en la base de datos, que ahora tiene la nueva parada
	return true, service.envioRepository.ActualizarEnvio(envioDB)
}

func (service *envioService) IniciarViaje(envio *dto.Envio) error {
	if envio == nil || envio.PatenteCamion == "" {
		err := errors.New("el envio o la patente del camión no pueden ser nulos")
		log.Printf("[service:EnvioService][method:IniciarViaje][reason:INVALID_INPUT][error:%v]", err)
		return err
	}

	// Cambiar el estado del envío a "En viaje"
	envio.Estado = "En viaje"
	err := service.envioRepository.ActualizarEnvio(envio.GetModel())
	if err != nil {
		log.Printf("[service:EnvioService][method:IniciarViaje][reason:ERROR][error:%v]", err)
		return err
	}

	log.Printf("[service:EnvioService][method:IniciarViaje][reason:SUCCESS][envio_id:%s]", envio.Id)
	return nil
}

func (service *envioService) ObtenerCantidadEnviosPorEstado(estado string) (int, error) {
	if estado == "" {
		err := errors.New("el estado no puede ser vacío")
		log.Printf("[service:EnvioService][method:ObtenerCantidadEnviosPorEstado][reason:INVALID_INPUT][error:%v]", err)
		return 0, err
	}
	cantidades, err := service.envioRepository.ObtenerCantidadEnviosPorEstado(estado)
	if err != nil {
		log.Printf("[service:EnvioService][method:ObtenerCantidadEnviosPorEstado][reason:ERROR][error:%v]", err)
		return cantidades, err
	}
	return cantidades, nil
}

func (service envioService) ObtnerBeneficiosEntreFecha(fecha *dto.Filtro) (int, error) {
	var beneficio int = 0

	if fecha.FechaCreacionDesde.IsZero() && fecha.FechaCreacionHasta.IsZero() {
		err := errors.New("las fechas no pueden ser nulas")
		log.Printf("[service:EnvioService][method:ObtnerBeneficiosEntreFecha][reason:INVALIDINPUT][error:%v]", err)
		return 0, err
	}
	envios, err := service.envioRepository.ObtenerEnviosFiltro("", "", "", fecha.FechaCreacionDesde, fecha.FechaCreacionHasta)
	if err != nil {
		return 0, err
	}
	for _, envio := range envios {
		for _, pedido := range envio.Pedido {
			pedidoDB, err := service.pedidoRepository.ObtenerPedidoPorId(pedido)
			if err != nil {
				return 0, err
			}
			for _, productoPedido := range pedidoDB.PedidoProductos {
				// Buscar el producto correspondiente al codigo
				producto, err := service.productoRepository.ObtenerProductoPorId(productoPedido.CodigoProducto)
				if err != nil {
					log.Printf("[service:ProductoService][method:ObtenerProductoPorId][reason:NOT_FOUND][id:%d]", productoPedido.CodigoProducto)
				}
				beneficio += int(producto.Precio) * productoPedido.Cantidad
			}
		}
	}
	return beneficio, nil
}

