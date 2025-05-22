package main

import (
	"testing"
)

// ... (существующие тесты остаются без изменений)

func TestSub(t *testing.T) {
	tests := []struct {
		name     string
		a, b     int
		expected int
	}{
		{"Positive", 5, 3, 2},
		{"Negative", 3, 5, -2},
		{"Zero", 5, 5, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Sub(tt.a, tt.b); got != tt.expected {
				t.Errorf("Sub(%d, %d) = %d, want %d", tt.a, tt.b, got, tt.expected)
			}
		})
	}
}

func TestMul(t *testing.T) {
	tests := []struct {
		name     string
		a, b     int
		expected int
	}{
		{"Positive", 5, 3, 15},
		{"WithZero", 5, 0, 0},
		{"Negative", -3, 4, -12},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Mul(tt.a, tt.b); got != tt.expected {
				t.Errorf("Mul(%d, %d) = %d, want %d", tt.a, tt.b, got, tt.expected)
			}
		})
	}
}

func TestDiv(t *testing.T) {
	tests := []struct {
		name     string
		a, b     int
		expected int
	}{
		{"Normal", 6, 3, 2},
		{"Floor", 5, 2, 2},
		{"ByZero", 5, 0, 0}, // Проверка деления на ноль
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Div(tt.a, tt.b); got != tt.expected {
				t.Errorf("Div(%d, %d) = %d, want %d", tt.a, tt.b, got, tt.expected)
			}
		})
	}
}

func TestPowerZeroZero(t *testing.T) {
	result := Power(0, 0)
	if result != 1.0 {
		t.Errorf("Power(0, 0) = %f, want 1.0", result)
	}
}
