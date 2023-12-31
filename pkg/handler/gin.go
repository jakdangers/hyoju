package handler

import (
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/fx"
	"pixelix/docs"
)

var Module = fx.Module("handler", fx.Provide(NewServeMux))

func NewServeMux() *gin.Engine {
	r := gin.Default()

	docs.SwaggerInfo.BasePath = "/api/v1"

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	return r
}
