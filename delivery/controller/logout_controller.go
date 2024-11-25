package controller

import (
	"strings"
	"test-mnc/config"
	"test-mnc/logger"
	"test-mnc/shared/common"
	"test-mnc/usecase"

	"github.com/gin-gonic/gin"
)

type blacklistController struct {
	uc  usecase.BlacklistUsecase
	rg  *gin.RouterGroup
	log *logger.Logger
}

func (b *blacklistController) Logout(ctx *gin.Context) {
	tokenString := ctx.GetHeader("Authorization")

	if tokenString == "" {
		b.log.Error("No token provided", nil)
		common.SendErrorResponse(ctx, 400, "No token provided")
		return
	}

	tokenString = strings.TrimPrefix(tokenString, "Bearer ")

	b.log.Info("Starting to add token to blacklist", tokenString)
	err := b.uc.AddTokenToBlacklist(tokenString)
	if err != nil {
		b.log.Error("Failed to add token to blacklist", err)
		common.SendErrorResponse(ctx, 500, err.Error())
		return
	}

	b.log.Info("Logged out successfully", nil)
	common.SendSingleResponse(ctx, nil, "Logged out successfully")
}

func (b *blacklistController) Route() {
	b.rg.POST(config.PostCustomerLogout, b.Logout)
}

func NewBlacklistController(uc usecase.BlacklistUsecase, rg *gin.RouterGroup, log *logger.Logger) *blacklistController {
	return &blacklistController{uc: uc, rg: rg, log: log}
}
