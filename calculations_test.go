package main

import (
	"math"
	"testing"
)

func TestCalculateProfit(t *testing.T) {
	tests := []struct {
		name     string
		revenue  int
		expenses int
		taxRate  float64
		expected CalculationResult
	}{
		{
			name:     "Basic calculation",
			revenue:  10000000, // €100,000
			expenses: 5000000,  // €50,000
			taxRate:  19.5,
			expected: CalculationResult{
				EBT:         5000000,
				Profit:      4025000,
				ProfitRatio: 0.4025,
			},
		},
		{
			name:     "Zero tax rate",
			revenue:  1000000, // €10,000
			expenses: 600000,  // €6,000
			taxRate:  0.0,
			expected: CalculationResult{
				EBT:         400000,
				Profit:      400000,
				ProfitRatio: 0.4,
			},
		},
		{
			name:     "High tax rate",
			revenue:  2000000, // €20,000
			expenses: 1000000, // €10,000
			taxRate:  50.0,
			expected: CalculationResult{
				EBT:         1000000,
				Profit:      500000,
				ProfitRatio: 0.25,
			},
		},
		{
			name:     "Loss scenario (expenses > revenue)",
			revenue:  500000, // €5,000
			expenses: 800000, // €8,000
			taxRate:  20.0,
			expected: CalculationResult{
				EBT:         -300000,
				Profit:      -240000, // Loss is reduced by tax benefit
				ProfitRatio: -0.48,
			},
		},
		{
			name:     "Zero revenue",
			revenue:  0,
			expenses: 100000, // €1,000
			taxRate:  20.0,
			expected: CalculationResult{
				EBT:         -100000,
				Profit:      -80000,
				ProfitRatio: 0, // Division by zero handled
			},
		},
		{
			name:     "Zero expenses",
			revenue:  1000000, // €10,000
			expenses: 0,
			taxRate:  25.0,
			expected: CalculationResult{
				EBT:         1000000,
				Profit:      750000,
				ProfitRatio: 0.75,
			},
		},
		{
			name:     "Decimal tax rate",
			revenue:  5000000, // €50,000
			expenses: 3000000, // €30,000
			taxRate:  19.25,
			expected: CalculationResult{
				EBT:         2000000,
				Profit:      1615000,
				ProfitRatio: 0.323,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := CalculateProfit(tt.revenue, tt.expenses, tt.taxRate)

			if result.EBT != tt.expected.EBT {
				t.Errorf("EBT = %d, want %d", result.EBT, tt.expected.EBT)
			}

			// Use a small epsilon for float comparisons
			if math.Abs(result.Profit-tt.expected.Profit) > 0.01 {
				t.Errorf("Profit = %.2f, want %.2f", result.Profit, tt.expected.Profit)
			}

			if math.Abs(result.ProfitRatio-tt.expected.ProfitRatio) > 0.001 {
				t.Errorf("ProfitRatio = %.6f, want %.6f", result.ProfitRatio, tt.expected.ProfitRatio)
			}
		})
	}
}

func TestFormatEuro(t *testing.T) {
	tests := []struct {
		name     string
		input    float64
		expected string
	}{
		{
			name:     "Positive amount with thousands",
			input:    50000.00,
			expected: "50,000.00",
		},
		{
			name:     "Positive amount with millions",
			input:    1500000.50,
			expected: "1,500,000.50",
		},
		{
			name:     "Small positive amount",
			input:    123.45,
			expected: "123.45",
		},
		{
			name:     "Zero amount",
			input:    0.00,
			expected: "0.00",
		},
		{
			name:     "Negative amount",
			input:    -2500.75,
			expected: "-2,500.75",
		},
		{
			name:     "Large negative amount",
			input:    -1234567.89,
			expected: "-1,234,567.89",
		},
		{
			name:     "Amount with no decimals",
			input:    1000.00,
			expected: "1,000.00",
		},
		{
			name:     "Single digit",
			input:    5.00,
			expected: "5.00",
		},
		{
			name:     "Rounding edge case",
			input:    99.999,
			expected: "100.00",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := FormatEuro(tt.input)
			if result != tt.expected {
				t.Errorf("FormatEuro(%.2f) = %s, want %s", tt.input, result, tt.expected)
			}
		})
	}
}

func TestFormatResult(t *testing.T) {
	result := CalculationResult{
		EBT:         5000000,
		Profit:      4025000,
		ProfitRatio: 0.4025,
	}

	expected := "EBT: 5000000 cents (50,000.00 euros)\nProfit: 4025000 cents (40,250.00 euros)\nProfit Ratio: 40.25%"
	formatted := FormatResult(result)

	if formatted != expected {
		t.Errorf("FormatResult() = %s, want %s", formatted, expected)
	}
}

// Benchmark tests for performance
func BenchmarkCalculateProfit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CalculateProfit(10000000, 5000000, 19.5)
	}
}

func BenchmarkFormatEuro(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FormatEuro(50000.00)
	}
}

// Edge case tests
func TestCalculateProfitEdgeCases(t *testing.T) {
	t.Run("Maximum int values", func(t *testing.T) {
		result := CalculateProfit(math.MaxInt32, 0, 0)
		if result.EBT != math.MaxInt32 {
			t.Errorf("Expected EBT to be MaxInt32, got %d", result.EBT)
		}
	})

	t.Run("Negative values", func(t *testing.T) {
		result := CalculateProfit(-1000, -500, 20)
		expected := CalculationResult{
			EBT:         -500,
			Profit:      -400,
			ProfitRatio: 0.4, // -400 / -1000 = 0.4
		}

		if result.EBT != expected.EBT {
			t.Errorf("EBT = %d, want %d", result.EBT, expected.EBT)
		}
		if math.Abs(result.Profit-expected.Profit) > 0.01 {
			t.Errorf("Profit = %.2f, want %.2f", result.Profit, expected.Profit)
		}
		if math.Abs(result.ProfitRatio-expected.ProfitRatio) > 0.001 {
			t.Errorf("ProfitRatio = %.6f, want %.6f", result.ProfitRatio, expected.ProfitRatio)
		}
	})

	t.Run("Very high tax rate", func(t *testing.T) {
		result := CalculateProfit(1000000, 500000, 100)
		expected := CalculationResult{
			EBT:         500000,
			Profit:      0, // 100% tax rate
			ProfitRatio: 0,
		}

		if result.EBT != expected.EBT {
			t.Errorf("EBT = %d, want %d", result.EBT, expected.EBT)
		}
		if math.Abs(result.Profit-expected.Profit) > 0.01 {
			t.Errorf("Profit = %.2f, want %.2f", result.Profit, expected.Profit)
		}
		if math.Abs(result.ProfitRatio-expected.ProfitRatio) > 0.001 {
			t.Errorf("ProfitRatio = %.6f, want %.6f", result.ProfitRatio, expected.ProfitRatio)
		}
	})
}
