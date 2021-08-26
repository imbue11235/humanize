package humanize

import (
	"fmt"
	"math"
)

// Fraction takes a decimal (float) value and converts
// it into a string representation of a fractional
// value.
// Example:
//
//		f := Fraction(0.1)
//		fmt.Println(f) => "1/10"
func Fraction(decimal float64) string {
	precision := float64(calculatePrecision(decimal))
	if precision == 0 {
		return fmt.Sprintf("%0.f", decimal)
	}

	factor := math.Pow(10, precision)
	denominator := int(factor)
	numerator := int(decimal * factor)

	gcf := calculateGCF(numerator, denominator)

	if gcf > 1 {
		denominator /= gcf
		numerator /= gcf
	}

	// count amount of "wholes" (integers) the fraction has
	// e.g. 6/2 => 3 integers
	wholes := numerator / denominator
	if wholes > 0 {
		numerator -= wholes * denominator
		return fmt.Sprintf("%d %d/%d", wholes, numerator, denominator)
	}

	return fmt.Sprintf("%d/%d", numerator, denominator)
}

func round(val float64, precision int) float64 {
	return math.Round(val*(math.Pow10(precision))) / math.Pow10(precision)
}

// calculateGCF calculates the greatest common factor
// from two integers, using the euclidean algorithm
// https://en.wikipedia.org/wiki/Greatest_common_divisor#Euclidean_algorithm
func calculateGCF(a, b int) int {
	for b != 0 {
		m := b
		b = a % b
		a = m
	}

	return a
}

// calculatePrecision counts how many trailing decimal places
// the decimal value has
func calculatePrecision(decimal float64) int {
	curr, precision := 1.0, 0

	for math.Round(decimal*curr)/curr != decimal {
		curr *= 10
		precision++
	}

	return precision
}
