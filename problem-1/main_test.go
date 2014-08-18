package main

import "testing"

func TestBelow10(t *testing.T) {
	expected := 23
	if result := SumOfMultiples(10); result != expected {
		t.Errorf("SumOfMultiples(10) = %v, want %v", result, expected)
	}
}

var triangleTests = []struct {
	input    int
	expected int
}{
	{1, 1},
	{2, 3},
	{3, 6},
	{4, 10},
	{5, 15},
	{6, 21},
	{7, 28},
	{8, 36},
	{9, 45},
	{10, 55},
}

func TestTriangleNumber(t *testing.T) {
	for _, tt := range triangleTests {
		if result := TriangleNumber(tt.input); result != tt.expected {
			t.Errorf(
				"TriangleNumber(%v) = %v, want %v",
				tt.input,
				result,
				tt.expected,
			)
		}
	}
}

func BenchmarkBelow10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SumOfMultiples(10)
	}
}

func BenchmarkBelow1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SumOfMultiples(1000)
	}
}
