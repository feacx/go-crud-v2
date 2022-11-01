package cmd

import (
	"fmt"
	"os"

	"github.com/feacx/study/bootstrap"
	"github.com/gin-gonic/gin"
)

func Runer() {
	bootstrap.LoadEnv()
	gin.SetMode("debug")
	router := gin.New()
	bootstrap.SetupRoute(router)
	bootstrap.SetupDB()
	bootstrap.SetupLogger()
	err := router.Run(os.Getenv("APP_PORT"))
	if err != nil {
		fmt.Println("err", err)
	}
}
