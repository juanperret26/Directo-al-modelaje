// Crear interface, structura y new ProductoService
package services

import (
	"errors"
	"log"

	"github.com/juanperret26/Directo-al-modelaje/go/dto"
	"github.com/juanperret26/Directo-al-modelaje/go/repositories"
)

type ProductoInterface interface {

    ObtenerProductos(filtroStockMinimo int) []*dto.Producto
	ObtenerProductoPorId(id string) *dto.Producto
	InsertarProducto(producto *dto.Producto) error
	EliminarProducto(id string) error
	ActualizarProducto(producto *dto.Producto) error
}

type productoService struct {
	productoRepository repositories.ProductoRepositoryInterface
}

func NewProductoService(productoRepository repositories.ProductoRepositoryInterface) *productoService {
	return &productoService{productoRepository: productoRepository}
}

func (service *productoService) ObtenerProductos(filtroStockMinimo int) []*dto.Producto {
	productoDB, _ := service.productoRepository.ObtenerProductos(filtroStockMinimo)
	var productos []*dto.Producto

	for _, productoDB := range productoDB {
		producto := dto.NewProducto(productoDB)

		// Si filtroStockMinimo > 0, filtrar los productos que tengan stock >= filtroStockMinimo
		if filtroStockMinimo > 0 && producto.Stock < filtroStockMinimo {
			continue
		}

		productos = append(productos, producto)
	}
	return productos
}



func (service *productoService) ObtenerProductoPorId(id string) *dto.Producto {
	productoDB, err := service.productoRepository.ObtenerProductoPorId(id)
	if err != nil {
		log.Printf("[service:productoService] [method:ObtenerProductoPorId] [reason: NOT_FOUND][id:%d]", id)
	}
	producto := dto.NewProducto(productoDB)
	return producto
}


func (service *productoService) InsertarProducto(producto *dto.Producto) error {

	if producto.Stock != 0 && producto.Precio != 0 && producto.Nombre != "" && producto.TipoProducto != "" {
		_, err := service.productoRepository.InsertarProducto(producto.GetModel())
		return err
	} else {
		err := errors.New("No se pasaron bien los datos")
		return err
	}
}

func (service *productoService) ActualizarProducto(producto *dto.Producto) error {
	_, err := service.productoRepository.ActualizarProducto(producto.GetModel())
	return err
}

func (service *productoService) EliminarProducto(id string) error {
	_, err := service.productoRepository.EliminarProducto(id)
	if err != nil {
		log.Printf("[service:productoService] [method:EliminarProducto] [reason: NOT_FOUND][id:%d]", id)
	}
	return err
}
