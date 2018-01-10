package main

import (
	"testing"
)

func TestGetDigitFromNumbers1s(t *testing.T) {

	var rst int
	rst = getDigitFromNumber(1, 2018)

	if rst != 8 {
		t.Error("Should return 8")
	}

}

func TestGetDigitFromNumbers10s(t *testing.T) {

	var rst int
	rst = getDigitFromNumber(10, 2018)

	if rst != 1 {
		t.Error("Should return 1")
	}

}

func TestGetDigitFromNumbers100s(t *testing.T) {

	var rst int
	rst = getDigitFromNumber(100, 2018)

	if rst != 0 {
		t.Error("Should return 0")
	}

}

func TestGetDigitFromNumbers1000s(t *testing.T) {

	var rst int
	rst = getDigitFromNumber(1000, 2018)

	if rst != 2 {
		t.Error("Should return 2")
	}

}

func TestDecimalToRoman5(t *testing.T) {

	var rst string
	rst = decimalToRoman(5)

	if rst != "V" {
		t.Error("Should return V")
	}

}

func TestDecimalToRoman1900(t *testing.T) {

	var rst string
	rst = decimalToRoman(1900)

	if rst != "MCM" {
		t.Error("Should return MCM")
	}

}

func TestDecimalToRoman499(t *testing.T) {

	var rst string
	rst = decimalToRoman(499)

	if rst != "CDXCIX" {
		t.Error("Should return CDXCIX")
	}

}
