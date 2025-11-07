package main

import (
	"expense-tracker-backend/config"
	"expense-tracker-backend/handlers"
	"expense-tracker-backend/middleware"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize database
	if err := config.InitDB(); err != nil {
		log.Fatal("Failed to initialize database:", err)
	}

	// Setup Gin router
	r := gin.Default()

	// CORS middleware
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})

	// Auth routes (public)
	auth := r.Group("/api/auth")
	{
		auth.POST("/register", handlers.Register)
		auth.POST("/login", handlers.Login)
	}

	// Expense routes (protected)
	expenses := r.Group("/api/expenses")
	expenses.Use(middleware.AuthMiddleware())
	{
		expenses.GET("", handlers.GetExpenses)
		expenses.POST("", handlers.CreateExpense)
		expenses.PUT("/:id", handlers.UpdateExpense)
		expenses.DELETE("/:id", handlers.DeleteExpense)
	}

	// Start server
	log.Println("Server starting on :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}

