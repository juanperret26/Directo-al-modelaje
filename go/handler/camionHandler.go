// Crear struct, new objeto y metodos
package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/juanperret26/Directo-al-modelaje/go/dto"
	"github.com/juanperret26/Directo-al-modelaje/go/services"
)

type CamionHandler struct {
	camionService services.CamionInterface
}

func NewCamionHandler(camionService services.CamionInterface) *CamionHandler {
	return &CamionHandler{camionService: camionService}
}

func (handler *CamionHandler) ObtenerCamiones(c *gin.Context) {
	camiones := handler.camionService.ObtenerCamiones()
	log.Printf("[handler:CamionHandler][method:ObtenerCamiones][cantidad:%d]", len(camiones))
	c.JSON(http.StatusOK, camiones)
}

func (handler *CamionHandler) ObtenerCamionPorPatente(c *gin.Context) {
	patente := c.Param("patente")
	camion := handler.camionService.ObtenerCamionPorPatente(patente)
	if camion == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "pedido no encontrado"})
	} else {
		c.JSON(http.StatusOK, camion)
	}
}

func (handler *CamionHandler) InsertarCamion(c *gin.Context) {
	var camion dto.Camion
	err := c.ShouldBindJSON(&camion)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

	} else {
		resultado := handler.camionService.InsertarCamion(&camion)
		if resultado != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": resultado.Error()})
		} else {
			c.JSON(http.StatusOK, "Creado Correctamente")
		}

	}

}

func (handler *CamionHandler) EliminarCamion(c *gin.Context) {
	id := c.Param("id")
	resultado := handler.camionService.EliminarCamion(id)

	if resultado != nil {
		c.JSON(http.StatusNotFound, gin.H{"mensaje": resultado.Error()})
	} else {

		c.JSON(http.StatusOK, "Se  elimino el camion correctamente")
	}
}

func (handler *CamionHandler) ActualizarCamion(c *gin.Context) {
	var camion dto.Camion
	err := c.ShouldBindJSON(&camion)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		resultado := handler.camionService.ActualizarCamion(&camion)
		if resultado != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": resultado.Error()})
		} else {
			c.JSON(http.StatusOK, "Creado Correctamente")
		}
	}
}
