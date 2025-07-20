package main

import (
	"fmt"
	"math"
	"strconv"
)

// CalculationResult holds the results of profit calculations
type CalculationResult struct {
	EBT         int     // Earnings Before Tax in cents
	Profit      float64 // Profit after tax in cents
	ProfitRatio float64 // Profit ratio as a decimal (0.0-1.0)
}

// CalculateProfit performs the main profit calculation
func CalculateProfit(revenue, expenses int, taxRate float64) CalculationResult {
	ebt := revenue - expenses
	profit := float64(ebt) * (1 - taxRate/100)

	var ratio float64
	if revenue != 0 {
		ratio = profit / float64(revenue)
	}

	return CalculationResult{
		EBT:         ebt,
		Profit:      profit,
		ProfitRatio: ratio,
	}
}

// FormatEuro formats a value in cents to a euro string with proper formatting
func FormatEuro(val float64) string {
	// Round to 2 decimal places to avoid floating-point precision issues
	rounded := math.Round(val*100) / 100
	intPart := int64(rounded)
	fracPart := int64(math.Round((rounded - float64(intPart)) * 100))
	if fracPart < 0 {
		fracPart = -fracPart
	}
	intStr := strconv.FormatInt(intPart, 10)
	n := len(intStr)
	withCommas := ""
	for i, c := range intStr {
		if i != 0 && (n-i)%3 == 0 {
			withCommas += ","
		}
		withCommas += string(c)
	}
	return fmt.Sprintf("%s.%02d", withCommas, fracPart)
}

// FormatResult formats the calculation result into a human-readable string
func FormatResult(result CalculationResult) string {
	return fmt.Sprintf(
		"EBT: %.0f cents (%s euros)\nProfit: %.0f cents (%s euros)\nProfit Ratio: %.2f%%",
		float64(result.EBT), FormatEuro(float64(result.EBT)/100),
		result.Profit, FormatEuro(result.Profit/100),
		result.ProfitRatio*100,
	)
}
