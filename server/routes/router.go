package routes

import (
	"server/controllers"
	v1 "server/controllers/v1"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.GET("/heartBeat", controllers.HeartBeat)

	apiv1 := router.Group("/v1")
	{
		apiv1.GET("/tickets", v1.GetTickets)
	}

	return router
}
