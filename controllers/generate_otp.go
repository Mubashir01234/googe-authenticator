package controllers

import (
	"authenticator/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pquerna/otp/totp"
)

func (ac *Controller) GenerateOTP(ctx *gin.Context) {
	var body *models.OTPInput

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	key, err := totp.Generate(totp.GenerateOpts{
		Issuer:      "digitalsstore.com",
		AccountName: "digitals@store.com",
		SecretSize:  15,
	})

	if err != nil {
		panic(err)
	}

	var user models.User
	result := ac.DB.First(&user, "id = ?", body.UserId)
	if result.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "invalid email or password"})
		return
	}

	dataToUpdate := models.User{
		Otp_secret:   key.Secret(),
		Otp_auth_url: key.URL(),
	}

	ac.DB.Model(&user).Updates(dataToUpdate)

	otpResponse := gin.H{
		"base32":      key.Secret(),
		"otpauth_url": key.URL(),
	}
	ctx.JSON(http.StatusOK, otpResponse)
}
