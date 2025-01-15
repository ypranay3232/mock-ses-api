package metrics

import (
	"sync"
	"time"
)

type EmailMetrics struct {
	mu sync.RWMutex
	// Daily email counts
	dailyCount map[string]int
	// Total emails sent
	totalCount int
	// Throttled requests
	throttledCount int
	// Failed requests
	failedCount int
}

var metrics *EmailMetrics
var once sync.Once

func GetMetrics() *EmailMetrics {
	once.Do(func() {
		metrics = &EmailMetrics{
			dailyCount: make(map[string]int),
		}
	})
	return metrics
}

func (m *EmailMetrics) IncrementDaily() {
	m.mu.Lock()
	defer m.mu.Unlock()
	
	today := time.Now().Format("2006-01-02")
	m.dailyCount[today]++
	m.totalCount++
}

func (m *EmailMetrics) GetDailyCount() int {
	m.mu.RLock()
	defer m.mu.RUnlock()
	
	today := time.Now().Format("2006-01-02")
	return m.dailyCount[today]
}

func (m *EmailMetrics) IncrementThrottled() {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.throttledCount++
}

func (m *EmailMetrics) IncrementFailed() {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.failedCount++
} 