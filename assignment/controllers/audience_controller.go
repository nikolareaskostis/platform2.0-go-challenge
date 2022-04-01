package api

import (
	request "GoRepo/source/assignment/codeHelpers"
	"GoRepo/source/assignment/services"

	"github.com/gin-gonic/gin"
)

type Audience_Controller interface {
	Get(c *gin.Context)
	GetFavourites(c *gin.Context)
	GetAudience(c *gin.Context)
	SetFavourite(c *gin.Context)
	UnsetFavourite(c *gin.Context)
	Delete(c *gin.Context)
}

type audience_controller struct {
	audienceService services.AudienceService
}

func NewAudienceController(audienceService services.AudienceService) Audience_Controller {
	return &audience_controller{audienceService: audienceService}
}

func (a *audience_controller) Get(c *gin.Context) {
	userId := c.Param("userId")
	Audiences, err := a.audienceService.GetAllAudiences(userId)

	if err != nil {
		request.BadRequest(c)
		return
	}

	request.OK(c, Audiences)
}

func (a *audience_controller) GetFavourites(c *gin.Context) {
	userId := c.Param("userId")
	audiences, err := a.audienceService.GetFavouriteAudiences(userId)

	if err != nil {
		request.BadRequest(c)
		return
	}

	request.OK(c, audiences)
}

func (a *audience_controller) GetAudience(c *gin.Context) {
	audienceId := c.Param("audienceId")
	audience, err := a.audienceService.GetAudienceById(audienceId)

	if err != nil {
		request.BadRequest(c)
		return
	}

	request.OK(c, audience)
}

func (a *audience_controller) SetFavourite(c *gin.Context) {
	audienceId := c.Param("audienceId")
	result, err := a.audienceService.SetFavouriteAudience(audienceId)

	if err != nil {
		request.BadRequest(c)
		return
	}

	request.OK(c, result)
}

func (a *audience_controller) UnsetFavourite(c *gin.Context) {
	audienceId := c.Param("audienceId")
	result, err := a.audienceService.UnsetFavouriteAudience(audienceId)

	if err != nil {
		request.BadRequest(c)
		return
	}

	request.OK(c, result)
}

func (a *audience_controller) Delete(c *gin.Context) {
	audienceId := c.Param("audienceId")
	result, err := a.audienceService.DeleteAudience(audienceId)

	if err != nil {
		request.BadRequest(c)
		return
	}

	request.OK(c, result)
}
