package day01

import (
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"
)

func Part01(input string) string {
	leftNums, rightNums := prepare(input)
	ds := distances(leftNums, rightNums)

	return fmt.Sprint(sum(ds))
}

func Part02(input string) string {
	leftNums, rightNums := prepare(input)
	countMap := count(rightNums)
	return fmt.Sprint(similartyScore(leftNums, countMap))
}

func prepare(input string) ([]int, []int) {
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
			return leftNums, rightNums
		}
		leftNums = append(leftNums, ln)

		rn, err := strconv.Atoi(numbers[1])
		if err != nil {
			log.Fatalf("could not convert '%v'", numbers[1])
			return leftNums, rightNums
		}
		rightNums = append(rightNums, rn)
	}

	sort.Ints(leftNums)
	sort.Ints(rightNums)

	return leftNums, rightNums
}

func count(rightNums []int) map[int]int {
	countMap := make(map[int]int)
	for _, rn := range rightNums {
		countMap[rn] = countMap[rn] + 1
	}

	return countMap
}

func similartyScore(leftNums []int, countMap map[int]int) int {
	score := 0
	for _, ln := range leftNums {
		if countMap[ln] != 0 {
			score = score + (ln * countMap[ln])
		}
	}
	return score
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
