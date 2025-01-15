package service

import (
	"regexp"
	"sync"
	"time"

	"mock-ses-api/internal/models"

	"github.com/google/uuid"
)

type SESService struct {
	stats    models.Statistics
	quota    models.QuotaInfo
	warmup   models.WarmupRules
	mu       sync.RWMutex
	lastSend time.Time
}

func NewSESService() *SESService {
	return &SESService{
		stats: models.Statistics{
			DailyQuota: 50000, // AWS SES starting quota
		},
		quota: models.QuotaInfo{
			Max24HourSend:  50000,
			MaxSendRate:    14, // Emails per second
			SendingEnabled: true,
		},
		warmup: models.WarmupRules{
			DailyLimit:     50, // Start with 50 emails per day
			CurrentDay:     1,
			IncreaseFactor: 1.5, // Increase limit by 50% each day
			IsWarmedUp:     false,
		},
		lastSend: time.Now(),
	}
}

func (s *SESService) SendEmail(input *models.SendEmailInput) (*models.SendEmailOutput, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	// Check email warming status
	if !s.warmup.IsWarmedUp {
		if s.stats.TotalEmails >= s.warmup.DailyLimit {
			return nil, &models.APIError{
				Code:    "DailyQuotaExceeded",
				Message: "Account is in warm-up period. Daily limit exceeded.",
			}
		}
	}

	// Validate email
	if err := s.validateEmail(input.To); err != nil {
		s.stats.FailedSends++
		return nil, err
	}

	// Check rate limiting
	if time.Since(s.lastSend) < time.Second/time.Duration(s.quota.MaxSendRate) {
		return nil, &models.APIError{
			Code:    "ThrottlingException",
			Message: "Rate limit exceeded",
		}
	}

	// Update statistics
	s.stats.TotalEmails++
	s.stats.SuccessfulSends++
	s.stats.UsedQuota++
	s.lastSend = time.Now()

	return &models.SendEmailOutput{
		MessageId: uuid.New().String(),
	}, nil
}

func (s *SESService) GetStatistics() models.Statistics {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.stats
}

func (s *SESService) GetQuota() models.QuotaInfo {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.quota
}

func (s *SESService) GetWarmupStatus() models.WarmupRules {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.warmup
}

func (s *SESService) GetDetailedStatistics() map[string]interface{} {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return map[string]interface{}{
		"basic": s.stats,
		"warmup": map[string]interface{}{
			"currentDay":   s.warmup.CurrentDay,
			"dailyLimit":   s.warmup.DailyLimit,
			"isWarmedUp":   s.warmup.IsWarmedUp,
			"nextIncrease": time.Now().Add(24 * time.Hour).Format(time.RFC3339),
		},
		"performance": map[string]interface{}{
			"averageLatency": "0.15s", // Mock value
			"bounceRate":     float64(s.stats.FailedSends) / float64(s.stats.TotalEmails) * 100,
			"successRate":    float64(s.stats.SuccessfulSends) / float64(s.stats.TotalEmails) * 100,
		},
	}
}

func (s *SESService) ListIdentities() *models.ListIdentitiesOutput {
	return &models.ListIdentitiesOutput{
		Identities: []string{"domain.com", "test@domain.com"},
		NextToken:  "",
	}
}

func (s *SESService) validateEmail(email string) error {
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
	if !emailRegex.MatchString(email) {
		return &models.APIError{
			Code:    "InvalidParameterValue",
			Message: "Invalid email format",
		}
	}
	return nil
}
