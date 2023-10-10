// Crear struct, new objeto y metodos
package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/juanperret/Directo-al-modelaje/dto"
	"github.com/juanperret/Directo-al-modelaje/services"
)

type EnvioHandler struct {
	envioService services.EnvioInterface
}

func NewEnvioHandler(envioService services.EnvioInterface) *EnvioHandler {
	return &EnvioHandler{
		envioService: envioService,
	}
}
func (handler *EnvioHandler) GetEnvios(c *gin.Context) {
	//invocamos al metodo
	envios := handler.envioService.GetEnvios()
	//Agregamos un log para indicar informacion
	log.Printf("[handelr:EnvioHandler] [method:GetEnvios] [envios:%v] [cantidad:%d]", envios, len(envios))
	c.JSON(http.StatusOK, envios)
}
func (handler *EnvioHandler) GetEnvio(c *gin.Context) {
	id := c.Param("id")
	//invocamos al metodo
	envio := handler.envioService.GetEnvio(id)
	//Agregamos un log para indicar informacion
	c.JSON(http.StatusOK, envio)
}
func (handler *EnvioHandler) InsertarEnvio(c *gin.Context) {
	var envio dto.Envio
	if err := c.ShouldBindJSON(&envio); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	resultado := handler.envioService.InsertarEnvio(&envio)
	c.JSON(http.StatusCreated, resultado)
}
