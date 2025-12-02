package main

import "os"

func main() {
	f, _ := os.Open("../input/day2.txt")
	defer func() {
		f.Close()
	}()
	run(f)
}

func run(*os.File) int {
	return 0
}
