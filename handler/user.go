package handler

import (
	"pustaka-api/loginservice"

	"github.com/gin-gonic/gin"
)

type LoginController interface {
	Login(ctx *gin.Context) string
}

type loginController struct {
	loginService loginservice.LoginService
	jwtService loginservice.JWTService
}

func LoginHandler(loginService loginservice.LoginService, jwtService loginservice.JWTService) LoginController {
	return &loginController{
		loginService: loginService,
		jwtService: jwtService,
	}
}

func (l *loginController) Login(ctx *gin.Context) string {
	var credential loginservice.LoginCredentials
	err := ctx.ShouldBind(&credential)
	if err != nil {
		return "no data found"
	}
	isUserAuthenticated := loginservice.LoginService.LoginUser(credential.Username, credential.Password)
	if isUserAuthenticated {
		return loginservice.JWTService.GenerateToken(credential.Username, true)

	}
	return ""
}

