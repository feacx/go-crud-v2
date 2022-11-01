package requests

import (
	"github.com/feacx/study/pkg/response"
	"github.com/gin-gonic/gin"
)

type ValidatorFunc func(interface{}, *gin.Context) map[string][]string

func Validate(ctx *gin.Context, obj interface{}, handler ValidatorFunc) bool {
	if err := ctx.ShouldBind(obj); err != nil {
		response.BadRequest(ctx, err, "请求解析错误，请确认格式是否正确。")
	}
	return true
}
