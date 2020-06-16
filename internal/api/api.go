package api

import (
	"multi-proxy/internal/controllers"

	"github.com/gin-gonic/gin"
)

func StartAPI() {
	router := gin.Default()
	router.POST("/deployments", controllers.HandlePostDeployments)
	router.Run(":8080")
}
