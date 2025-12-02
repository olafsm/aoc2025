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

func overflow_add(a int, b int, max int) (int, bool) {
	sum := a + b
	if sum >= max {
		return sum - max, true
	}
	return sum, false
}

func underflow_sub(a int, b int, max int) (int, bool) {
	diff := a - b
	if diff < 0 {
		return diff + max, true
	}
	return diff, false
}

func (d *Dial) Rotate(direction string, steps int, max int) int {
	times_passed_zero := steps / max
	remainder := steps % max
	wrapped := false
	previous_position := d.Position
	switch direction {
	case "R":
		d.Position, wrapped = overflow_add(d.Position, remainder, max)
		if d.Position == 0 {
			return times_passed_zero + 1
		}
		if wrapped {
			return times_passed_zero + 1
		}

	case "L":
		d.Position, wrapped = underflow_sub(d.Position, remainder, max)
		if d.Position == 0 {
			return times_passed_zero + 1
		}
		if wrapped && previous_position != 0 {
			times_passed_zero += 1
		}
	}
	return times_passed_zero
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
		times_passed_zero := dial.Rotate(direction, steps, 100)

		fmt.Printf("Rotating %s%d to point at %d, passing zero %d times\n", direction, steps, dial.Position, times_passed_zero)
		zero_count += times_passed_zero
	}
	return zero_count
}
