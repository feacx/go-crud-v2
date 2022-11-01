package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/feacx/study/app/http/models/user"
	"github.com/feacx/study/app/requests"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	BaseController
}

func GetUser(ctx *gin.Context) {
	id := ctx.Param("id")
	data := user.Get(id)
	ctx.JSON(http.StatusOK, gin.H{
		"message": "GetUser",
		"data":    data,
	})
}

func CreateUser(ctx *gin.Context) {
	log.Println("按时大按时大")
	user := requests.CreateUserRequest{}
	ctx.Bind(&user)
	fmt.Println(user)
}
