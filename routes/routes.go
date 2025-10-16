package routes

import (
	"test/pract.com/controllers"
	"github.com/gin-gonic/gin"
	
)


func UserRoutes(incomingRoutes *gin.Engine){
	incomingRoutes.POST("/user/signup", controllers.Signup())
	incomingRoutes.POST("/user/login", controllers.Login())
	incomingRoutes.POST("/admin/addproduct", controllers.ProductViewerAdmin())
	incomingRoutes.GET("/user/search", controllers.SearchProduct())
	incomingRoutes.GET("/users/productview", controllers.SearchProductByQuery())
	


}