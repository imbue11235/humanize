package humanize

import (
	"fmt"
	"math"
	"strconv"
)

// Fraction takes a decimal (float) value and converts
// it into a string representation of a fractional
// value.
//
//		f := humanize.Fraction(0.1)
//		fmt.Println(f) => "1/10"
//
func Fraction(decimal float64) string {
	if math.IsNaN(decimal) {
		return "NaN"
	}

	precision := float64(calculatePrecision(decimal))
	if precision == 0 {
		return strconv.FormatFloat(decimal, 'f', 0, 64)
	}

	factor := math.Pow(10, precision)
	denominator := int(factor)
	numerator := int(decimal * factor)

	// to see if the fraction can be further reduced,
	// we will calculate the greatest common factor
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

	// if wholes are negative
	if wholes < 0 {
		numerator -= wholes * denominator
		return fmt.Sprintf("%d %d/%d", wholes, -numerator, denominator)
	}

	return fmt.Sprintf("%d/%d", numerator, denominator)
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
	// if number is whole, we can't calculate precision
	if decimal == math.Trunc(decimal) {
		return 0
	}

	curr, precision := 1.0, 0
	for math.Round(decimal*curr)/curr != decimal {
		curr *= 10
		precision++
	}

	return precision
}
