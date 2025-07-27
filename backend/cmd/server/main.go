package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"rent-help-backend/internal/config"
	"rent-help-backend/internal/handlers"
	"rent-help-backend/internal/middleware"
	"rent-help-backend/internal/services"
	"rent-help-backend/pkg/database"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	// Load config
	cfg := config.Load()

	// Connect to MongoDB
	db, err := database.Connect(cfg.MongoURI)
	if err != nil {
		log.Fatal("Failed to connect to MongoDB:", err)
	}
	defer database.Disconnect(db)

	// Initialize services
	userService := services.NewUserService(db)
	propertyService := services.NewPropertyService(db)
	bookingService := services.NewBookingService(db)

	// Initialize handlers
	userHandler := handlers.NewUserHandler(userService)
	propertyHandler := handlers.NewPropertyHandler(propertyService)
	bookingHandler := handlers.NewBookingHandler(bookingService)

	// Setup Gin router
	router := gin.Default()

	// CORS middleware
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Health check
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	// API routes
	api := router.Group("/api/v1")
	{
		// Auth routes
		auth := api.Group("/auth")
		{
			auth.POST("/register", userHandler.Register)
			auth.POST("/login", userHandler.Login)
			auth.POST("/refresh", userHandler.RefreshToken)
		}

		// Protected routes
		protected := api.Group("")
		protected.Use(middleware.AuthMiddleware())
		{
			// User routes
			users := protected.Group("/users")
			{
				users.GET("/profile", userHandler.GetProfile)
				users.PUT("/profile", userHandler.UpdateProfile)
			}

			// Property routes
			properties := protected.Group("/properties")
			{
				properties.GET("", propertyHandler.GetProperties)
				properties.GET("/:id", propertyHandler.GetProperty)
				properties.POST("", propertyHandler.CreateProperty)
				properties.PUT("/:id", propertyHandler.UpdateProperty)
				properties.DELETE("/:id", propertyHandler.DeleteProperty)
			}

			// Booking routes
			bookings := protected.Group("/bookings")
			{
				bookings.GET("", bookingHandler.GetBookings)
				bookings.GET("/:id", bookingHandler.GetBooking)
				bookings.POST("", bookingHandler.CreateBooking)
				bookings.PUT("/:id", bookingHandler.UpdateBooking)
				bookings.DELETE("/:id", bookingHandler.DeleteBooking)
			}
		}
	}

	// Create server
	srv := &http.Server{
		Addr:    ":" + cfg.Port,
		Handler: router,
	}

	// Graceful shutdown
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	log.Printf("Server started on port %s", cfg.Port)

	// Wait for interrupt signal to gracefully shutdown the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	// The context is used to inform the server it has 5 seconds to finish
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exiting")
}
