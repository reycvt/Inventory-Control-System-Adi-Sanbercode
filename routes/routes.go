package routes

import (
	"inverntory-adi-sanbercode/controllers"
	"inverntory-adi-sanbercode/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {

	r.POST("/login", controllers.Login)
	r.POST("/register", controllers.Register)

	protected := r.Group("/")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.GET("/masterbarang", controllers.GetALLBarang)
		protected.GET("/masterbarang/:id", controllers.GetBarang)
		protected.POST("/masterbarang", controllers.CreateBarang)
		protected.PUT("/masterbarang/:id", controllers.UpdateBarang)
		protected.DELETE("/masterbarang/:id", controllers.DeleteBarang)

		protected.GET("/masterbarang/stock", controllers.GetAllStock)
		protected.GET("/masterbarang/:id/stock", controllers.GetStock)
		protected.POST("/masterbarang/:id/stock", controllers.CreateStockById)
		protected.PUT("/masterbarang/:id/stock", controllers.UpdateStock)

		protected.GET("/masterbarang/transaksi", controllers.GetALLTransaksi)
		protected.GET("/masterbarang/:id/transaksi", controllers.GetTransaksi)
		protected.POST("/masterbarang/:id/transaksi", controllers.CreateTransaksiById)
		protected.PUT("/masterbarang/:id/transaksi", controllers.UpdateTransaksi)
	}
}
