package day11

import (
	"fmt"
	"strconv"
	"strings"
    "sync"
)

// broken async brute force solution
// I think the only way really is to take each initial stone and split it, iterate 75 times and
// then count that.. store the count and work your way through to avoid consuming too much memory..
// doesn't address the cpu time problem though, it's taking an age to calculate at 75 blinks

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
		stones = papply(blink, stones)
		numberOfBlinks++
	}

	return toString(stones)
}

func apply(fn func(stone int) []int, stones []int) []int {
	result := []int{}
	for _, s := range stones {
		result = append(result, blink(s)...)
	}
	return result
}

func papply(fn func(stone int) []int, stones []int) []int {
	result := []int{}
    ch := make(chan int)
    var mu sync.Mutex
    var wg sync.WaitGroup

    wg.Add(len(stones))

	for _, stone := range stones {
        go func() {
            defer wg.Done()
            for _, bs := range blink(stone) {
                ch <- bs
            }
        }()
	}

    go func() {
        for bs := range ch {
            mu.Lock()
            result = append(result, bs)
            mu.Unlock()
        }
    }()


    wg.Wait()
    close(ch)

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
