package controller

import (
	"test-mnc/config"
	"test-mnc/entity"
	"test-mnc/entity/dto"
	"test-mnc/logger"
	"test-mnc/shared/common"
	"test-mnc/usecase"

	"github.com/gin-gonic/gin"
)

type customerController struct {
	uc  usecase.AuthUsecase
	rg  *gin.RouterGroup
	log *logger.Logger
}

func (c *customerController) registerHandler(ctx *gin.Context) {
	var payload entity.Customer

	c.log.Info("Starting to binding a payload", nil)
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		c.log.Error("Invalid payload for register", err)
		common.SendErrorResponse(ctx, 400, err.Error())
		return
	}

	payloadMap := map[string]any{"email": payload.Email, "password": payload.Password}
	c.log.Info("Starting to validation required fields a new customer", payloadMap)
	if payload.Name == "" || payload.Email == "" || payload.Password == "" {
		c.log.Error("Missing required fields for a new customer", payload)
		common.SendErrorResponse(ctx, 400, "Missing required fields")
		return
	}

	c.log.Info("Starting to create a new customer", payloadMap)
	data, err := c.uc.Register(payload)
	if err != nil {
		c.log.Error("Failed to create a new customer", err)
		common.SendErrorResponse(ctx, 500, err.Error())
		return
	}

	c.log.Info("Successfully created a new customer", data)
	common.SendCreateResponse(ctx, data)
}

func (c *customerController) loginHandler(ctx *gin.Context) {
	var payload dto.AuthRequestDto

	c.log.Info("Starting to binding a payload", nil)
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		c.log.Error("Invalid payload for login", err)
		common.SendErrorResponse(ctx, 400, err.Error())
		return
	}

	payloadMap := map[string]any{"email": payload.Email, "password": payload.Password}
	c.log.Info("Starting to validation required fields for login", payloadMap)
	if payload.Email == "" || payload.Password == "" {
		c.log.Error("Missing required fields for login", payloadMap)
		common.SendErrorResponse(ctx, 400, "Missing required fields")
		return
	}

	c.log.Info("Starting to login a customer", payloadMap)
	data, err := c.uc.Login(payload)
	if err != nil {
		c.log.Error("Failed to login a customer", err)
		common.SendErrorResponse(ctx, 500, err.Error())
		return
	}

	c.log.Info("Successfully logged in a customer", data)
	common.SendSingleResponse(ctx, data, "OK")
}

func (c *customerController) Route() {
	c.rg.POST(config.PostCustomerRegister, c.registerHandler)
	c.rg.POST(config.PostCustomerLogin, c.loginHandler)
}

func NewCustomerController(uc usecase.AuthUsecase, rg *gin.RouterGroup, log *logger.Logger) *customerController {
	return &customerController{uc: uc, rg: rg, log: log}
}
