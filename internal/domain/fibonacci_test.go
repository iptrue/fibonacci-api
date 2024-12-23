package domain

import (
	"math/big"
	"testing"
)

func TestFibonacciResult_Marshal(t *testing.T) {
	tests := []struct {
		name    string
		result  FibonacciResult
		want    string
		wantErr bool
	}{
		{
			name:    "Marshal int64 value",
			result:  FibonacciResult{Value: 10},
			want:    "10",
			wantErr: false,
		},
		{
			name:    "Marshal big.Int value",
			result:  FibonacciResult{BigValue: big.NewInt(1000)},
			want:    "1000",
			wantErr: false,
		},
		{
			name:    "Marshal zero value",
			result:  FibonacciResult{Value: 0},
			want:    "0",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.result.Marshal()
			if (err != nil) != tt.wantErr {
				t.Errorf("FibonacciResult.Marshal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("FibonacciResult.Marshal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFibonacciResult_Unmarshal(t *testing.T) {

	bigValue, _ := new(big.Int).SetString("1000000000000000000000000000000000", 10)

	tests := []struct {
		name    string
		value   string
		want    FibonacciResult
		wantErr bool
	}{
		{
			name:    "Unmarshal int64 value",
			value:   "10",
			want:    FibonacciResult{Value: 10},
			wantErr: false,
		},
		{
			name:    "Unmarshal big.Int value",
			value:   "1000000000000000000000000000000000",
			want:    FibonacciResult{BigValue: bigValue},
			wantErr: false,
		},
		{
			name:    "Unmarshal empty string",
			value:   "",
			want:    FibonacciResult{},
			wantErr: true,
		},
		{
			name:    "Unmarshal invalid string",
			value:   "invalid",
			want:    FibonacciResult{},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got FibonacciResult
			err := got.Unmarshal(tt.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("FibonacciResult.Unmarshal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err == nil && (got.Value != tt.want.Value || (got.BigValue != nil && got.BigValue.Cmp(tt.want.BigValue) != 0)) {
				t.Errorf("FibonacciResult.Unmarshal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFibonacciInt64(t *testing.T) {
	tests := []struct {
		name string
		n    int64
		want int64
	}{
		{"Fibonacci number for 0", 0, 0},
		{"Fibonacci number for 1", 1, 1},
		{"Fibonacci number for 2", 2, 1},
		{"Fibonacci number for 3", 3, 2},
		{"Fibonacci number for 4", 4, 3},
		{"Fibonacci number for 10", 10, 55},
		{"Fibonacci number for 92", 92, 7540113804746346429},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FibonacciInt64(tt.n)
			if got != tt.want {
				t.Errorf("FibonacciInt64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFibonacciBig(t *testing.T) {

	fib93, _ := new(big.Int).SetString("12200160415121876738", 10)
	fib100, _ := new(big.Int).SetString("354224848179261915075", 10)

	tests := []struct {
		name string
		n    int64
		want *big.Int
	}{
		{"Fibonacci number for 93", 93, fib93},
		{"Fibonacci number for 100", 100, fib100},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FibonacciBig(tt.n)
			if got.Cmp(tt.want) != 0 {
				t.Errorf("FibonacciBig() = %v, want %v", got, tt.want)
			}
		})
	}
}
