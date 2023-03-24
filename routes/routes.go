package routes

import (
	"authenticator/controllers"

	"github.com/gin-gonic/gin"
)

type RouteController struct {
	controller controllers.Controller
}

func NewRouteController(authController controllers.Controller) RouteController {
	return RouteController{authController}
}

func (c *RouteController) AuthRoute(rg *gin.RouterGroup) {
	// auth routes
	auth := rg.Group("auth")
	auth.POST("/register", c.controller.SignUp)
	auth.POST("/login", c.controller.Login)

	// OTP routes
	otp := auth.Group("otp")
	otp.POST("/generate", c.controller.GenerateOTP)
	otp.POST("/verify", c.controller.VerifyOTP)
	otp.POST("/validate", c.controller.ValidateOTP)
	otp.POST("/disable", c.controller.DisableOTP)
}
