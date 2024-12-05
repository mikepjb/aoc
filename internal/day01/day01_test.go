package day01

import (
	"reflect"
	"sort"
	"testing"
)

func TestListOrdering(t *testing.T) {
	list := []int{3, 4, 2, 1, 3, 3}
	expectation := []int{1, 2, 3, 3, 3, 4}

	sort.Ints(list)
	if !reflect.DeepEqual(list, expectation) {
		t.Fatalf("%v is not %v", list, expectation)
	}
}

func TestPart01Sample(t *testing.T) {
	sample := `3   4
4   3
2   5
1   3
3   9
3   3`

	Part01(sample)
}
