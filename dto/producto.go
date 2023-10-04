//Crear struct y un new producto con los metodos que sean necesario
package dto
import (
	"time"
	"github.com/juanperret/Directo-al-modelaje/model"
)
type Producto struct {
	id_Producto int
	nombre string
	peso_unitario float64
	precio float64
	stock int
	stock_minimo int
	tipo string
	ultima_actualizacion time.Date
	fecha_creacion time.Date
}
func NewProducto(producto model.Producto) *Producto {
	return &Producto{
		id_Producto: model.GetIntIDFromObjectID(producto.id_Producto),
		nombre: producto.nombre,
		peso_unitario: producto.peso_unitario,
		precio: producto.precio,
		stock: producto.stock,
		stock_minimo: producto.stock_minimo,
		tipo: producto.tipo,
		ultima_actualizacion: producto.ultima_actualizacion,
		fecha_creacion: producto.fecha_creacion,
	}
}