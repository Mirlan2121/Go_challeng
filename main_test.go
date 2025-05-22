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

func TestCalculator(t *testing.T) {
	sum, sub, mul, div := Calculator()
	if sum != 200 || sub != 0 || mul != 10000 || div != 1 {
		t.Errorf("Calculator() = %d, %d, %d, %d, ожидалось 200, 0, 10000, 1",
			sum, sub, mul, div)
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
