package day01

import (
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"
)

func Part01(input string) string {
	inputLines := strings.Split(input, "\n")

	leftNums := []int{}
	rightNums := []int{}

	for _, line := range inputLines {
		numbers := strings.Fields(line)
		if len(numbers) == 0 {
			break
		}

		ln, err := strconv.Atoi(numbers[0])
		if err != nil {
			log.Fatalf("could not convert '%v'", numbers[0])
			return "-1"
		}
		leftNums = append(leftNums, ln)

		rn, err := strconv.Atoi(numbers[1])
		if err != nil {
			log.Fatalf("could not convert '%v'", numbers[1])
			return "-1"
		}
		rightNums = append(rightNums, rn)
	}

	sort.Ints(leftNums)
	sort.Ints(rightNums)

	ds := distances(leftNums, rightNums)

	return fmt.Sprint(sum(ds))
}

func distances(leftNums []int, rightNums []int) []int {
	ds := []int{}
	for i, ln := range leftNums {
		var d int
		rn := rightNums[i]
		if ln > rn {
			d = ln - rn

		} else {
			d = rn - ln
		}
		ds = append(ds, d)
	}
	return ds
}

func sum(ds []int) int {
	total := 0
	for _, d := range ds {
		total += d
	}
	return total
}
