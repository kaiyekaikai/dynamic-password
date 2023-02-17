package routes

import (
	"dynamic-password/user/controllers"
	"github.com/gin-gonic/gin"
)

func Init() *gin.Engine {
	r := gin.Default()

	user := r.Group("/user")
	{
		user.POST("/register", controllers.Register)
		user.POST("/send_code", controllers.SendEmailCode)
	}

	admin := r.Group("/admin")
	{
		admin.PUT("/generate_rsakey", controllers.GenerateRsaKey)
		admin.GET("/test_get_text", controllers.TestGetCiphertext)
	}

	return r
}
