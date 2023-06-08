package router

import (
	"github.com/gin-gonic/gin"
	"github.com/ropel12/URL-Shortner/handlers"
)

func NewRouter(urlHandler handlers.URLHandler) *gin.Engine {
	r := gin.Default()

	r.GET("/:path", urlHandler.Get)
	r.POST("/", urlHandler.Create)
	r.POST("/custom", urlHandler.CreateCustom)
	return r
}
