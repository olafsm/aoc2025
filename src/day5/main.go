package main

import (
	"bufio"
	"fmt"
	"log/slog"
	"os"
	"strconv"
	"strings"
)

var handler = slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
	Level: slog.LevelWarn,
})

var logger = slog.New(
	handler,
)

type IngredientRange struct {
	Min int
	Max int
}

func run(file *os.File) int {
	scanner := bufio.NewScanner(file)
	ranges := make([]IngredientRange, 0)
	parsingRanges := true
	ingredientsInRange := 0
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		println(line)
		if line == "" {
			parsingRanges = false
		}
		if parsingRanges {
			splitLine := strings.SplitN(line, "-", 2)
			min, _ := strconv.Atoi(splitLine[0])
			max, _ := strconv.Atoi(splitLine[1])
			ranges = append(ranges, IngredientRange{min, max})
			continue
		}
		for _, r := range ranges {
			ingredient, _ := strconv.Atoi(line)
			if ingredient >= r.Min && ingredient <= r.Max {
				fmt.Printf("found %s in range %d-%d\n", line, r.Min, r.Max)
				ingredientsInRange++
				break
			}
		}
	}
	return ingredientsInRange
}

func main() {
	f, _ := os.Open("input.txt")
	defer func() {
		f.Close()
	}()
	out := run(f)

	fmt.Println(out)
}
