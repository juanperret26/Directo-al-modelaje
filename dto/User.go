package dto

<<<<<<< HEAD:go/dto/User.go
import "github.com/juanperret26/Directo-al-modelaje/go/clients/responses"
=======
import "github.com/juanperret/Directo-al-modelaje/clients/responses"
>>>>>>> parent of ad42c9c (prueba cambio de directorio):dto/User.go

type User struct {
	Codigo   string `json:"codigo"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Rol      string `json:"rol"`
}

func NewUser(userInfo *responses.UserInfo) User {
	user := User{}
	if userInfo != nil {
		user.Codigo = userInfo.Codigo
		user.Email = userInfo.Email
		user.Username = userInfo.Username
		user.Rol = userInfo.Rol
	}
	return user
}
