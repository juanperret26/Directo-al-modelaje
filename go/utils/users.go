package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/juanperret26/Directo-al-modelaje/go/clients/responses"
)

const (
	RolAdministrador = "Administrador"
	RolUsuario       = "Usuario"
	RolConductor     = "Conductor"
)

func SetUserInContext(c *gin.Context, user *responses.UserInfo) {
	c.Set("UserInfo", user)
}

func GetUserInfoFromContext(c *gin.Context) *responses.UserInfo {
	userInfo, _ := c.Get("UserInfo")

	user, _ := userInfo.(*responses.UserInfo)

	return user
}
