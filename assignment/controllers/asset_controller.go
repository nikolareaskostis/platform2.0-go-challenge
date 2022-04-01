package api

import (
	request "GoRepo/source/assignment/codeHelpers"
	"GoRepo/source/assignment/services"
	"github.com/gin-gonic/gin"
)

type Asset_Controller interface {
	Get(c *gin.Context)
	GetFavourites(c *gin.Context)
}

type asset_controller struct {
	assetService services.AssetService
}

func NewAssetsController(assetService services.AssetService) Asset_Controller {
	return &asset_controller{assetService: assetService}
}

func (a *asset_controller) Get(c *gin.Context) {
	userId := c.Param("userId")
	assets, err := a.assetService.GetAllAssets(userId)

	if err != nil {
		request.BadRequest(c)
		return
	}

	request.OK(c, assets)
}

func (a *asset_controller) GetFavourites(c *gin.Context) {
	userId := c.Param("userId")
	assets, err := a.assetService.GetFavouriteAssets(userId)

	if err != nil {
		request.BadRequest(c)
		return
	}

	request.OK(c, assets)
}
