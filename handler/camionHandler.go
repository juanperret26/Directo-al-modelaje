// Crear struct, new objeto y metodos
package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/juanperret/Directo-al-modelaje/services"
)

type CamionHandler struct {
	camionService services.CamionInterface
}

func NewCamionHandler(camionService services.CamionInterface) *CamionHandler {
	return &CamionHandler{camionService: camionService}
}
func (handler *CamionHandler) GetCamion() (c *gin.Context) {
	return handler.GetCamion() //No Hecho
}
