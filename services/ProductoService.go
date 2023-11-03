// Crear interface, structura y new ProductoService
package services

import (
	"log"

	"github.com/juanperret/Directo-al-modelaje/dto"
	"github.com/juanperret/Directo-al-modelaje/repositories"
)

type ProductoInterface interface {
	ObtenerProductos() []*dto.Producto
	ObtenerProductosStockMinimo(tipoProducto string) []*dto.Producto
	ObtenerProductoPorId(id string) *dto.Producto
	InsertarProducto(producto *dto.Producto) bool
	EliminarProducto(id string) bool
	ActualizarProducto(producto *dto.Producto) bool
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
func (service *productoService) InsertarProducto(producto *dto.Producto) bool {
	service.productoRepository.InsertarProducto(producto.GetModel())

	return true
}

func (service *productoService) ActualizarProducto(producto *dto.Producto) bool {
	service.productoRepository.ActualizarProducto(producto.Id)

	return true
}

func (service *productoService) EliminarProducto(id string) bool {
	service.productoRepository.EliminarProducto(id)
	return true
}
