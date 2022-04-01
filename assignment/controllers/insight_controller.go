package api

import (
	request "GoRepo/source/assignment/codeHelpers"
	"GoRepo/source/assignment/services"
	"github.com/gin-gonic/gin"
)

type Insight_Controller interface {
	Get(c *gin.Context)
	GetFavourites(c *gin.Context)
	GetInsight(c *gin.Context)
	SetFavourite(c *gin.Context)
	UnsetFavourite(c *gin.Context)
	Delete(c *gin.Context)
}

type insight_controller struct {
	insightService services.InsightService
}

func NewInsightController(insightService services.InsightService) Insight_Controller {
	return &insight_controller{insightService: insightService}
}

func (insight_c *insight_controller) Get(c *gin.Context) {
	userId := c.Param("userId")
	Insights, err := insight_c.insightService.GetAllInsights(userId)

	if err != nil {
		request.BadRequest(c)
		return
	}

	request.OK(c, Insights)
}

func (insight_c *insight_controller) GetFavourites(c *gin.Context) {
	userId := c.Param("userId")
	Insights, err := insight_c.insightService.GetFavouriteInsights(userId)

	if err != nil {
		request.BadRequest(c)
		return
	}

	request.OK(c, Insights)
}

func (insight_c *insight_controller) GetInsight(c *gin.Context) {
	insightId := c.Param("insightId")
	Insight, err := insight_c.insightService.GetInsightById(insightId)

	if err != nil {
		request.BadRequest(c)
		return
	}

	request.OK(c, Insight)
}

func (insight_c *insight_controller) SetFavourite(c *gin.Context) {
	insightId := c.Param("insightId")
	res, err := insight_c.insightService.SetFavouriteInsight(insightId)

	if err != nil {
		request.BadRequest(c)
		return
	}

	request.OK(c, res)
}

func (insight_c *insight_controller) UnsetFavourite(c *gin.Context) {
	insightId := c.Param("insightId")
	res, err := insight_c.insightService.UnsetFavouriteInsight(insightId)

	if err != nil {
		request.BadRequest(c)
		return
	}

	request.OK(c, res)
}

func (insight_c *insight_controller) Delete(c *gin.Context) {
	insightId := c.Param("insightId")
	res, err := insight_c.insightService.DeleteInsight(insightId)

	if err != nil {
		request.BadRequest(c)
		return
	}

	request.OK(c, res)
}
