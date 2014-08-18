package main

import "testing"

func TestBelow10(t *testing.T) {
	expected := 23
	if result := SumOfMultiples(10); result != expected {
		t.Errorf("SumOfMultiples(10) = %v, want %v", result, expected)
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
