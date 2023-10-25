package controller

import (
	"api-payment/model"
	"api-payment/usecase"
	"api-payment/utils/common"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	router *gin.Engine
	userUc usecase.UserUseCase
}

func (u *UserController) createHandler(c *gin.Context) {
	var user model.UserCredential
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error" : err.Error(),
		})
		return
	}

	user.Id = common.GenerateUUID()
	if err := u.userUc.RegisterNewUser(user); err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error" : err.Error(),
		})
		return
	}

	userResponse := map[string]any{
		"id" : user.Id,
		"username" : user.Username,
	}

	c.JSON(http.StatusOK,userResponse)
}

func NewUserController(r *gin.Engine, userUseCase usecase.UserUseCase){
	controller := UserController{
		router: r,
		userUc: userUseCase,
	}
	rg := r.Group("/api/v1")
	rg.POST("/register",controller.createHandler)
}