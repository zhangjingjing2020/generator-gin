package routes

import (
	"<%= moduleName %>/app/controllers/app"
	"<%= moduleName %>/app/controllers/common"

	"<%= moduleName %>/app/middleware"
	"<%= moduleName %>/app/services"

	"github.com/gin-gonic/gin"
)

func SetApiGroupRoutes(router *gin.RouterGroup) {
	router.POST("/auth/register", app.Register)
	router.POST("/auth/login", app.Login)

	authRouter := router.Group("").Use(middleware.JWTAuth(services.AppGuardName))
	{
		authRouter.POST("/auth/info", app.Info)
		authRouter.POST("/auth/logout", app.Logout)
		authRouter.POST("/image_upload", common.ImageUpload)
	}
}
