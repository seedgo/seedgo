package seedgo

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Response(ctx *gin.Context, httpStatus int, code int, data interface{}, msg string) {
	ctx.JSON(httpStatus, gin.H{"code": code, "data": data, "msg": msg})
}

func Success(ctx *gin.Context, data interface{}) {
	Response(ctx, http.StatusOK, http.StatusOK, data, "success")
}

func ParamValidateFail(ctx *gin.Context, msg string) {
	Response(ctx, http.StatusOK, http.StatusUnprocessableEntity, nil, msg)
}

func Fail(ctx *gin.Context, msg string, errcode int) {
	Response(ctx, http.StatusOK, errcode, nil, msg)
}

func ValidateFail(ctx *gin.Context, data map[string][]string) {
	Response(ctx, http.StatusUnprocessableEntity, 422, data, "params validation error")
}

func FailWithErr(ctx *gin.Context, err BusErr) {
	Response(ctx, http.StatusOK, err.Code, nil, err.Message)
}
