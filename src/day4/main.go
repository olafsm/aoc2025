package main

import (
	"bufio"
	"fmt"
	"log/slog"
	"os"
	"strings"
)

var handler = slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
	Level: slog.LevelWarn,
})

var logger = slog.New(
	handler,
)

type Space rune

const (
	Empty              Space = '.'
	RollOfPaper        Space = '@'
	CountedRollOfPaper Space = 'x'
)

type Grid [][]Space

func (g *Grid) String() string {
	var sb strings.Builder
	for _, line := range *g {
		sb.WriteString(string(line))
		sb.WriteString("\n")
	}
	return sb.String()
}

func (g *Grid) numberOfNeighborsWithType(x, y int) int {
	directions := [][2]int{
		{-1, -1}, {0, -1}, {1, -1},
		{-1, 0}, {1, 0},
		{-1, 1}, {0, 1}, {1, 1},
	}
	neighborsWithType := 0
	for _, dir := range directions {
		nx, ny := x+dir[0], y+dir[1]
		if nx >= 0 && nx < len((*g)[0]) && ny >= 0 && ny < len(*g) {
			if (*g)[ny][nx] == RollOfPaper || (*g)[ny][nx] == CountedRollOfPaper {
				logger.Info(fmt.Sprintf("route %d, %d", nx, ny))
				neighborsWithType++
			}
		}
	}
	return neighborsWithType
}

type Line []Space

func run(file *os.File) int {
	scanner := bufio.NewScanner(file)
	grid := Grid{}
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		grid = append(grid, Line(line))
	}
	accessibleRolls := 0
	for y := range len(grid[0]) {
		for x := range len(grid) {
			if grid[y][x] == RollOfPaper && grid.numberOfNeighborsWithType(x, y) < 4 {
				grid[y][x] = CountedRollOfPaper
				accessibleRolls++
			}
		}
	}
	return accessibleRolls
}

func main() {
	f, _ := os.Open("input.txt")
	defer func() {
		f.Close()
	}()
	out := run(f)

	fmt.Println(out)
}
