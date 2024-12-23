package service

import (
	"context"
	"errors"
	"fibonacci-api/internal/domain"
	"fibonacci-api/pkg/logger"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Мокаем кэш
type MockCache struct {
	mock.Mock
}

func (m *MockCache) Set(key string, value string) error {
	args := m.Called(key, value)
	return args.Error(0)
}

func (m *MockCache) Get(key string) (string, error) {
	args := m.Called(key)
	return args.String(0), args.Error(1)
}

func (m *MockCache) GetWithFallback(key string) (string, error) {
	args := m.Called(key)
	return args.String(0), args.Error(1)
}

func (m *MockCache) Delete(key string) error {
	args := m.Called(key)
	return args.Error(0)
}

func (m *MockCache) Exists(key string) (bool, error) {
	args := m.Called(key)
	return args.Bool(0), args.Error(1)
}

func (m *MockCache) Clear() error {
	args := m.Called()
	return args.Error(0)
}

func TestFibonacciService_GetOrCalculateFibonacci_CacheHit(t *testing.T) {
	logger := logger.NewLogger("stdout", "debug")
	mockCache := new(MockCache)
	service := NewFibonacciService(mockCache, logger)

	// Подготовка мока
	cacheKey := "fibonacci:10"
	expectedResult := &domain.FibonacciResult{Value: 55}
	mockCache.On("GetWithFallback", cacheKey).Return("55", nil)

	ctx := context.Background()
	// Вызов метода
	result, err := service.GetOrCalculateFibonacci(ctx, 10)

	// Проверка
	assert.NoError(t, err)
	assert.Equal(t, expectedResult, result)
	mockCache.AssertExpectations(t)
}

func TestFibonacciService_GetOrCalculateFibonacci_CacheMiss(t *testing.T) {
	logger := logger.NewLogger("stdout", "debug")
	mockCache := new(MockCache)
	service := NewFibonacciService(mockCache, logger)

	// Подготовка мока
	cacheKey := "fibonacci:10"
	expectedResult := &domain.FibonacciResult{Value: 55}
	mockCache.On("GetWithFallback", cacheKey).Return("", errors.New("cache miss"))
	mockCache.On("Set", cacheKey, "55").Return(nil)

	ctx := context.Background()

	// Вызов метода
	result, err := service.GetOrCalculateFibonacci(ctx, 10)

	// Проверка
	assert.NoError(t, err)
	assert.Equal(t, expectedResult, result)
	mockCache.AssertExpectations(t)
}

func TestFibonacciService_CalculateFibonacci(t *testing.T) {
	logger := logger.NewLogger("stdout", "debug")
	mockCache := new(MockCache)
	service := NewFibonacciService(mockCache, logger)

	// Проверка для маленького числа
	result := service.calculateFibonacci(10)
	assert.Equal(t, int64(55), result.Value)

	// Проверка для большого числа
	resultBig := service.calculateFibonacci(100)
	assert.NotNil(t, resultBig.BigValue)
}

func TestFibonacciService_GetFibonacciSequenceConcurrent_Success(t *testing.T) {
	logger := logger.NewLogger("stdout", "debug")
	mockCache := new(MockCache)
	service := NewFibonacciService(mockCache, logger)

	// Подготовка мока
	maxIndex := int64(20)
	mockCache.On("GetWithFallback", mock.Anything).Return("", errors.New("cache miss"))
	mockCache.On("Set", mock.Anything, mock.Anything).Return(nil)

	ctx := context.Background()

	// Вызов метода
	results, err := service.GetFibonacciSequenceConcurrent(ctx, maxIndex)

	// Проверка результатов
	assert.NoError(t, err)
	assert.Len(t, results, int(maxIndex)+1)

	// Проверка последовательности чисел Фибоначчи
	expected := []int64{0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, 377, 610, 987, 1597, 2584, 4181, 6765}
	for i, result := range results {
		assert.Equal(t, expected[i], result.Value)
	}
	mockCache.AssertExpectations(t)
}

func TestFibonacciService_GetFibonacciSequenceConcurrent_HighLoad(t *testing.T) {
	logger := logger.NewLogger("stdout", "debug")
	mockCache := new(MockCache)
	service := NewFibonacciService(mockCache, logger)

	// Подготовка мока
	maxIndex := int64(50)
	mockCache.On("GetWithFallback", mock.Anything).Return("", errors.New("cache miss"))
	mockCache.On("Set", mock.Anything, mock.Anything).Return(nil)

	ctx := context.Background()

	var wg sync.WaitGroup
	numWorkers := 100 // Количество конкурентных запросов

	wg.Add(numWorkers)
	for i := 0; i < numWorkers; i++ {
		go func() {
			defer wg.Done()
			_, err := service.GetFibonacciSequenceConcurrent(ctx, maxIndex)
			assert.NoError(t, err)
		}()
	}

	wg.Wait()
	mockCache.AssertExpectations(t)
}
