package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/javice/bazar-go/backend/internal/handlers"
)

func SetupRoutes(r *gin.Engine) {
	api := r.Group("/api/v1")
	{
		// Productos
		api.GET("/products", handlers.GetAllProducts)
		api.GET("/products/:id", handlers.GetProductByID)
		api.POST("/products", handlers.CreateProduct)
		api.PUT("/products/:id", handlers.UpdateProduct)
		api.DELETE("/products/:id", handlers.DeleteProduct)

		// Clientes
		api.GET("/clients", handlers.GetAllClients)
		api.GET("/clients/:id", handlers.GetClientByID)
		api.POST("/clients", handlers.CreateClient)
		api.PUT("/clients/:id", handlers.UpdateClient)
		api.DELETE("/clients/:id", handlers.DeleteClient)

		// Ventas
		api.GET("/sales", handlers.GetAllSales)
		api.GET("/sales/:id", handlers.GetSaleByID)
		api.POST("/sales", handlers.CreateSale)
		api.PUT("/sales/:id", handlers.UpdateSale)
		api.DELETE("/sales/:id", handlers.DeleteSale)
	}
}