package repository

import (
	"context"
	"fibonacci-api/pkg/logger"
	"testing"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/stretchr/testify/assert"
)

const redisAddr = "localhost:6379"

func newRedisClient(t *testing.T) *redis.Client {
	client := redis.NewClient(&redis.Options{Addr: redisAddr})
	if err := client.Ping(context.Background()).Err(); err != nil {
		t.Fatalf("Failed to connect to Redis: %v", err)
	}
	return client
}

func newTestLogger() *logger.Logger {
	return logger.NewLogger("console", "debug")
}

func TestRedisCache_SetAndGet(t *testing.T) {
	redisClient := newRedisClient(t)
	defer redisClient.Close()

	log := newTestLogger()
	cache := NewRedisCache(redisClient, time.Minute, log)

	key := "testKey"
	value := "testValue"

	err := cache.Set(key, value)
	assert.NoError(t, err)

	gotValue, err := cache.Get(key)
	assert.NoError(t, err)
	assert.Equal(t, value, gotValue)
}

func TestRedisCache_GetWithFallback(t *testing.T) {
	redisClient := newRedisClient(t)
	defer redisClient.Close()

	log := newTestLogger()
	cache := NewRedisCache(redisClient, time.Minute, log)

	key := "testFallbackKey"
	value := "fallbackValue"

	err := cache.Set(key, value)
	assert.NoError(t, err)

	gotValue, err := cache.GetWithFallback(key)
	assert.NoError(t, err)
	assert.Equal(t, value, gotValue)
}

func TestRedisCache_KeyDoesNotExist(t *testing.T) {
	redisClient := newRedisClient(t)
	defer redisClient.Close()

	log := newTestLogger()
	cache := NewRedisCache(redisClient, time.Minute, log)

	key := "nonExistentKey"
	value, err := cache.Get(key)
	assert.Error(t, err)
	assert.Empty(t, value)
}

func TestRedisCache_Delete(t *testing.T) {
	redisClient := newRedisClient(t)
	defer redisClient.Close()

	log := newTestLogger()
	cache := NewRedisCache(redisClient, time.Minute, log)

	key := "testKeyToDelete"
	value := "valueToDelete"
	err := cache.Set(key, value)
	assert.NoError(t, err)

	err = cache.Delete(key)
	assert.NoError(t, err)

	gotValue, err := cache.Get(key)
	assert.Error(t, err)
	assert.Empty(t, gotValue)
}

func TestRedisCache_Exists(t *testing.T) {
	redisClient := newRedisClient(t)
	defer redisClient.Close()

	log := newTestLogger()
	cache := NewRedisCache(redisClient, time.Minute, log)

	key := "testKeyExist"
	value := "valueExist"
	err := cache.Set(key, value)
	assert.NoError(t, err)

	exists, err := cache.Exists(key)
	assert.NoError(t, err)
	assert.True(t, exists)

	nonExistentKey := "nonExistentKey"
	exists, err = cache.Exists(nonExistentKey)
	assert.NoError(t, err)
	assert.False(t, exists)
}

func TestRedisCache_Clear(t *testing.T) {
	redisClient := newRedisClient(t)
	defer redisClient.Close()

	log := newTestLogger()
	cache := NewRedisCache(redisClient, time.Minute, log)

	cache.Set("key1", "value1")
	cache.Set("key2", "value2")

	err := cache.Clear()
	assert.NoError(t, err)

	exists1, _ := cache.Exists("key1")
	exists2, _ := cache.Exists("key2")
	assert.False(t, exists1)
	assert.False(t, exists2)
}
