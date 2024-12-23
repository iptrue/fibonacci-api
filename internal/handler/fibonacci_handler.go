package handler

import (
	"context"
	"fibonacci-api/configs"
	"fibonacci-api/internal/domain"
	"fibonacci-api/internal/metrics"
	"fibonacci-api/internal/service"
	"fibonacci-api/pkg/logger"
	ratelimiter "fibonacci-api/pkg/rate_limiter"
	proto "fibonacci-api/proto/generated"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

type FibonacciHandler struct {
	svc         *service.FibonacciService
	log         *logger.Logger
	rateLimiter *ratelimiter.RateLimiter
	metrics     *metrics.Metrics
	config      *configs.AppConfig
	proto.UnimplementedFibonacciServiceServer
}

func NewFibonacciHandler(svc *service.FibonacciService, logger *logger.Logger, config *configs.AppConfig, rateLimiter *ratelimiter.RateLimiter, metrics *metrics.Metrics) *FibonacciHandler {
	return &FibonacciHandler{
		svc:         svc,
		log:         logger,
		config:      config,
		rateLimiter: rateLimiter,
		metrics:     metrics,
	}
}

func (h *FibonacciHandler) GetFibonacciNumber(ctx context.Context, req *proto.FibonacciRequest) (*proto.FibonacciResponse, error) {
	startTime := time.Now()

	timeoutCtx, cancel := context.WithTimeout(ctx, h.config.GRPCTimeout)
	defer cancel()

	h.metrics.RequestCount.WithLabelValues("grpc", "GetFibonacciNumber").Inc()

	index := req.GetIndex()

	if err := h.rateLimiter.Acquire(); err != nil {
		h.log.Error("Rate limit exceeded for Fibonacci request: ", index)
		return nil, status.Errorf(codes.ResourceExhausted, "Rate limit exceeded")
	}

	if index <= 0 {
		h.log.Error("Invalid index, must be > 0: ", index)
		return nil, status.Errorf(codes.InvalidArgument, "Index must be positive")
	}

	if index > h.config.MaxFibonacciN {
		h.log.Error("Index too large, Max allowed index is ", h.config.MaxFibonacciN)
		return nil, status.Errorf(codes.InvalidArgument, "Index is too big, current max index is %d", h.config.MaxFibonacciN)
	}

	resultCh := make(chan *domain.FibonacciResult)
	errCh := make(chan error)

	go func() {
		result, err := h.svc.GetOrCalculateFibonacci(timeoutCtx, index)
		if err != nil {
			errCh <- err
		} else {
			resultCh <- result
		}
	}()

	select {
	case <-timeoutCtx.Done():
		h.log.Warn("Fibonacci calculation cancelled for index", index)
		return nil, status.Errorf(codes.DeadlineExceeded, "Request cancelled or timed out")
	case err := <-errCh:
		h.log.Error("Fibonacci service error for index", index, err)
		return nil, status.Errorf(codes.Internal, "Calculation error")
	case result := <-resultCh:
		response := &proto.FibonacciResponse{}
		if result.BigValue != nil {
			response.Value = &proto.FibonacciResponse_BigValue{BigValue: result.BigValue.String()}
		} else {
			response.Value = &proto.FibonacciResponse_IntValue{IntValue: result.Value}
		}

		duration := time.Since(startTime).Seconds()
		h.metrics.RequestDuration.WithLabelValues("grpc", "GetFibonacciNumber").Observe(duration)
		h.log.Debug("Successfully retrieved Fibonacci number for index", index)
		return response, nil
	}
}

func (h *FibonacciHandler) GetFibonacciSequence(ctx context.Context, req *proto.FibonacciSequenceRequest) (*proto.FibonacciSequenceResponse, error) {
	startTime := time.Now()

	timeoutCtx, cancel := context.WithTimeout(ctx, h.config.GRPCTimeout)
	defer cancel()

	h.metrics.RequestCount.WithLabelValues("grpc", "GetFibonacciSequence").Inc()

	maxIndex := req.GetMaxIndex()

	if maxIndex == 0 {
		h.log.Error("Invalid max_index, Max index cannot be zero: ", maxIndex)
		return nil, status.Errorf(codes.InvalidArgument, "Max index cannot be zero")
	}
	if maxIndex < 0 || maxIndex > h.config.MaxFibonacciN {
		return nil, status.Errorf(codes.InvalidArgument, "Max index must be in range [1, %d]", h.config.MaxFibonacciN)
	}

	resultCh := make(chan []*domain.FibonacciResult)
	errCh := make(chan error)

	go func() {
		result, err := h.svc.GetFibonacciSequenceConcurrent(timeoutCtx, maxIndex)
		if err != nil {
			h.log.Error("Failed to generate Fibonacci sequence", err)
			errCh <- err
		} else {
			resultCh <- result
		}
	}()

	select {
	case <-timeoutCtx.Done():
		h.log.Warn("Fibonacci Sequence calculation cancelled for max index", maxIndex)
		return nil, status.Errorf(codes.DeadlineExceeded, "Request cancelled or timed out")
	case err := <-errCh:
		h.log.Error("Fibonacci service error for max index", maxIndex, err)
		return nil, status.Errorf(codes.Internal, "Calculation error")
	case sequence := <-resultCh:
		protoSequence := make([]string, len(sequence))
		for i, res := range sequence {
			if res.BigValue != nil {
				protoSequence[i] = res.BigValue.String()
			} else {
				protoSequence[i] = fmt.Sprintf("%d", res.Value)
			}
		}

		duration := time.Since(startTime).Seconds()
		h.metrics.RequestDuration.WithLabelValues("grpc", "GetFibonacciSequence").Observe(duration)
		h.log.Debug("Successfully retrieved Fibonacci Sequence for max index", maxIndex)
		return &proto.FibonacciSequenceResponse{Sequence: protoSequence}, nil
	}
}
