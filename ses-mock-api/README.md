# Mock SES API

A mock implementation of AWS Simple Email Service (SES) API using Go and Gin framework.

## Table of Contents
- [Prerequisites](#prerequisites)
- [Installation](#installation)
- [Project Structure](#project-structure)
- [Running the Application](#running-the-application)
- [API Endpoints](#api-endpoints)
- [Testing with Postman](#testing-with-postman)
- [Running Tests](#running-tests)
- [Additional Features](#additional-features)

## Prerequisites

- Go 1.16 or higher
- Git
- Postman (for testing APIs)

## Installation

1. Clone the repository

bash
git clone https://github.com/yourusername/mock-ses-api.git

cd mock-ses-api


2. Initialize Go module

go mod init mock-ses-api



3. Install dependencies


go get github.com/gin-gonic/gin
go get github.com/google/uuid
go mod tidy



## Project Structure
mock-ses-api/
├── cmd/
│ └── main.go # Server entry point
├── internal/
│ ├── api/
│ │ ├── handlers/ # HTTP handlers
│ │ └── routes/ # Route definitions
│ ├── models/ # Data structures
│ └── service/ # Business logic
│ ├── ses_service.go
│ └── ses_service_test.go
└── README.md



## Running the Application

1. Start the server

go run cmd/main.go


2. The server will start on `http://localhost:8080`

## Testing with Postman
## API Endpoints

### 1. Send Email
- **Method**: POST

- **Endpoint**: `/v1/email/send`

- **Body**:
json
{
"to": "test@example.com",
"subject": "Test Email",
"body": "Hello, this is a test email!"
}



### 2. Test Email Warming Rules
Using Postman Runner:
1. Click "Runner" in Postman
2. Add the send email request
3. Set:
   - Iterations: 51 (to test 50/day limit)
   - Delay: 100ms
4. Expected Results:
   - First 50: Success
   - 51st: Quota exceeded error

### 3. Test Invalid Email

json
{
"to": "invalid-email",
"subject": "Test Email",
"body": "Should fail!"
}


### 4. Check Statistics
- Method: GET
- URL: `http://localhost:8080/v1/email/statistics`

### 5. Check Quota
- Method: GET
- URL: `http://localhost:8080/v1/email/quota`

## Expected Responses

### Successful Send
json
{
"messageId": "123e4567-e89b-12d3-a456-426614174000"
}


### Quota Exceeded
json
{
"code": "DailyQuotaExceeded",
"message": "Account is in warm-up period. Daily limit exceeded."
}


### Invalid Email
json
{
"code": "InvalidParameterValue",
"message": "Invalid email format"
}


## Additional Features

### Email Warming Up
- Starts with 50 emails per day
- Increases limit by 50% each day
- Full quota after 14 days
- Automatic daily limit tracking

### Statistics Tracking
- Total emails sent
- Success/failure rates
- Current warming up status
- Performance metrics
- Quota usage

### Rate Limiting
- Maximum 14 emails per second
- Daily quota enforcement
- Warming up period restrictions

## AWS SES Documentation References
- [AWS SES API Reference](https://docs.aws.amazon.com/ses/latest/APIReference/Welcome.html)
- [AWS SES Developer Guide](https://docs.aws.amazon.com/ses/latest/DeveloperGuide/Welcome.html)
- [Email Sending Limits](https://docs.aws.amazon.com/ses/latest/DeveloperGuide/manage-sending-limits.html)
- [Account Warm-up](https://docs.aws.amazon.com/ses/latest/DeveloperGuide/warm-up.html)



## Error Handling

The API returns appropriate error messages for invalid requests:
json
{
"code": "ThrottlingException",
"error": "invalid email format"
}



## Security Notice
This is a mock implementation intended for testing and development purposes only. Do not use this for sending actual emails in production environments.


## Troubleshooting

Common issues and solutions:

1. **Port 8080 already in use**
   - Error: `listen tcp :8080: bind: address already in use`
   - Solution: Kill the process using port 8080 or change the port in main.go

2. **Missing dependencies**
   - Error: `cannot find package "github.com/gin-gonic/gin"`
   - Solution: Run `go mod tidy` to install all dependencies

3. **Invalid email format**
   - Error: `{"code": "ThrottlingException", "error": "invalid email format"}`
   - Solution: Ensure email address follows standard format (e.g., "user@example.com")

