// Crear struct, new objeto y metodos
package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/juanperret/Directo-al-modelaje/dto"
	"github.com/juanperret/Directo-al-modelaje/services"
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
	if err := c.ShouldBindJSON(&patente); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//invocamos al metodo
	camion := handler.camionService.ObtenerCamionPorPatente(patente)
	//Agregamos un log para indicar informacion
	c.JSON(http.StatusOK, camion)
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
			c.JSON(http.StatusOK, gin.H{"error": resultado.Error()})
		}

	}

}

func (handler *CamionHandler) EliminarCamion(c *gin.Context) {
	id := c.Param("id")
	err := c.ShouldBindJSON(&id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"mensaje": err.Error()})
	} else {
		resultado := handler.camionService.EliminarCamion(id)
		if resultado != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": resultado.Error()})
		} else {
			c.JSON(http.StatusOK, gin.H{"error": resultado.Error()})
		}
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
			c.JSON(http.StatusOK, gin.H{"error": resultado.Error()})
		}
	}
}
