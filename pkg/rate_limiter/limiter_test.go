package ratelimiter_test

import (
	"fibonacci-api/pkg/rate_limiter"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestRateLimiter(t *testing.T) {
	limiter := ratelimiter.NewRateLimiter(2, time.Second)

	// Пытаемся получить токен
	err := limiter.Acquire()
	assert.NoError(t, err)

	err = limiter.Acquire()
	assert.NoError(t, err)

	// Должен вернуть ошибку, так как лимит исчерпан
	err = limiter.Acquire()
	assert.Error(t, err)

	// Ждем, чтобы пополнились токены
	time.Sleep(1 * time.Second)

	// Токен должен быть снова доступен
	err = limiter.Acquire()
	assert.NoError(t, err)
}
