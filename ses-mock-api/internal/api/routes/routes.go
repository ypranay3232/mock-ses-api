package routes

import (
    "github.com/gin-gonic/gin"
    "mock-ses-api/internal/api/handlers"
)

func SetupRoutes(router *gin.Engine, sesHandler *handlers.SESHandler) {
    v1 := router.Group("/v1")
    {
        // Email endpoints
        v1.POST("/email/send", sesHandler.SendEmail)
        v1.GET("/email/statistics", sesHandler.GetSendStatistics)
        v1.GET("/email/detailed-statistics", sesHandler.GetDetailedStatistics)
        v1.GET("/email/quota", sesHandler.GetSendQuota)
        v1.GET("/email/warmup-status", sesHandler.GetWarmupStatus)

        // Identity endpoints
        v1.GET("/identities", sesHandler.ListIdentities)
    }
}