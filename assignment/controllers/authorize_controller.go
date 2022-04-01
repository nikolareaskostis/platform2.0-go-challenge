package api

import (
	request "GoRepo/source/assignment/codeHelpers"
	"GoRepo/source/assignment/models"
	"GoRepo/source/assignment/services"

	"github.com/gin-gonic/gin"
)

type Authorize_Controller interface {
	Login(c *gin.Context)
	Register(c *gin.Context)
	Check(c *gin.Context)
}

type authorize_controller struct {
	authorizeService services.AuthorizeService
	tokenService     services.TokenService
}

func NewAuthorizeController(authorizeService services.AuthorizeService, tokenService services.TokenService) Authorize_Controller {
	return &authorize_controller{authorizeService: authorizeService, tokenService: tokenService}
}

func (a *authorize_controller) Login(c *gin.Context) {
	var user models.User
	err := c.BindJSON(&user)
	if err == nil {
		result, err := a.authorizeService.Login(user)

		if err != nil {
			request.BadRequest(c)
			return
		}

		request.OK(c, result)
	}
	request.BadRequest(c)
}

func (a *authorize_controller) Register(c *gin.Context) {
	var user models.User
	err := c.BindJSON(&user)
	if err == nil {
		result, err := a.authorizeService.Register(user)

		if err != nil {
			request.BadRequest(c)
			return
		}

		request.OK(c, result)
	}
	request.BadRequest(c)
}

func (a *authorize_controller) Check(c *gin.Context) {
	var user models.User
	err := c.BindJSON(&user)
	if err == nil {
		result, err := a.authorizeService.Check(user)

		if err != nil {
			request.BadRequest(c)
			return
		}

		request.OK(c, result)
	}
	request.BadRequest(c)
}
