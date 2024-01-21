package main

import "testing"

func TestFunc1(t *testing.T) {
	// Test cases for Func1
	testCases := []struct {
		a, b, expected int
	}{
		{2, 3, 6},
		{-2, 5, -10},
		{0, 0, 0},
		{4, 7, 28},
	}

	for _, tc := range testCases {
		result := Func1(tc.a, tc.b)
		if result != tc.expected {
			t.Errorf("Func1(%d, %d) = %d; expected %d", tc.a, tc.b, result, tc.expected)
		}
	}
}

func TestFunc2(t *testing.T) {
	// Test cases for Func2
	testCases := []struct {
		a, b, c  int
		s        string
		expected int
	}{
		{2, 3, 4, "lol", 24},
		{1, 2, 3, "bol", 6},
		{4, 5, 2, "yourass", 24},
		{1, 2, 3, "unknown", 0},
	}

	for _, tc := range testCases {
		result := Func2(tc.a, tc.b, tc.c, tc.s)
		if result != tc.expected {
			t.Errorf("Func2(%d, %d, %d, %q) = %d; expected %d", tc.a, tc.b, tc.c, tc.s, result, tc.expected) // ldfsfdsfsdfdsfdsfsdfsf
		}
	}
}
