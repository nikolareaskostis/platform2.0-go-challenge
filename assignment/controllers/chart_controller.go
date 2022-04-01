package api

import (
	request "GoRepo/source/assignment/codeHelpers"
	"GoRepo/source/assignment/services"
	"github.com/gin-gonic/gin"
)

type Chart_Controller interface {
	Get(c *gin.Context)
	GetFavourites(c *gin.Context)
	GetChart(c *gin.Context)
	SetFavourite(c *gin.Context)
	UnsetFavourite(c *gin.Context)
	Delete(c *gin.Context)
}

type chart_controller struct {
	chartService services.ChartService
}

func NewChartsController(chartService services.ChartService) Chart_Controller {
	return &chart_controller{chartService: chartService}
}

func (chart_c *chart_controller) Get(c *gin.Context) {
	userId := c.Param("userId")
	Charts, err := chart_c.chartService.GetAllCharts(userId)

	if err != nil {
		request.BadRequest(c)
		return
	}

	request.OK(c, Charts)
}

func (chart_c *chart_controller) GetFavourites(c *gin.Context) {
	userId := c.Param("userId")
	Charts, err := chart_c.chartService.GetFavouriteCharts(userId)

	if err != nil {
		request.BadRequest(c)
		return
	}

	request.OK(c, Charts)
}

func (chart_c *chart_controller) GetChart(c *gin.Context) {
	chartId := c.Param("chartId")
	Chart, err := chart_c.chartService.GetChartById(chartId)

	if err != nil {
		request.BadRequest(c)
		return
	}

	request.OK(c, Chart)
}

func (chart_c *chart_controller) SetFavourite(c *gin.Context) {
	chartId := c.Param("chartId")
	res, err := chart_c.chartService.SetFavouriteChart(chartId)

	if err != nil {
		request.BadRequest(c)
		return
	}

	request.OK(c, res)
}

func (chart_c *chart_controller) UnsetFavourite(c *gin.Context) {
	chartId := c.Param("chartId")
	res, err := chart_c.chartService.UnsetFavouriteChart(chartId)

	if err != nil {
		request.BadRequest(c)
		return
	}

	request.OK(c, res)
}

func (chart_c *chart_controller) Delete(c *gin.Context) {
	chartId := c.Param("chartId")
	res, err := chart_c.chartService.DeleteChart(chartId)

	if err != nil {
		request.BadRequest(c)
		return
	}

	request.OK(c, res)
}
