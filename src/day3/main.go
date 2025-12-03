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

func findHighestDigitInRestOfUsableBank(bank string) (rune, int) {
	highest := '0'
	highestPos := 0
	for i, ch := range bank {
		if ch > highest {
			highest = ch
			highestPos = i
		}
	}
	return highest, highestPos
}

func findHighestJoltageInBank(bank string) int {
	sumDigits := []rune{}
	lastIndex := 0
	for i := 12; i > 0; i-- {
		usableBank := bank[lastIndex : len(bank)-i+1]
		logger.Info("Searching for substring", "len", i, "bank", bank, "bankLen", len(bank), "searchableBank", bank[:len(bank)-i], "rest", bank[len(bank)-i:], "restLen", len(bank[len(bank)-i:]))
		val, posInRest := findHighestDigitInRestOfUsableBank(usableBank)
		logger.Info("  -> Found highest digit", "digit", string(val), "pos", posInRest)
		lastIndex = lastIndex + posInRest + 1
		sumDigits = append(sumDigits, val)
	}
	sum, _ := strconv.Atoi(string(sumDigits))
	return sum
}

func run(file *os.File) int {
	scanner := bufio.NewScanner(file)
	joltageSum := 0
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		joltageSum += findHighestJoltageInBank(line)
	}
	return joltageSum
}

func main() {
	f, _ := os.Open("input.txt")
	defer func() {
		f.Close()
	}()
	out := run(f)

	fmt.Println(out)
}
