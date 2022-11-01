package routes

import (
	"github.com/feacx/study/app/http/controllers/api"
	"github.com/feacx/study/app/http/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterAPIRoutes(r *gin.Engine) {
	r.GET("/user", middleware.Logger(), api.GetUser)
	r.POST("/user", middleware.Logger(), api.CreateUser)
}
