package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log/slog"
	"os"
	"strconv"
	"strings"
)

var handler = slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
	Level: slog.LevelInfo,
})

var logger = slog.New(
	handler,
)

func scanComma(data []byte, atEOF bool) (advance int, token []byte, err error) {
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

func isRepeatingSubstring(s string, length int) bool {
	if len(s)%length != 0 {
		return false
	}
	substr := s[0:length]
	for i := 0; i <= len(s)-length; i += length {
		logger.Debug("Comparing substrings", "substr", substr, "current", s[i:i+length])
		if s[i:i+length] != substr {
			return false
		}
	}
	return true
}

func findInvalidIds(start string, end string) []int {
	var invalid_ids []int

	startInt, _ := strconv.Atoi(start)
	endInt, _ := strconv.Atoi(end)

	diff := endInt - startInt

	logger.Debug("Checking range", "start", start, "end", end, "diff", diff)

	for i := range diff + 1 {
		id := startInt + i
		idStr := strconv.Itoa(id)
		mid := len(idStr) / 2
		for j := 1; j <= mid; j++ {
			if isRepeatingSubstring(idStr, j) {
				logger.Info("Found repeating substring", "id", idStr, "substring", idStr[0:j], "length", j)
				invalid_ids = append(invalid_ids, id)
				break
			}
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
	scanner.Split(scanComma)
	id_sum := 0
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		start, end := splitRange(line)
		id_sum += sumInts(findInvalidIds(start, end))
	}
	return id_sum
}

func main() {
	f, _ := os.Open("input.txt")
	defer func() {
		f.Close()
	}()
	out := run(f)

	fmt.Println(out)
}
