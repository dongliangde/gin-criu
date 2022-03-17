package routes

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go-criu/api"
	"go-criu/cmd"
	"go-criu/middleware"
)

func InitRouter() {
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(middleware.Cors())
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.POST("/dump", api.Dump)
	r.POST("/restore", api.Restore)
	_ = r.Run(":" + cmd.Port)
}
