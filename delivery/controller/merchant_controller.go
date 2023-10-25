package controller

import (
	"api-payment/delivery/middleware"
	"api-payment/model"
	"api-payment/usecase"
	"api-payment/utils/common"
	"net/http"

	"github.com/gin-gonic/gin"
)

type MerchantController struct {
	router  *gin.Engine
	useCase usecase.MerchantUseCase
}

func (m *MerchantController) createHandler(c *gin.Context) {
	var merchant model.Merchant
	if err := c.ShouldBindJSON(&merchant); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	merchant.Id = common.GenerateUUID()
	err := m.useCase.RegisterNewMerchant(merchant)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusCreated,
		"message": "Success Create New Merchant",
		"data":    merchant,
	})
}

func (m *MerchantController) listHandler(c *gin.Context) {
	merchants, err := m.useCase.FindAllMerchantList()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	status := map[string]any{
		"code":        200,
		"description": "Get All Data Successfully",
	}
	c.JSON(http.StatusOK, gin.H{
		"status": status,
		"data":   merchants,
	})
}

func NewMerchantController(router *gin.Engine, merchantUseCase usecase.MerchantUseCase) {
	ctr := &MerchantController{
		router:  router,
		useCase: merchantUseCase,
	}

	routerGroup := ctr.router.Group("/api/v1",middleware.AuthMiddleware())
	routerGroup.POST("/merchant", ctr.createHandler)
	routerGroup.GET("/merchant", ctr.listHandler)
}