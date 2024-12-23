package main

import (
	"context"
	"fibonacci-api/configs"
	"fibonacci-api/internal/handler"
	"fibonacci-api/internal/metrics"
	"fibonacci-api/internal/repository"
	"fibonacci-api/internal/service"
	"fibonacci-api/pkg/logger"
	ratelimiter "fibonacci-api/pkg/rate_limiter"
	pb "fibonacci-api/proto/generated"
	"fmt"
	"github.com/go-redis/redis/v8"
	"google.golang.org/grpc"
	"net"
	"os"
	"os/signal"
	"time"
)

func main() {

	// Загрузка конфигурации
	config, err := configs.LoadConfig("configs/config.yaml")
	if err != nil {
		panic(fmt.Sprintf("Failed to load config: %v", err))
	}

	log := logger.NewLogger(config.LogOutput, config.LogLevel)

	// Инициализация метрик
	appMetrics := metrics.NewMetrics("fibonacci_api")
	appMetrics.Register()
	go metrics.ServeMetrics(config.PrometheusAddr)

	log.Info("Configuration loaded successfully", "config", config)

	// Инициализация Redis
	redisClient := redis.NewClient(&redis.Options{
		Addr: config.RedisAddr,
	})
	defer redisClient.Close()

	if err := redisClient.Ping(context.Background()).Err(); err != nil {
		log.Fatal("Failed to connect to Redis", "error", err)
	}

	redisCache := repository.NewRedisCache(redisClient, config, log)
	svc := service.NewFibonacciService(redisCache, log)

	// Настройка gRPC сервера
	grpcServer := grpc.NewServer()
	rateLimiter := ratelimiter.NewRateLimiter(config.MaxTokens, time.Duration(config.RefillRate))
	server := handler.NewFibonacciHandler(svc, log, config, rateLimiter, appMetrics)
	pb.RegisterFibonacciServiceServer(grpcServer, server)

	listener, err := net.Listen("tcp", config.GRPCPort)
	if err != nil {
		log.Fatal("Failed to listen on gRPC port", "port", config.GRPCPort, "error", err)
	}

	go func() {
		log.Info("Starting gRPC server", "port", config.GRPCPort)
		if err := grpcServer.Serve(listener); err != nil {
			log.Fatal("Failed to start gRPC server", "error", err)
		}
	}()

	// Обработка завершения приложения
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	grpcServer.GracefulStop()
	log.Info("gRPC server stopped gracefully")
}
