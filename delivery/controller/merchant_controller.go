package controller

import (
	"test-mnc/config"
	"test-mnc/entity"
	"test-mnc/logger"
	"test-mnc/shared/common"
	"test-mnc/usecase"

	"github.com/gin-gonic/gin"
)

type merchantController struct {
	uc  usecase.MerchantUsecase
	rg  *gin.RouterGroup
	log *logger.Logger
}

// CreateMerchantHandler creates a new merchant.
func (c *merchantController) CreateMerchantHandler(ctx *gin.Context) {
	var merchant entity.Merchant

	c.log.Info("Starting to binding a payload", nil)
	if err := ctx.ShouldBindJSON(&merchant); err != nil {
		c.log.Error("Failed to bind payload", err)
		common.SendErrorResponse(ctx, 400, err.Error())
		return
	}

	c.log.Info("Starting to create a new merchant", merchant)
	createdMerchant, err := c.uc.CreateMerchant(merchant)
	if err != nil {
		c.log.Error("Failed to create merchant", err)
		common.SendErrorResponse(ctx, 500, err.Error())
		return
	}

	c.log.Info("Successfully created merchant", createdMerchant)
	common.SendCreateResponse(ctx, createdMerchant)
}

func (c *merchantController) Route() {
	c.rg.POST(config.PostMerchant, c.CreateMerchantHandler)
}

func NewMerchantController(uc usecase.MerchantUsecase, rg *gin.RouterGroup, log *logger.Logger) *merchantController {
	return &merchantController{uc: uc, rg: rg, log: log}
}
