package controller

import (
	"api-payment/delivery/middleware"
	"api-payment/model"
	"api-payment/usecase"
	"api-payment/utils/common"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PaymentController struct {
	router  *gin.Engine
	useCase usecase.PaymentUseCase
}

func (p *PaymentController) createHandler(c *gin.Context) {
	var payment model.Payment
	if err := c.ShouldBindJSON(&payment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	payment.Id = common.GenerateUUID()
	err := p.useCase.RegisterNewPayment(payment)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusCreated,
		"message": "Success Create New Payment",
		"data":    payment,
	})
}

func (p *PaymentController) listHandler(c *gin.Context) {
	payments, err := p.useCase.FindAllPaymentList()
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
		"data":   payments,
	})
}

func NewPaymentController(router *gin.Engine, PaymentUseCase usecase.PaymentUseCase) {
	ctr := &PaymentController{
		router:  router,
		useCase: PaymentUseCase,
	}

	routerGroup := ctr.router.Group("/api/v1",middleware.AuthMiddleware())
	routerGroup.POST("/payment", ctr.createHandler)
	routerGroup.GET("/payment", ctr.listHandler)
}