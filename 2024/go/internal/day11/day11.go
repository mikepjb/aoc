package day11

import (
	"fmt"
	"strconv"
	"strings"
)

func Part01(input string, blinks int) int {
	result := Observe(strings.TrimSpace(input), blinks)
	stones := strings.Split(result, " ")
	return len(stones)
}

func Observe(input string, blinks int) string {
	stoneStrs := strings.Split(input, " ")
	stones := []int64{}

	for _, s := range stoneStrs {
		i, err := strconv.ParseInt(s, 10, 64)
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

func apply(fn func(stone int64) []int64, stones []int64) []int64 {
	result := []int64{}
	for _, s := range stones {
		result = append(result, blink(s)...)
	}
	return result
}

func toString(stones []int64) string {
	strs := []string{}
	for _, s := range stones {
		strs = append(strs, fmt.Sprintf("%v", s))
	}
	return strings.Join(strs, " ")
}

func blink(stone int64) []int64 {
	stoneStr := fmt.Sprintf("%v", stone)
	stoneLen := len(stoneStr)
	switch {
	case stone == 0:
		return []int64{1}
	case stoneLen%2 == 0:
		halfway := stoneLen / 2
		first, _ := strconv.ParseInt(stoneStr[0:halfway], 10, 64)
		second, _ := strconv.ParseInt(stoneStr[halfway:stoneLen], 10, 64)
		return []int64{first, second}
	default:
		return []int64{stone * 2024}
	}
}
