package common

import (
	"net/http"
	"test-mnc/shared/model"

	"github.com/gin-gonic/gin"
)

func SendCreateResponse(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusCreated, &model.SingleResponse{
		Status: model.Status{
			Code:    http.StatusCreated,
			Message: "Created",
		},
		Data: data,
	})
}

func SendSingleResponse(ctx *gin.Context, data interface{}, message string) {
	ctx.JSON(http.StatusOK, &model.SingleResponse{
		Status: model.Status{
			Code:    http.StatusOK,
			Message: message,
		},
		Data: data,
	})
}

func SendDeleteResponse(ctx *gin.Context, message string) {
	ctx.JSON(http.StatusNoContent, &model.Status{
		Code:    http.StatusNoContent,
		Message: message,
	})
}

func SendErrorResponse(ctx *gin.Context, code int, message string) {
	ctx.AbortWithStatusJSON(code, &model.Status{
		Code:    code,
		Message: message,
	})
}
