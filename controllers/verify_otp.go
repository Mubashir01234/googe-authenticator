package controllers

import (
	"authenticator/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pquerna/otp/totp"
)

func (ac *Controller) VerifyOTP(ctx *gin.Context) {
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

	valid := totp.Validate(body.Token, user.Otp_secret)
	if !valid {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "token is invalid"})
		return
	}

	dataToUpdate := models.User{
		OTPEnabled:  true,
		OTPVerified: true,
	}

	ac.DB.Model(&user).Updates(dataToUpdate)

	userResponse := gin.H{
		"id":          user.ID.String(),
		"name":        user.Name,
		"email":       user.Email,
		"otp_enabled": user.OTPEnabled,
	}
	ctx.JSON(http.StatusOK, gin.H{"otp_verified": true, "user": userResponse})
}
