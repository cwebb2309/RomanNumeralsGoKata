package main

import (
	"testing"
)

type numPartTest struct {
	n        int
	testNum  int
	expected int
}

var numPartTests = []numPartTest{
	{1, 2018, 8}, {10, 2018, 1}, {100, 2018, 0}, {1000, 2018, 2},
}

func TestNumParts(t *testing.T) {
	for _, tt := range numPartTests {
		actual := getDigitFromNumber(tt.n, tt.testNum)
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
