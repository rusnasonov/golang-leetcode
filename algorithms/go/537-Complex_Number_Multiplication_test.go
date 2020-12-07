package main

import "testing"

func TestComplexNumberMultiply(t *testing.T) {
	expected := "0+-2i"
	if res := complexNumberMultiply("1+-1i", "1+-1i"); res != expected {
		t.Errorf("got %s, expected %s", res, expected)
	}
	expected = "0+2i"
	if res := complexNumberMultiply("1+1i", "1+1i"); res != expected {
		t.Errorf("got %s, expected %s", res, expected)
	}
}

func TestParseComplex(t *testing.T) {
	testCase := "1+1i"
	real, imag, err := parseComplex(testCase)
	if real != 1 || imag != 1 || err != nil {
		t.Errorf("got %d+%di, expected %s; err: %s", real, imag, testCase, err)
	}
	testCase = "2+-2i"
	real, imag, err = parseComplex(testCase)
	if real != 2 || imag != -2 || err != nil {
		t.Errorf("got %d+%di, expected %s; err: %s", real, imag, testCase, err)
	}
	testCase = "2+2"
	real, imag, err = parseComplex(testCase)
	if err != nil {
		t.Errorf("got %d+%di, expected %s; err: %s", real, imag, testCase, err)
	}
}
