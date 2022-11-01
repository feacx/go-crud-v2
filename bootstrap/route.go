package bootstrap

import (
	"github.com/feacx/study/routes"
	"github.com/gin-gonic/gin"
)

func SetupRoute(router *gin.Engine) {
	routes.RegisterAPIRoutes(router)
}
