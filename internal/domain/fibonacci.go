package domain

import (
	"fmt"
	"math/big"
	"strconv"
)

type FibonacciResult struct {
	Value    int64    // Используется для чисел <= 92
	BigValue *big.Int // Используется для чисел > 92
}

// Marshal сериализует FibonacciResult в строку.
func (r *FibonacciResult) Marshal() (string, error) {
	if r.BigValue != nil {
		// Для больших значений возвращаем строку из big.Int
		return r.BigValue.String(), nil
	}
	// Для меньших значений возвращаем строковое представление числа
	return strconv.FormatInt(r.Value, 10), nil
}

// Unmarshal десериализует строку в FibonacciResult.
func (r *FibonacciResult) Unmarshal(value string) error {
	if value == "" {
		return fmt.Errorf("cached value is empty")
	}

	// Сначала пробуем распарсить как int64
	parsedInt, err := strconv.ParseInt(value, 10, 64)
	if err == nil {
		r.Value = parsedInt
		r.BigValue = nil // Сбрасываем BigValue для согласованности
		return nil
	}

	// Если значение превышает диапазон int64, пробуем использовать big.Int
	bigValue := new(big.Int)
	if _, ok := bigValue.SetString(value, 10); ok {
		r.BigValue = bigValue
		r.Value = 0 // Сбрасываем Value для согласованности
		return nil
	}

	return fmt.Errorf("failed to parse value: %v", err)
}

// FibonacciInt64 вычисляет число Фибоначчи для числа <= 92 (тип int64).
func FibonacciInt64(n int64) int64 {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}

	a, b := int64(0), int64(1)
	for i := int64(2); i <= n; i++ {
		a, b = b, a+b
	}
	return b
}

// FibonacciBig вычисляет число Фибоначчи для чисел > 92 (тип big.Int).
func FibonacciBig(n int64) *big.Int {
	a := big.NewInt(0)
	b := big.NewInt(1)

	for i := int64(2); i <= n; i++ {
		next := new(big.Int).Add(a, b)
		a, b = b, next
	}
	return b
}
