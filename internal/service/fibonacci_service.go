package service

import (
	"context"
	"fibonacci-api/internal/domain"
	"fibonacci-api/internal/repository"
	"fibonacci-api/pkg/logger"
	"fmt"
	"sync"
)

type FibonacciService struct {
	cache  repository.Cache
	logger *logger.Logger
}

func NewFibonacciService(cache repository.Cache, logger *logger.Logger) *FibonacciService {
	return &FibonacciService{
		cache:  cache,
		logger: logger,
	}
}

func (s *FibonacciService) GetFibonacciSequenceConcurrent(ctx context.Context, maxIndex int64) ([]*domain.FibonacciResult, error) {
	s.logger.Debug("Starting concurrent Fibonacci sequence calculation", maxIndex)

	results := make([]*domain.FibonacciResult, maxIndex+1)
	var wg sync.WaitGroup
	var mu sync.Mutex
	errCh := make(chan error, 1)

	for i := int64(0); i <= maxIndex; i++ {
		select {
		case <-ctx.Done():
			s.logger.Warn("Fibonacci calculation cancelled before starting for max index", maxIndex)
			return nil, ctx.Err()
		default:
		}

		wg.Add(1)
		go func(index int64) {
			defer wg.Done()
			select {
			case <-ctx.Done(): // Если контекст отменён, прекращаем выполнение
				s.logger.Debug("Goroutine cancelled for index", index)
				return
			default:
				s.logger.Debug("Processing Fibonacci calculation for index", index)
				result, err := s.GetOrCalculateFibonacci(ctx, index)
				if err != nil {
					s.sendErrorOnce(errCh, err)
					s.logger.Error("Error during Fibonacci calculation for index", index, err)
					return
				}

				mu.Lock()
				results[index] = result
				mu.Unlock()
				s.logger.Debug("Successfully calculated Fibonacci for index", index)
			}
		}(i)
	}

	// Ожидание завершения горутин с проверкой отмены
	doneCh := make(chan struct{})
	go func() {
		wg.Wait()
		close(doneCh)
	}()

	select {
	case <-ctx.Done(): // Отмена контекста
		s.logger.Warn("Fibonacci calculation cancelled during execution for max index", maxIndex)
		return nil, ctx.Err()
	case err := <-errCh: // Ошибка из одной из горутин
		s.logger.Error("Error occurred during concurrent Fibonacci calculation", err)
		return nil, err
	case <-doneCh: // Завершение всех горутин
		s.logger.Info("Completed Fibonacci sequence calculation", maxIndex)
		return results, nil
	}
}

// Метод для записи в канал ошибок единожды
func (s *FibonacciService) sendErrorOnce(errCh chan error, err error) {
	select {
	case errCh <- err:
	default: // Избегаем блокировки, если канал уже заполнен
	}
}

func (s *FibonacciService) GetOrCalculateFibonacci(ctx context.Context, index int64) (*domain.FibonacciResult, error) {
	s.logger.Debug("Fetching or calculating Fibonacci number for index", index)

	cacheKey := fmt.Sprintf("fibonacci:%d", index)
	cacheValue, err := s.cache.GetWithFallback(cacheKey)
	if err == nil {
		var result domain.FibonacciResult
		if err := result.Unmarshal(cacheValue); err == nil {
			s.logger.Debug("Cache hit for Fibonacci index", index)
			return &result, nil
		}
		s.logger.Debug("Cache miss: unable to unmarshal cached value for index", index, err)
	}

	// Прерывание, если контекст отменен
	if ctx.Err() != nil {
		return nil, ctx.Err()
	}

	s.logger.Debug("Cache miss: calculating Fibonacci for index", index)
	result := s.calculateFibonacci(index)
	if err := s.cacheResult(index, result); err != nil {
		s.logger.Debug("Failed to cache Fibonacci result for index", index, err)
	}
	return result, nil
}

func (s *FibonacciService) calculateFibonacci(index int64) *domain.FibonacciResult {
	s.logger.Debug("Calculating Fibonacci value for index", index)
	var result domain.FibonacciResult
	if index <= 92 {
		result.Value = domain.FibonacciInt64(index)
		s.logger.Debug("Calculated small Fibonacci value for index", index, result.Value)
	} else {
		result.BigValue = domain.FibonacciBig(index)
		s.logger.Debug("Calculated large Fibonacci value for index", index, result.BigValue.String())
	}
	return &result
}

func (s *FibonacciService) cacheResult(index int64, result *domain.FibonacciResult) error {
	s.logger.Debug("Caching Fibonacci result for index", index)
	cacheKey := fmt.Sprintf("fibonacci:%d", index)
	cacheValue, err := result.Marshal()
	if err != nil {
		s.logger.Error("Failed to marshal Fibonacci result for caching", index, err)
		return fmt.Errorf("failed to marshal Fibonacci result: %v", err)
	}
	if err := s.cache.Set(cacheKey, cacheValue); err != nil {
		s.logger.Warn("Failed to set Fibonacci result in cache", index, err)
		return err
	}
	s.logger.Debug("Successfully cached Fibonacci result for index", index)
	return nil
}
