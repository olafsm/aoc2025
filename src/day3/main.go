package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func findHigestJoltage(bank string) int {
	fmt.Printf("bank: %s\n", bank)
	highest := '0'
	highestAfter := '0'
	for i, ch := range bank {
		if ch > highest && i != len(bank)-1 {
			highestAfter = '0'
			highest = ch
		} else if ch > highestAfter {
			highestAfter = ch
		}
	}
	fmt.Printf("Found highest and highestAfter: %d, %d\n", highest-'0', highestAfter-'0')
	val, _ := strconv.Atoi(fmt.Sprintf("%d%d", int(highest-'0'), int(highestAfter-'0')))
	return val
}
func run(file *os.File) int {
	scanner := bufio.NewScanner(file)
	joltageSum := 0
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		joltageSum += findHigestJoltage(line)
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
