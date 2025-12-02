package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, _ := os.Open("input.txt")
	defer func() {
		f.Close()
	}()
	out := run(f)
	fmt.Println(out)
}
func ScanComma(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}
	if i := bytes.IndexByte(data, ','); i >= 0 {
		return i + 1, data[0:i], nil
	}
	if atEOF {
		return len(data), data, nil
	}
	return 0, nil, nil
}

func splitRange(s string) (string, string) {
	var start, end string
	split := strings.SplitN(s, "-", 2)
	start = split[0]
	end = split[1]
	return start, end
}

func findInvalidIds(start string, end string) []int {
	var invalid_ids []int

	lenStart := len(start)
	lenEnd := len(end)

	if lenStart%2 != 0 && lenEnd%2 != 0 {
		return invalid_ids
	}
	startInt, _ := strconv.Atoi(start)
	endInt, _ := strconv.Atoi(end)

	diff := endInt - startInt

	fmt.Printf("Checking range %s to %s with diff %d\n", start, end, diff)

	for i := range diff + 1 {
		id := startInt + i
		idStr := strconv.Itoa(id)

		if len(idStr)%2 != 0 {
			continue
		}

		mid := len(idStr) / 2
		firstHalf := idStr[:mid]
		secondHalf := idStr[mid:]
		if firstHalf == secondHalf {
			fmt.Printf("Checking id %d: %s vs %s\n", id, firstHalf, secondHalf)
			invalid_ids = append(invalid_ids, id)
		}
	}

	return invalid_ids
}

func sumInts(ints []int) int {
	sum := 0
	for _, v := range ints {
		sum += v
	}
	return sum
}

func run(file *os.File) int {
	scanner := bufio.NewScanner(file)
	scanner.Split(ScanComma)
	id_sum := 0
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		start, end := splitRange(line)
		id_sum += sumInts(findInvalidIds(start, end))
	}
	return id_sum
}
