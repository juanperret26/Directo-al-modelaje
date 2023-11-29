// Crear struct, new objeto y metodos
package handler

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
<<<<<<< HEAD:go/handler/envioHandler.go
	"github.com/juanperret26/Directo-al-modelaje/go/dto"
	"github.com/juanperret26/Directo-al-modelaje/go/services"
=======
	"github.com/juanperret/Directo-al-modelaje/dto"
	"github.com/juanperret/Directo-al-modelaje/services"
>>>>>>> parent of ad42c9c (prueba cambio de directorio):handler/envioHandler.go
)

type EnvioHandler struct {
	envioService services.EnvioInterface
}

func NewEnvioHandler(envioService services.EnvioInterface) *EnvioHandler {
	return &EnvioHandler{
		envioService: envioService,
	}
}
func (handler *EnvioHandler) ObtenerEnvios(c *gin.Context) {
	//invocamos al metodo
	envios := handler.envioService.ObtenerEnvios()
	if envios == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No se encontraron envios"})
	} else {
		log.Printf("[handler:EnvioHandler] [method:ObtenerEnvios] [envios:%v] [cantidad:%d]", envios, len(envios))
		c.JSON(http.StatusOK, envios)
	}

}
func (handler *EnvioHandler) IniciarViaje(c *gin.Context) {
	id := c.Param("id")
	envio := handler.envioService.ObtenerEnvioPorId(id)
	if err := c.ShouldBindJSON(&envio); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} else {
		resultado := handler.envioService.IniciarViaje(envio)
		if resultado != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": resultado.Error()})
		} else {
			c.JSON(http.StatusCreated, gin.H{"estado": "Viaje iniciado correctamente"})
		}

	}
}

func (handler *EnvioHandler) ObtenerEnvioPorId(c *gin.Context) {
	id := c.Param("id")
	//invocamos al metodo
	envio := handler.envioService.ObtenerEnvioPorId(id)
	if envio == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "envio no encontrado"})
	} else {
		//Agregamos un log para indicar informacion
		c.JSON(http.StatusOK, envio)
	}

}

func (handler *EnvioHandler) ObtenerPedidosFiltrados(c *gin.Context) {
	codigoEnvio := c.Param("codigoEnvio")
	estado := c.Param("estado")
	fechaInicioStr := c.Param("fechaInicio")
	fechaFinalStr := c.Param("fechaFinal")

	// Convertir strings a time.Time
	fechaInicio, errInicio := time.Parse("2006-01-02", fechaInicioStr)
	if errInicio != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Formato de fechaInicio incorrecto"})
		return
	}

	fechaFinal, errFinal := time.Parse("2006-01-02", fechaFinalStr)
	if errFinal != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Formato de fechaFinal incorrecto"})
		return
	}

	// Manejo de errores
	pedidos, err := handler.envioService.ObtenerPedidosFiltrados(codigoEnvio, estado, fechaInicio, fechaFinal)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	} else {
		// Registro de información
		log.Printf("Se obtuvieron pedidos filtrados para codigoEnvio %s, estado %s, fechaInicio %s, fechaFinal %s", codigoEnvio, estado, fechaInicio, fechaFinal)
		c.JSON(http.StatusOK, pedidos)
	}
}

func (handler *EnvioHandler) ObtenerEnviosPorParametros(c *gin.Context) {
	patente := c.Param("patente")
	estado := c.Param("estado")
	ultimaParada := c.Param("ultimaParada")
	fechaCreacionDesdeStr := c.Param("fechaCreacionDesde")
	fechaCreacionHastaStr := c.Param("fechaCreacionHasta")

	fechaCreacionDesde, errDesde := time.Parse("2006-01-02", fechaCreacionDesdeStr)
	if errDesde != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Formato de fechaCreacionDesde incorrecto"})
		return
	}

	fechaCreacionHasta, errHasta := time.Parse("2006-01-02", fechaCreacionHastaStr)
	if errHasta != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Formato de fechaCreacionHasta incorrecto"})
		return
	}

	envios, err := handler.envioService.ObtenerEnviosPorParametros(patente, estado, ultimaParada, fechaCreacionDesde, fechaCreacionHasta)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	} else {
		log.Printf("Se obtuvieron envios filtrados para patente %s, estado %s, ultimaParada %s, fechaCreacionDesde %s, fechaCreacionHasta %s", patente, estado, ultimaParada, fechaCreacionDesde, fechaCreacionHasta)
		c.JSON(http.StatusOK, envios)
	}

}

func (handler *EnvioHandler) InsertarEnvio(c *gin.Context) {

	var envio dto.Envio
	err := c.ShouldBindJSON(&envio)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		resultado := handler.envioService.InsertarEnvio(&envio)
		if resultado != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": resultado.Error()})
		} else {
			c.JSON(http.StatusOK, gin.H{"error": "Creado correctamente"})
		}
	}
}

func (handler *EnvioHandler) EliminarEnvio(c *gin.Context) {
	id := c.Param("id")
	resultado := handler.envioService.EliminarEnvio(id)
	if resultado != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": resultado.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"error": "Eliminado correctamente"})
	}
}

func (handler *EnvioHandler) ActualizarEnvio(c *gin.Context) {
	var envio dto.Envio
	err := c.ShouldBindJSON(&envio)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

	} else {
		resultado := handler.envioService.ActualizarEnvio(&envio)
		if resultado != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": resultado.Error()})
		} else {
			c.JSON(http.StatusOK, gin.H{"error": "Actualizado correctamente"})
		}
	}

}
func (handler *EnvioHandler) ObtenerCantidadEnviosPorEstado(c *gin.Context) {
	estado := c.Param("estado")
	cantidad, err := handler.envioService.ObtenerCantidadEnviosPorEstado(estado)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

	} else {
		c.JSON(http.StatusOK, cantidad)
	}

}
func (handler *EnvioHandler) AgregarParada(c *gin.Context) {

	//Recibimos el id como parametro
	id := c.Param("id")

	//Obtenemos la nueva parada
	var parada dto.Parada
	err := c.ShouldBindJSON(&parada)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		envio := dto.Envio{
			Id: id,
			Paradas: []dto.Parada{
				parada,
			},
		}

		_, err := handler.envioService.AgregarParada(&envio)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		} else {
			//Agregamos un log para indicar información relevante del resultado
			log.Printf("[handler:EnvioHandler][method:AgregarParada][envio:%+v][user:%s]", envio)
			c.JSON(http.StatusOK, gin.H{"estado": "Parada agregada correctamente"})
		}

	}
}
func (handler *EnvioHandler) ObtenerBeneficiosEntreFechas(c *gin.Context) {
	fechaDesdeStr := c.DefaultQuery("fechaDesde", "0001-01-01T00:00:00Z")
	fechaDesde, err := time.Parse(time.RFC3339, fechaDesdeStr)
	if err != nil {
		fechaDesde = time.Time{}
	}
	fechaHastaStr := c.DefaultQuery("fechaHasta", "0001-01-01T00:00:00Z")
	fechaHasta, err := time.Parse(time.RFC3339, fechaHastaStr)
	if err != nil {
		fechaHasta = time.Time{}
	}
	// Manejo de errores
	beneficios, err := handler.envioService.ObtnerBeneficiosEntreFecha(fechaDesde, fechaHasta)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Registro de información
	log.Printf("Se obtuvieron beneficios entre fechas %s, %s", fechaDesde, fechaHasta)
	response := map[string]int{"beneficio": beneficios}
	// Respuesta exitosa
	c.JSON(http.StatusOK, response)
}
