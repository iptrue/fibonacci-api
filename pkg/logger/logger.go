package logger

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// Logger - глобальная структура для логирования
type Logger struct {
	*log.Logger
	logLevel string
}

// NewLogger - создание нового экземпляра логгера с конфигурацией
func NewLogger(outputType string, logLevel string) *Logger {
	var logOutput *os.File

	if strings.ToLower(outputType) == "file" {
		// Генерация имени файла с текущей датой и временем
		timestamp := time.Now().Format("2006-01-02_15-04-05")
		logFileName := fmt.Sprintf("fibonacci-api-%s.log", timestamp)

		// Создание или открытие файла для записи логов
		logFile, err := os.OpenFile(filepath.Join(".", logFileName), os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Printf("Could not open log file: %v, using default stderr\n", err)
			logOutput = os.Stderr
		} else {
			logOutput = logFile
		}
	} else {
		// Логирование в консоль по умолчанию
		logOutput = os.Stdout
	}

	// Если уровень логирования не задан, используем уровень по умолчанию
	if logLevel == "" {
		logLevel = "info"
	}

	// Создание нового логгера с пользовательскими флагами
	logger := log.New(logOutput, "", log.Lmsgprefix)

	return &Logger{
		Logger:   logger,
		logLevel: strings.ToLower(logLevel),
	}
}

// logMessage - универсальный метод для логирования сообщений
func (l *Logger) logMessage(level, msg string, args ...any) {
	levelPriority := map[string]int{"debug": 4, "info": 3, "warn": 2, "error": 1, "fatal": 0}
	currentPriority := levelPriority[strings.ToLower(l.logLevel)]

	if levelPriority[level] > currentPriority {
		return
	}

	// Форматирование времени с миллисекундами
	timestamp := time.Now().Format("2006-01-02 15:04:05.000")

	// Преобразование структур в форматированный JSON
	serializedArgs := make([]string, len(args))
	for i, arg := range args {
		if jsonBytes, err := json.MarshalIndent(arg, "", "  "); err == nil {
			serializedArgs[i] = string(jsonBytes)
		} else {
			serializedArgs[i] = fmt.Sprintf("serialization error: %v", err)
		}
	}

	// Формирование итогового сообщения
	formattedMessage := fmt.Sprintf("%s [%s] %s %s", timestamp, strings.ToUpper(level), msg, strings.Join(serializedArgs, " "))

	if level == "fatal" {
		l.Logger.Fatalf(formattedMessage)
	} else {
		l.Logger.Println(formattedMessage)
	}
}

// Info - логирование информационных сообщений
func (l *Logger) Info(msg string, args ...any) {
	l.logMessage("info", msg, args...)
}

// Debug - логирование отладочных сообщений
func (l *Logger) Debug(msg string, args ...any) {
	l.logMessage("debug", msg, args...)
}

// Warn - логирование предупреждений
func (l *Logger) Warn(msg string, args ...any) {
	l.logMessage("warn", msg, args...)
}

// Error - логирование ошибок
func (l *Logger) Error(msg string, args ...any) {
	l.logMessage("error", msg, args...)
}

// Fatal - логирование фатальных ошибок
func (l *Logger) Fatal(msg string, args ...any) {
	l.logMessage("fatal", msg, args...)
}
