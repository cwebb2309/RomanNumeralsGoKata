package main

import (
	"testing"
)

type datePartTest struct {
	n        int
	expected int
}

var datePartTests = []datePartTest{
	{1, 8}, {10, 1}, {100, 0}, {1000, 2},
}

func TestDateParts(t *testing.T) {
	for _, tt := range datePartTests {
		actual := getDigitFromNumber(tt.n, 2018)
		if actual != tt.expected {
			t.Errorf("Expected : %d", tt.expected)
		}
	}
}

type romanTest struct {
	n        int
	expected string
}

var romanTests = []romanTest{
	{5, "V"}, {1900, "MCM"}, {499, "CDXCIX"},
}

func TestDecimalToRoman(t *testing.T) {
	for _, tt := range romanTests {
		actual := decimalToRoman(tt.n)
		if actual != tt.expected {
			t.Errorf("Expected : %s", tt.expected)
		}
	}
}
