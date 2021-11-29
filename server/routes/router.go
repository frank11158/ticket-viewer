package routes

import (
	"server/controllers"
	v1 "server/controllers/v1"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func InitRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	config := cors.DefaultConfig()
	config.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"}
	config.AllowHeaders = []string{"Content-Type", "Authorization"}
	config.AllowOrigins = viper.GetStringSlice("ALLOWED_ORIGINS")
	config.AllowCredentials = true
	router.Use(cors.New(config))

	router.GET("/heartBeat", controllers.HeartBeat)

	apiv1 := router.Group("api/v1")
	{
		apiv1.GET("/tickets", v1.GetTickets)
		apiv1.GET("/users/:user_id", v1.GetUserInfo)
	}

	return router
}
