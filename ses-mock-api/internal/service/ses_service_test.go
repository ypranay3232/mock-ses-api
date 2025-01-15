package service

import (
	"testing"
	"mock-ses-api/internal/models"
)

func TestSendEmail(t *testing.T) {
	sesService := NewSESService()

	tests := []struct {
		name     string
		input    *models.SendEmailInput
		wantErr  bool
	}{
		{
			name: "Valid email",
			input: &models.SendEmailInput{
				To:      "test@example.com",
				Subject: "Test Subject",
				Body:    "Test Body",
			},
			wantErr: false,
		},
		{
			name: "Invalid email",
			input: &models.SendEmailInput{
				To:      "invalid-email",
				Subject: "Test Subject",
				Body:    "Test Body",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := sesService.SendEmail(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("SESService.SendEmail() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err == nil && result.MessageId == "" {
				t.Error("Expected MessageId but got empty string")
			}
		})
	}
}