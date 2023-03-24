package controllers

import (
	"authenticator/models"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func (ac *Controller) Login(ctx *gin.Context) {
	var body *models.LoginUserInput

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	var user models.User
	result := ac.DB.First(&user, "email = ?", strings.ToLower(body.Email))
	if result.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "invalid email or password"})
		return
	}

	userResponse := gin.H{
		"id":          user.ID.String(),
		"name":        user.Name,
		"email":       user.Email,
		"otp_enabled": user.OTPEnabled,
	}
	ctx.JSON(http.StatusOK, gin.H{"status": "success", "user": userResponse})
}
