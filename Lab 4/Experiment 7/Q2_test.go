package main

import "testing"

func TestDivide(t *testing.T) {
	result, err := divide(10, 2)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if result != 5 {
		t.Errorf("expected 5, got %v", result)
	}

	_, err = divide(10, 0)

	if err == nil {
		t.Error("expected division by zero error")
	}
}

func TestSquareRoot(t *testing.T) {
	result, err := squareRoot(25)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if result != 5 {
		t.Errorf("expected 5, got %v", result)
	}

	_, err = squareRoot(-25)

	if err == nil {
		t.Error("expected negative number error")
	}
}

func TestParseNumber(t *testing.T) {
	result, err := parseNumber("123")

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if result != 123 {
		t.Errorf("expected 123, got %v", result)
	}

	_, err = parseNumber("")

	if err == nil {
		t.Error("expected empty input error")
	}

	_, err = parseNumber("hello")

	if err == nil {
		t.Error("expected invalid input error")
	}
}