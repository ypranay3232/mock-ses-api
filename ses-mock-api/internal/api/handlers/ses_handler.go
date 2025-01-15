package handlers

import (
	"mock-ses-api/internal/models"
	"mock-ses-api/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SESHandler struct {
	sesService *service.SESService
}

func NewSESHandler(sesService *service.SESService) *SESHandler {
	return &SESHandler{
		sesService: sesService,
	}
}

func (h *SESHandler) SendEmail(c *gin.Context) {
	var input models.SendEmailInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, &models.APIError{
			Code:    "InvalidParameterValue",
			Message: "Invalid request body",
		})
		return
	}

	result, err := h.sesService.SendEmail(&input)
	if err != nil {
		if apiErr, ok := err.(*models.APIError); ok {
			c.JSON(http.StatusBadRequest, apiErr)
			return
		}
		c.JSON(http.StatusInternalServerError, &models.APIError{
			Code:    "InternalError",
			Message: "An internal error occurred",
		})
		return
	}

	c.JSON(http.StatusOK, result)
}

func (h *SESHandler) GetSendStatistics(c *gin.Context) {
	stats := h.sesService.GetStatistics()
	c.JSON(http.StatusOK, stats)
}

func (h *SESHandler) GetDetailedStatistics(c *gin.Context) {
	stats := h.sesService.GetDetailedStatistics()
	c.JSON(http.StatusOK, stats)
}

func (h *SESHandler) GetSendQuota(c *gin.Context) {
	quota := h.sesService.GetQuota()
	c.JSON(http.StatusOK, quota)
}

func (h *SESHandler) GetWarmupStatus(c *gin.Context) {
	status := h.sesService.GetWarmupStatus()
	c.JSON(http.StatusOK, status)
}

func (h *SESHandler) ListIdentities(c *gin.Context) {
	identities := h.sesService.ListIdentities()
	c.JSON(http.StatusOK, identities)
}
