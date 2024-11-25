package controller

import (
	"fmt"
	"strconv"
	"test-mnc/config"
	"test-mnc/delivery/middleware"
	"test-mnc/entity"
	"test-mnc/logger"
	"test-mnc/shared/common"
	"test-mnc/usecase"

	"github.com/gin-gonic/gin"
)

type topupController struct {
	topupUc        usecase.TopupUsecase
	productUc      usecase.ProductUsecase
	authMiddleware middleware.AuthMiddleware
	rg             *gin.RouterGroup
	log            *logger.Logger
}

// CreateTopupHandler creates a new topup.
func (t *topupController) CreateTopupHandler(ctx *gin.Context) {
	var topup entity.Topup

	t.log.Info("Starting to binding a payload", nil)
	if err := ctx.ShouldBindJSON(&topup); err != nil {
		t.log.Error("Failed to bind payload", err)
		common.SendErrorResponse(ctx, 400, err.Error())
		return
	}

	t.log.Info("Starting to create a new topup", topup)
	createdTopup, err := t.topupUc.CreateTopup(topup)
	if err != nil {
		t.log.Error("Failed to create topup", err)
		common.SendErrorResponse(ctx, 500, err.Error())
		return
	}

	getProduct, err := t.productUc.GetProductById(createdTopup.ProductID)
	if err != nil {
		t.log.Error("Failed to get product", err)
		common.SendErrorResponse(ctx, 500, err.Error())
		return
	}

	confirmTopup := &entity.ConfirmTopup{
		TopupID:       createdTopup.ID,
		Amount:        getProduct.Nominal,
		Price:         getProduct.Price,
		PaymentMethod: createdTopup.PaymentMethod,
		PaymentStatus: "Done",
	}

	reply := fmt.Sprintf("Silahkan transfer sebesar %.2f ke bank %s, rekening 123456789 a/n PT EMONEY INDONESIA. Silahkan konfirmasi : {topup_id : %d, amount : %.2f, price : %.2f, payment_method : %s, payment_status : %s} pada link : 'http://localhost:8080/api/v1/topup/callback' jika sudah melakukan transfer", confirmTopup.Price, confirmTopup.PaymentMethod, confirmTopup.TopupID, confirmTopup.Amount, confirmTopup.Price, confirmTopup.PaymentMethod, confirmTopup.PaymentStatus)

	t.log.Info("Successfully created topup", createdTopup)
	common.SendCreateResponse(ctx, reply)
}

func (t *topupController) GetTopupById(ctx *gin.Context) {
	id := ctx.Param("id")
	idInt, err := strconv.Atoi(id)

	if err != nil {
		t.log.Error("Invalid topup ID", err)
		common.SendErrorResponse(ctx, 400, err.Error())
		return
	}

	t.log.Info("Starting to get topup by ID", id)
	topup, err := t.topupUc.GetTopupById(idInt)
	if err != nil {
		t.log.Error("Failed to get topup", err)
		common.SendErrorResponse(ctx, 500, err.Error())
		return
	}

	t.log.Info("Successfully fetched topup", topup)
	common.SendSingleResponse(ctx, topup, "Topup ditemukan")
}

func (t *topupController) PaymentCallbackHandler(ctx *gin.Context) {
	var confirmPayment entity.ConfirmTopup

	t.log.Info("Starting to binding a payload", nil)
	if err := ctx.ShouldBindJSON(&confirmPayment); err != nil {
		t.log.Error("Failed to bind payload", err)
		common.SendErrorResponse(ctx, 400, err.Error())
		return
	}

	t.log.Info("Starting to update the topup data", confirmPayment)
	err := t.topupUc.UpdateBalanceAfterPayment(confirmPayment)
	if err != nil {
		t.log.Error("Failed to update topup", err)
		common.SendErrorResponse(ctx, 500, err.Error())
		return
	}

	t.log.Info("Successfully updated topup", nil)
	common.SendSingleResponse(ctx, confirmPayment, "Topup berhasil")
}

func (t *topupController) Route() {
	t.rg.POST(config.PostTopup, t.authMiddleware.Middleware, t.CreateTopupHandler)
	t.rg.POST(config.PostTopupCallback, t.authMiddleware.Middleware, t.PaymentCallbackHandler)
}

func NewTopupController(topupUc usecase.TopupUsecase, productUc usecase.ProductUsecase, authMiddleware middleware.AuthMiddleware, rg *gin.RouterGroup, log *logger.Logger) *topupController {
	return &topupController{topupUc: topupUc, productUc: productUc, authMiddleware: authMiddleware, rg: rg, log: log}
}
