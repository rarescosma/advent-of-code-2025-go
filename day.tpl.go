package main

import (
	"bufio"
	"github.com/davecgh/go-spew/spew"
	"os"
)

func main() {
	file, _ := os.Open("inputs/${DAY}.ex")

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		spew.Dump(scanner.Text())
	}
}
