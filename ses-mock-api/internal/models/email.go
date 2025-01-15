package models

// SendEmailInput represents the input for sending an email
type SendEmailInput struct {
	To      string `json:"to" binding:"required,email"`
	Subject string `json:"subject" binding:"required"`
	Body    string `json:"body" binding:"required"`
}

// SendEmailOutput represents the response after sending an email
type SendEmailOutput struct {
	MessageId string `json:"messageId"`
}

// Statistics represents email sending statistics
type Statistics struct {
	TotalEmails     int     `json:"totalEmails"`
	SuccessfulSends int     `json:"successfulSends"`
	FailedSends     int     `json:"failedSends"`
	DailyQuota      int     `json:"dailyQuota"`
	UsedQuota       int     `json:"usedQuota"`
	SendRate        float64 `json:"sendRate"`
}

// QuotaInfo represents quota information
type QuotaInfo struct {
	Max24HourSend   int     `json:"max24HourSend"`
	MaxSendRate     float64 `json:"maxSendRate"`
	SentLast24Hours int     `json:"sentLast24Hours"`
	SendingEnabled  bool    `json:"sendingEnabled"`
}

// WarmupRules represents email warming up configuration
type WarmupRules struct {
	DailyLimit     int     `json:"dailyLimit"`
	CurrentDay     int     `json:"currentDay"`
	IncreaseFactor float64 `json:"increaseFactor"`
	IsWarmedUp     bool    `json:"isWarmedUp"`
}

// ListIdentitiesOutput represents the response for listing identities
type ListIdentitiesOutput struct {
	Identities []string `json:"identities"`
	NextToken  string   `json:"nextToken,omitempty"`
}

// APIError represents an API error response
type APIError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

// Error implements the error interface
func (e *APIError) Error() string {
	return e.Message
}
