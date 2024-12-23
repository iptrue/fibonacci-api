package repository

import (
	"context"
	"fibonacci-api/configs"
	"fibonacci-api/pkg/logger"
	"fmt"
	"github.com/go-redis/redis/v8"
	"sync"
	"time"
)

// RedisCache - структура для работы с Redis
type RedisCache struct {
	client            *redis.Client
	expiry            time.Duration
	context           context.Context
	logger            *logger.Logger
	localCache        sync.Map
	cacheSize         int
	maxLocalCacheSize int
	cacheMutex        sync.Mutex
}

func NewRedisCache(client *redis.Client, config *configs.AppConfig, logger *logger.Logger) Cache {
	return &RedisCache{
		client:            client,
		context:           context.Background(),
		logger:            logger,
		maxLocalCacheSize: config.LocalCacheSize,
		expiry:            config.CacheExpiration,
	}
}

// Set - установка значения в Redis и локальный кэш
func (r *RedisCache) Set(key string, value string) error {
	// Сохраняем значение в Redis
	err := r.client.Set(r.context, key, value, r.expiry).Err()
	if err != nil {
		r.logError("Set", key, err)
		return fmt.Errorf("could not set key %s in redis: %v", key, err)
	}

	r.cacheMutex.Lock()
	defer r.cacheMutex.Unlock()

	if r.cacheSize >= r.maxLocalCacheSize {
		r.localCache.Range(func(key, value interface{}) bool {
			// Удаляем первое найденное значение
			r.localCache.Delete(key)
			r.cacheSize--
			return false
		})
	}

	// Сохраняем в локальный кэш
	r.localCache.Store(key, value)
	r.cacheSize++

	return nil
}

// Get - получение значения из Redis
func (r *RedisCache) Get(key string) (string, error) {
	val, err := r.client.Get(r.context, key).Result()
	if err != nil {
		if err == redis.Nil {
			return "", redis.Nil
		}
		r.logError("Get", key, err)
		return "", fmt.Errorf("could not get key %s from redis: %v", key, err)
	}
	return val, nil
}

// GetWithFallback - получение значения с fallback: сначала из map, потом из Redis
func (r *RedisCache) GetWithFallback(key string) (string, error) {
	// Проверяем локальный кэш
	if val, ok := r.localCache.Load(key); ok {
		return val.(string), nil
	}

	// Если в map нет, идем в Redis
	val, err := r.Get(key)
	if err == nil {
		// Если данные есть в Redis, сохраняем их в локальный кэш для быстрого доступа
		r.localCache.Store(key, val)
		r.cacheSize++ // Увеличиваем размер кэша при добавлении
	}
	return val, err
}

// Delete - удаление значения из Redis и локального кэша
func (r *RedisCache) Delete(key string) error {
	err := r.client.Del(r.context, key).Err()
	if err != nil {
		r.logError("Delete", key, err)
		return fmt.Errorf("could not delete key %s from redis: %v", key, err)
	}

	// Удаляем из локального кэша
	r.localCache.Delete(key)
	r.cacheSize--
	return nil
}

// Exists - проверка существования ключа в кэше
func (r *RedisCache) Exists(key string) (bool, error) {
	val, err := r.client.Exists(r.context, key).Result()
	if err != nil {
		r.logError("Exists", key, err)
		return false, fmt.Errorf("could not check existence of key %s in redis: %v", key, err)
	}
	return val > 0, nil
}

// Clear - очистка всего кэша
func (r *RedisCache) Clear() error {
	err := r.client.FlushAll(r.context).Err()
	if err != nil {
		r.logError("Clear", "", err)
		return fmt.Errorf("could not clear redis cache: %v", err)
	}

	// Очистка локального кэша
	r.localCache = sync.Map{}
	r.cacheSize = 0
	return nil
}

// Логирование ошибок Redis
func (r *RedisCache) logError(action, key string, err error) {
	r.logger.Error("Redis %s error for key %s: %v", action, key, err)
}
