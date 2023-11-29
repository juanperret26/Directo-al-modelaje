// Crear interface, structura y new ProductoService
package services

import (
	"log"

	"github.com/juanperret/Directo-al-modelaje/go/dto"
	"github.com/juanperret/Directo-al-modelaje/go/repositories"
)

type ProductoInterface interface {
	ObtenerProductos() []*dto.Producto
	ObtenerProductosStockMinimo(tipoProducto string) []*dto.Producto
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

func (service *productoService) ObtenerProductos() []*dto.Producto {
	productoDB, _ := service.productoRepository.ObtenerProductos()
	var productos []*dto.Producto
	for _, productoDB := range productoDB {
		producto := dto.NewProducto(productoDB)
		productos = append(productos, producto)
	}
	return productos
}
func (service *productoService) ObtenerProductosStockMinimo(tipoProducto string) []*dto.Producto {
	productoDB, _ := service.productoRepository.ObtenerProductosStockMinimo(tipoProducto)
	var productos []*dto.Producto
	for _, productoDB := range productoDB {
		producto := dto.NewProducto(productoDB)
		if producto.Stock < producto.Stock_minimo && producto.TipoProducto == tipoProducto {
			productos = append(productos, producto)
		}
	}
	return productos
}

func (service *productoService) ObtenerProductoPorId(id string) *dto.Producto {
	productoDB, err := service.productoRepository.ObtenerProductoPorId(id)
	var producto *dto.Producto
	if err == nil {
		producto = dto.NewProducto(productoDB)
	} else {
		log.Printf("[service:productoService] [method:ObtenerProductoPorId] [reason: NOT_FOUND][id:%d]", id)
	}
	return producto
}
func (service *productoService) InsertarProducto(producto *dto.Producto) error {
	_, err := service.productoRepository.InsertarProducto(producto.GetModel())

	return err
}

func (service *productoService) ActualizarProducto(producto *dto.Producto) error {
	_, err := service.productoRepository.ActualizarProducto(producto.GetModel())

	return err
}

func (service *productoService) EliminarProducto(id string) error {
	_, err := service.productoRepository.EliminarProducto(id)
	return err
}
