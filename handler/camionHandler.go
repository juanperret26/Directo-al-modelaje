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

func (handler *CamionHandler) ObtenerCamionPorId(c *gin.Context) {
	id := c.Param("id")
	//invocamos al metodo
	camion := handler.camionService.ObtenerCamionPorId(id)
	//Agregamos un log para indicar informacion
	c.JSON(http.StatusOK, camion)
}

func (handler *CamionHandler) InsertarCamion(c *gin.Context) {
	var camion dto.Camion
	if err := c.ShouldBindJSON(&camion); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	resultado := handler.camionService.InsertarCamion(&camion)
	c.JSON(http.StatusCreated, resultado)
}

func (handler *CamionHandler) EliminarCamion(c *gin.Context) {
	id := c.Param("id")
	resultado := handler.camionService.EliminarCamion(id)
	c.JSON(http.StatusOK, resultado)
}

func (handler *CamionHandler) ActualizarCamion(c *gin.Context) {
	var camion dto.Camion
	if err := c.ShouldBindJSON(&camion); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	resultado := handler.camionService.ActualizarCamion(&camion)
	c.JSON(http.StatusOK, resultado)

}
