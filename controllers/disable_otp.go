package controllers

import (
	"authenticator/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (ac *Controller) DisableOTP(ctx *gin.Context) {
	var body *models.OTPInput

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	var user models.User
	result := ac.DB.First(&user, "id = ?", body.UserId)
	if result.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "user doesn't exist"})
		return
	}

	user.OTPEnabled = false
	ac.DB.Save(&user)

	userResponse := gin.H{
		"id":          user.ID.String(),
		"name":        user.Name,
		"email":       user.Email,
		"otp_enabled": user.OTPEnabled,
	}
	ctx.JSON(http.StatusOK, gin.H{"otp_disabled": true, "user": userResponse})
}
