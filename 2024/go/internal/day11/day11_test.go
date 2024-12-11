package day11_test

import (
	"reflect"
	"testing"

	"hypalynx.com/aoc/internal/day11"
)

func TestPart01Sample(t *testing.T) {
	input := "0 1 10 99 999"
	expectation := "1 2024 1 0 9 9 2021976"

	result := day11.Part01(input)
	if !reflect.DeepEqual(result, expectation) {
		t.Fatalf("%v is not %v", result, expectation)
	}
}
