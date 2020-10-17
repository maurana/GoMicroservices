package routes

import (
	"net/http"
	"github.com/gin-gonic/gin"
	user "service_persistance1/controllers"
)

func Routes(){
	router := gin.Default()
	api := router.Group("/api") 
	{
		api.GET("/users", user.ReadAll)
		api.GET("/users:id", user.Read)
		api.POST("/users", user.Create)
		api.GET("/users/:id",user.Update)
		api.DELETE("/users/:id", user.Delete)
	}

	router.NoRoute(func(c *gin.Context) {
		c.AbortWithStatus(http.StatusNotFound)
	})

	router.Run(":8001")
}