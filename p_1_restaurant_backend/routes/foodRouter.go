package routes

import (
	"github.com/gin-gonic/gin"
	"restaurant_backend/controllers"
)

func FoodRoutes(router *gin.Engine) {
	router.GET("/foods", controllers.GetFoods())
	router.GET("/foods/:food_id", controllers.GetFood())
	router.POST("/foods", controllers.CreateFood())
	router.PATCH("/foods/:food_id", controllers.UpdateFood())
}
