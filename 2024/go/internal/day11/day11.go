package day11

import (
	"fmt"
	"strconv"
	"strings"
)

func Both(input string, blinks int) int {
	result := Observe(strings.TrimSpace(input), blinks)
	stones := strings.Split(result, " ")
	return len(stones)
}

func Observe(input string, blinks int) string {
	stoneStrs := strings.Split(input, " ")
	stones := []int{}

	for _, s := range stoneStrs {
		i, err := strconv.Atoi(s)
		if err == nil {
			stones = append(stones, i)
		}
	}

	numberOfBlinks := 0
	targetNumberOfBlinks := blinks

	for numberOfBlinks < targetNumberOfBlinks {
		stones = apply(blink, stones)
		numberOfBlinks++
	}

	return toString(stones)
}

func Part02(input string) string {
	return "no"
}

func apply(fn func(stone int) []int, stones []int) []int {
	result := []int{}
	for _, s := range stones {
		result = append(result, blink(s)...)
	}
	return result
}

func toString(stones []int) string {
	strs := []string{}
	for _, s := range stones {
		strs = append(strs, fmt.Sprintf("%v", s))
	}
	return strings.Join(strs, " ")
}

func blink(stone int) []int {
	stoneStr := fmt.Sprintf("%v", stone)
	stoneLen := len(stoneStr)
	switch {
	case stone == 0:
		return []int{1}
	case stoneLen%2 == 0:
		halfway := stoneLen / 2
		first, _ := strconv.Atoi(stoneStr[0:halfway])
		second, _ := strconv.Atoi(stoneStr[halfway:stoneLen])
		return []int{first, second}
	default:
		return []int{stone * 2024}
	}
}
