package main

import (
	"embed"
	"fmt"

	"hypalynx.com/aoc/internal/day01"
	"hypalynx.com/aoc/internal/day11"
)

//go:embed input
var inputDir embed.FS

func main() {
	fmt.Println("Hello AOC!")

	file01, err := inputDir.ReadFile("input/day01.txt")

	if err == nil {
		input01 := string(file01)
		fmt.Printf("Day 01, Part 01: %v\n", day01.Part01(string(input01)))
		fmt.Printf("Day 01, Part 02: %v\n", day01.Part02(string(input01)))
	}

	file11, err := inputDir.ReadFile("input/day11.txt")
	if err == nil {
		input11 := string(file11)
		fmt.Printf("Day 11, Part 01: %v\n", day11.Both(string(input11), 25))
		fmt.Printf("Day 11, Part 02: %v\n", day11.Both(string(input11), 75))
	}
}
