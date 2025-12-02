package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Dial struct {
	Position int
}

func overflow_add(a int, b int, max int) int {
	sum := a + b
	if sum >= max {
		return sum - max
	}
	return sum
}

func underflow_sub(a int, b int, max int) int {
	diff := a - b
	if diff < 0 {
		return diff + max
	}
	return diff
}

func (d *Dial) Rotate(direction string, steps int, max int) {
	switch direction {
	case "R":
		d.Position = overflow_add(d.Position, steps%max, max)
	case "L":
		d.Position = underflow_sub(d.Position, steps%max, max)
	}
}

func main() {
	f, _ := os.Open("input.txt")
	defer func() {
		f.Close()
	}()
	out := run(f)
	fmt.Println(out)
}

func run(file *os.File) int {
	dial := Dial{Position: 50}
	scanner := bufio.NewScanner(file)
	zero_count := 0
	for scanner.Scan() {
		line := scanner.Text()
		direction := line[0:1]
		steps, _ := strconv.Atoi(line[1:])
		dial.Rotate(direction, steps, 100)
		fmt.Printf("Rotating %d steps %s to point to pos %d\n", steps, direction, dial.Position)
		if dial.Position == 0 {
			zero_count++
		}
	}
	return zero_count
}
