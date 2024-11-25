package controller

import (
	"strconv"
	"test-mnc/config"
	"test-mnc/entity"
	"test-mnc/logger"
	"test-mnc/shared/common"
	"test-mnc/usecase"

	"github.com/gin-gonic/gin"
)

type productController struct {
	uc  usecase.ProductUsecase
	rg  *gin.RouterGroup
	log *logger.Logger
}

func (c *productController) CreateProductHandler(ctx *gin.Context) {
	var product entity.Product

	c.log.Info("Starting to binding a payload", nil)
	if err := ctx.ShouldBindJSON(&product); err != nil {
		c.log.Error("Invalid payload for create product", err)
		common.SendErrorResponse(ctx, 400, err.Error())
		return
	}

	c.log.Info("Starting to create a new product", product)
	createdProduct, err := c.uc.CreateProduct(product)
	if err != nil {
		c.log.Error("Failed to create a new product", err)
		common.SendErrorResponse(ctx, 500, err.Error())
		return
	}

	c.log.Info("Successfully created a new product", createdProduct)
	common.SendCreateResponse(ctx, createdProduct)
}

// DeleteProductHandler deletes a product by its ID.
func (c *productController) DeleteProductHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	idInt, err := strconv.Atoi(id)

	if err != nil {
		c.log.Error("Invalid product ID", err)
		common.SendErrorResponse(ctx, 400, err.Error())
		return
	}

	c.log.Info("Starting to delete a product", idInt)
	err = c.uc.DeleteProduct(idInt)
	if err != nil {
		c.log.Error("Failed to delete a product", err)
		common.SendErrorResponse(ctx, 500, err.Error())
		return
	}

	c.log.Info("Successfully deleted a product", id)
	common.SendDeleteResponse(ctx, "Successfully deleted product")
}

// UpdateProductHandler updates a product by its ID.
func (c *productController) UpdateProductHandler(ctx *gin.Context) {
	var product entity.Product

	id := ctx.Param("id")
	idInt, err := strconv.Atoi(id)

	if err != nil {
		c.log.Error("Invalid product ID", err)
		common.SendErrorResponse(ctx, 400, err.Error())
		return
	}

	c.log.Info("Starting to binding a payload", nil)
	if err := ctx.ShouldBindJSON(&product); err != nil {
		c.log.Error("Invalid payload for update product", err)
		common.SendErrorResponse(ctx, 400, err.Error())
		return
	}

	product.ID = idInt

	c.log.Info("Starting to update a product", product)
	updatedProduct, err := c.uc.UpdateProduct(product)
	if err != nil {
		c.log.Error("Failed to update a product", err)
		common.SendErrorResponse(ctx, 500, err.Error())
		return
	}

	c.log.Info("Successfully updated a product", updatedProduct)
	common.SendSingleResponse(ctx, updatedProduct, "Successfully updated product")
}

// GetAllProductsHandler gets all products.
func (c *productController) GetAllProductsHandler(ctx *gin.Context) {
	c.log.Info("Starting to get all products", nil)
	products, err := c.uc.GetAllProducts()
	if err != nil {
		c.log.Error("Failed to get all products", err)
		common.SendErrorResponse(ctx, 500, err.Error())
		return
	}

	c.log.Info("Successfully retrieved all products", products)
	common.SendSingleResponse(ctx, products, "Successfully retrieved all products")
}

// GetProductByIDHandler gets a product by its ID.
func (c *productController) GetProductByIDHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	idInt, err := strconv.Atoi(id)

	if err != nil {
		c.log.Error("Invalid product ID", err)
		common.SendErrorResponse(ctx, 400, err.Error())
		return
	}

	c.log.Info("Starting to get a product", idInt)
	product, err := c.uc.GetProductById(idInt)
	if err != nil {
		c.log.Error("Failed to get a product", err)
		common.SendErrorResponse(ctx, 500, err.Error())
		return
	}

	c.log.Info("Successfully retrieved a product", product)
	common.SendSingleResponse(ctx, product, "Successfully retrieved product")
}

func (c *productController) Route() {
	c.rg.POST(config.PostProduct, c.CreateProductHandler)
	c.rg.DELETE(config.DelProduct, c.DeleteProductHandler)
	c.rg.PUT(config.PutProduct, c.UpdateProductHandler)
	c.rg.GET(config.GetAllProduct, c.GetAllProductsHandler)
	c.rg.GET(config.GetProductId, c.GetProductByIDHandler)
}

func NewProductController(uc usecase.ProductUsecase, rg *gin.RouterGroup, log *logger.Logger) *productController {
	return &productController{uc: uc, rg: rg, log: log}
}
