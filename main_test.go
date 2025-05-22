package main

import (
	"testing"
)

func TestAdd(t *testing.T) {
	tests := []struct {
		name     string
		numbers  []int
		expected int
	}{
		{"Positive numbers", []int{1, 2, 3}, 6},
		{"Negative numbers", []int{-1, -2, -3}, -6},
		{"Mixed numbers", []int{-1, 2, -3}, -2},
		{"Empty list", []int{}, 0}, // крайний случай: пустой список
		{"Single number", []int{42}, 42},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Add(tt.numbers...)
			if result != tt.expected {
				t.Errorf("Add(%v) = %d, expected %d", tt.numbers, result, tt.expected)
			}
		})
	}
}

func TestPower(t *testing.T) {
	tests := []struct {
		name     string
		n        float64 // Изменено на float64
		power    int
		expected float64 // Изменено на float64
	}{
		{"Positive exponent", 2.0, 3, 8.0},
		{"Zero exponent", 5.0, 0, 1.0},
		{"Negative exponent", 2.0, -3, 0.125}, // Теперь поддерживается
		{"Zero base", 0.0, 5, 0.0},
		{"One to any power", 1.0, 100, 1.0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Power(tt.n, tt.power)
			if result != tt.expected {
				t.Errorf("Power(%v, %d) = %v, expected %v", tt.n, tt.power, result, tt.expected)
			}
		})
	}
}

func TestHandler(t *testing.T) {
	// Можно добавить HTTP-тесты с помощью httptest, но это выходит за рамки примера
}

// main_test.go
func TestHello(t *testing.T) {
	tests := []struct {
		name     string
		expected string
	}{
		{"Default case", "Правильно"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hello(); got != tt.expected {
				t.Errorf("hello() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestAddEmpty(t *testing.T) {
	result := Add() // Пустой список
	if result != 0 {
		t.Errorf("add() = %d, ожидалось 0", result)
	}
}

func TestPowerZeroBase(t *testing.T) {
	result := Power(0, 5)
	if result != 0 {
		t.Errorf("Power(0,5) = %f, ожидалось 0", result)
	}
}

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

func TestSum(t *testing.T) {
	if Add(100, 100) != 200 {
		t.Error("Ошибка: Sum(100, 100) != 200")
	}
}
