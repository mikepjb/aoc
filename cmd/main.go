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
		day01.Part01(string(input01))
	}
}
