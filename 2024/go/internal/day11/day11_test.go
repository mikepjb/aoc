package day11_test

import (
	"reflect"
	"testing"

	"hypalynx.com/aoc/internal/day11"
)

func TestPart01Sample(t *testing.T) {
	input := "0 1 10 99 999"
	expectation := "1 2024 1 0 9 9 2021976"

	result := day11.Observe(input, 1)
	if !reflect.DeepEqual(result, expectation) {
		t.Fatalf("%v is not %v", result, expectation)
	}
}

func TestPart01SecondSample(t *testing.T) {
	input := "125 17"
	expectation := "2097446912 14168 4048 2 0 2 4 40 48 2024 40 48 80 96 2 8 6 7 6 0 3 2"

	result := day11.Observe(input, 6)
	if !reflect.DeepEqual(result, expectation) {
		t.Fatalf("%v is not %v", result, expectation)
	}

	firstCount := day11.Part01(input, 6)
	if firstCount != 22 {
		t.Fatalf("%v is not 22", firstCount)
	}

	secondCount := day11.Part01(input, 25)
	if secondCount != 55312 {
		t.Fatalf("%v is not 22", secondCount)
	}
}
