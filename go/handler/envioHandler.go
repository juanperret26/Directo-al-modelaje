// Crear struct, new objeto y metodos
package handler

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/juanperret26/Directo-al-modelaje/go/dto"
	"github.com/juanperret26/Directo-al-modelaje/go/services"
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
	if envio == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "envio no encontrado"})
		return
	} else {
		resultado := handler.envioService.IniciarViaje(envio)
		if resultado != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": resultado.Error()})
		} else {
			c.JSON(http.StatusCreated, "Viaje iniciado correctamente")
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

func (handler *EnvioHandler) ObtenerEnviosFiltro(c *gin.Context) {
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
	filtro := dto.Filtro{
		PatenteCamion:      patente,
		EstadoEnvio:        estado,
		UltimaParada:       ultimaParada,
		FechaCreacionDesde: fechaCreacionDesde,
		FechaCreacionHasta: fechaCreacionHasta,
	}
	envios, err := handler.envioService.ObtenerEnviosFiltro(&filtro)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	} else {
		log.Printf("Se obtuvieron envios filtrados para patente %s, estado %s, ultimaParada %s, fechaCreacionDesde %s, fechaCreacionHasta %s", patente, estado, ultimaParada, fechaCreacionDesde, fechaCreacionHasta)
		c.JSON(http.StatusOK, envios)
	}

}

func (handler *EnvioHandler) InsertarEnvio(c *gin.Context) {
	log.Println("[handler:EnvioHandler][method:InsertarEnvio][info:Inicio]")
	var envio dto.Envio
	err := c.ShouldBindJSON(&envio)
	if err != nil {
		log.Printf("[handler:EnvioHandler][method:InsertarEnvio][reason:INVALID_INPUT][error:%v]", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	log.Printf("[handler:EnvioHandler][method:InsertarEnvio][info:Datos recibidos][envio:%+v]", envio)

	resultado := handler.envioService.InsertarEnvio(&envio)
	if resultado != nil {
		log.Printf("[handler:EnvioHandler][method:InsertarEnvio][reason:SERVICE_ERROR][error:%v]", resultado)
		c.JSON(http.StatusBadRequest, gin.H{"error": "No se pudo insertar el envio"})
		return
	}

	log.Println("[handler:EnvioHandler][method:InsertarEnvio][info:Envio insertado correctamente]")
	c.JSON(http.StatusCreated, gin.H{"message": "Envio insertado correctamente"})
}


func (handler *EnvioHandler) EliminarEnvio(c *gin.Context) {
	id := c.Param("id")
	resultado := handler.envioService.EliminarEnvio(id)
	if resultado != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": resultado.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"error": "Envio eliminado correctamente"})
	}
}

func (handler *EnvioHandler) ActualizarEnvio(c *gin.Context) {
    log.Println("[handler:EnvioHandler][method:ActualizarEnvio][info:Inicio]")
    
    var envio dto.Envio
    err := c.ShouldBindJSON(&envio)
    if err != nil {
        log.Printf("[handler:EnvioHandler][method:ActualizarEnvio][reason:INVALID_INPUT][error:%v]", err)
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    log.Printf("[handler:EnvioHandler][method:ActualizarEnvio][info:Datos recibidos][envio:%+v]", envio)

    // Cambiar el manejo del error
    if resultado := handler.envioService.ActualizarEnvio(&envio); resultado != nil {
        log.Printf("[handler:EnvioHandler][method:ActualizarEnvio][reason:SERVICE_ERROR][error:%v]", resultado)
        c.JSON(http.StatusBadRequest, gin.H{"error": resultado.Error()})
        return
    }

    log.Println("[handler:EnvioHandler][method:ActualizarEnvio][info:Envio actualizado correctamente]")
    c.JSON(http.StatusOK, gin.H{"message": "Envio actualizado correctamente"})
}


func (handler *EnvioHandler) ObtenerCantidadEnviosPorEstado(c *gin.Context) {
	estado := c.Param("estado")
	cantidad, err := handler.envioService.ObtenerCantidadEnviosPorEstado(estado)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

	} else {
		c.JSON(http.StatusOK, gin.H{"cantidad": cantidad})
	}

}
func (handler *EnvioHandler) AgregarParada(c *gin.Context) {
    // Recibimos la nueva parada desde el body
    var parada dto.Parada
    err := c.ShouldBindJSON(&parada)
    if err != nil {
        log.Printf("[handler:EnvioHandler][method:AgregarParada][error:Datos inválidos][detail:%s]", err.Error())
        c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos: " + err.Error()})
        return
    }
	id := c.Param("id")
    log.Printf("[handler:EnvioHandler][method:AgregarParada][info:Parada recibida][parada:%+v]", parada)

    envio := dto.Envio{
        Id: id, 
        Paradas: []dto.Parada{
            parada,
        },
    }
    log.Printf("[handler:EnvioHandler][method:AgregarParada][info:Envío preparado][envio:%+v]", envio)

		_, err := handler.envioService.AgregarParada(&envio)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		} else {
			//Agregamos un log para indicar información relevante del resultado
			log.Printf("[handler:EnvioHandler][method:AgregarParada][envio:%+v]", envio)
			c.JSON(http.StatusOK, gin.H{"error": "parada cargada correctamente"})
		}

    // Llamamos al servicio para agregar la parada
    success, err := handler.envioService.AgregarParada(&envio)
    if err != nil {
        log.Printf("[handler:EnvioHandler][method:AgregarParada][error:Servicio falló][detail:%s]", err.Error())
        c.JSON(http.StatusBadRequest, gin.H{"error": "Error al agregar la parada: " + err.Error()})
        return
    }

    if success {
        log.Printf("[handler:EnvioHandler][method:AgregarParada][info:Parada agregada correctamente][envio:%+v]", envio)
        c.JSON(http.StatusOK, gin.H{"message": "Parada agregada correctamente"})
    } else {
        log.Printf("[handler:EnvioHandler][method:AgregarParada][error:No se pudo agregar la parada][envio:%+v]", envio)
        c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo agregar la parada"})
    }
}

func (handler *EnvioHandler) ObtenerBeneficiosEntreFechas(c *gin.Context) {

	//Convierte las fechas string a time.Time
	fechaDesdeStr := c.DefaultQuery("fechaDesde", "0001-01-01")
	fechaDesde, err := time.Parse("2006-01-02", fechaDesdeStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Formato de fechaDesde incorrecto"})
		return
	}

	fechaHastaStr := c.DefaultQuery("fechaHasta", "0001-01-01")
	fechaHasta, err := time.Parse("2006-01-02", fechaHastaStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Formato de fechaHasta incorrecto"})
		return
	}

	filtro := dto.Filtro{
		FechaUltimaActualizacionDesde: fechaDesde,
		FechaUltimaActualizacionHasta: fechaHasta,
	}

	beneficio, err := handler.envioService.ObtnerBeneficiosEntreFecha(&filtro)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	log.Printf("Se obtuvieron beneficios entre fechas %s, %s", fechaDesde, fechaHasta)
	response := map[string]int{"beneficio": beneficio}
	c.JSON(http.StatusOK, response)

}
