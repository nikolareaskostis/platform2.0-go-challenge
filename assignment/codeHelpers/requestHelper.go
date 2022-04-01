package codeHelpers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func OK(c *gin.Context, obj interface{}) {
	c.JSON(http.StatusOK, obj)
}

func ActionOK(c *gin.Context) {
	c.Status(http.StatusOK)
}

func BadRequest(c *gin.Context) {
	c.Status(http.StatusBadRequest)
	c.Abort()
}

func Unauthorized(c *gin.Context) {
	c.Status(http.StatusUnauthorized)
	c.Abort()
}

func InternalServerError(c *gin.Context) {
	c.Status(http.StatusInternalServerError)
	c.Abort()
}
