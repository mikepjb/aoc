package main

import (
	"embed"
	"fmt"

	"hypalynx.com/aoc/internal/day01"
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
}
