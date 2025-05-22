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
			result := Add(tt.numbers)
			if result != tt.expected {
				t.Errorf("Add(%v) = %d, expected %d", tt.numbers, result, tt.expected)
			}
		})
	}
}

func TestPower(t *testing.T) {
	tests := []struct {
		name     string
		n        int
		power    int
		expected int
	}{
		{"Positive exponent", 2, 3, 8},
		{"Zero exponent", 5, 0, 1},                    // крайний случай: степень 0
		{"Negative exponent (unsupported)", 2, -3, 1}, // (в текущей реализации не поддерживается)
		{"Zero base", 0, 5, 0},
		{"One to any power", 1, 100, 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Power(tt.n, tt.power)
			if result != tt.expected {
				t.Errorf("Power(%d, %d) = %d, expected %d", tt.n, tt.power, result, tt.expected)
			}
		})
	}
}

func TestHandler(t *testing.T) {
	// Можно добавить HTTP-тесты с помощью httptest, но это выходит за рамки примера
}
