package main

import (
    "github.com/gin-gonic/gin"
    "mock-ses-api/internal/api/handlers"
    "mock-ses-api/internal/api/routes"
    "mock-ses-api/internal/service"
)

func main() {
    router := gin.Default()

    // Initialize services and handlers
    sesService := service.NewSESService()
    sesHandler := handlers.NewSESHandler(sesService)

    // Setup routes
    routes.SetupRoutes(router, sesHandler)

    // Start the server
    router.Run(":8080")
}