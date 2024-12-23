package ratelimiter

import (
	"errors"
	"sync"
	"time"
)

type RateLimiter struct {
	mu          sync.Mutex
	tokens      int
	maxTokens   int
	refillRate  time.Duration
	lastRefill  time.Time
	refillMutex sync.Mutex
}

func NewRateLimiter(maxTokens int, refillRate time.Duration) *RateLimiter {
	return &RateLimiter{
		tokens:     maxTokens,
		maxTokens:  maxTokens,
		refillRate: refillRate,
		lastRefill: time.Now(),
	}
}

// acquire пытается забрать токен. Если токенов нет, возвращает ошибку.
func (r *RateLimiter) Acquire() error {
	r.mu.Lock()
	defer r.mu.Unlock()

	// Обновляем токены в случае, если они не были обновлены.
	r.refillTokens()

	// Если токены есть, забираем один
	if r.tokens > 0 {
		r.tokens--
		return nil
	}

	// Если токенов нет, отклоняем запрос
	return errors.New("rate limit exceeded")
}

// refillTokens пополняет количество токенов на основе времени.
func (r *RateLimiter) refillTokens() {
	r.refillMutex.Lock()
	defer r.refillMutex.Unlock()

	now := time.Now()
	elapsed := now.Sub(r.lastRefill)

	// Вычисляем, сколько токенов можем добавить
	refillCount := int(elapsed / r.refillRate)

	if refillCount > 0 {
		// Пополняем токены, но не больше максимально допустимого
		r.tokens = min(r.maxTokens, r.tokens+refillCount)
		r.lastRefill = now
	}
}

// min возвращает минимальное значение из двух.
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
