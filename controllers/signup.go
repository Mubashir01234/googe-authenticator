package controllers

import (
	"authenticator/models"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func (ac *Controller) SignUp(ctx *gin.Context) {
	var body *models.RegisterUserInput

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	newUser := models.User{
		Name:     body.Name,
		Email:    strings.ToLower(body.Email),
		Password: body.Password,
	}

	result := ac.DB.Create(&newUser)

	if result.Error != nil && strings.Contains(result.Error.Error(), "duplicate key value violates unique") {
		ctx.JSON(http.StatusConflict, gin.H{"status": "fail", "message": "Email already exist, please use another email address"})
		return
	} else if result.Error != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": result.Error.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "message": "registered successfully, please login"})
}
