package main

import (
	"aoc_2025/lib"
	"bufio"
	"os"
	"strings"
)

type Interval struct {
	start uint64
	end   uint64
}

func (i *Interval) contains(num uint64) bool {
	return num >= i.start && num <= i.end
}

func (i *Interval) length() uint64 {
	return i.end - i.start + 1
}

func main() {
	file, _ := os.Open("inputs/05.in")
	scanner := bufio.NewScanner(file)

	var intervals []Interval
	var nums []uint64

	parsingRanges := true
	for scanner.Scan() {
		line := scanner.Text()
		if strings.TrimSpace(line) == "" {
			parsingRanges = false
			continue
		}
		if parsingRanges {
			ints := lib.UIntsPlease(line, "-")
			intervals = append(intervals, Interval{ints[0], ints[1]})
		} else {
			nums = append(nums, lib.UIntPlease(line))
		}
	}

	consolidating := true
	for consolidating {
		intervals, consolidating = consolidate(intervals)
	}

	p1 := 0
	for _, num := range nums {
		for _, interval := range intervals {
			if interval.contains(num) {
				p1++
				break
			}
		}
	}
	println("p1:", p1)

	p2 := uint64(0)
	for _, interval := range intervals {
		p2 += interval.length()
	}
	println("p2:", p2)
}

func consolidate(intervals []Interval) ([]Interval, bool) {
	var ret []Interval
	replaced := make(map[int]bool)

	// Loop through all distinct pairs of intervals and check whether they overlap.
	// If they do, mark them as replaced and insert the union interval in the return map.
	for n, i1 := range intervals {
		for m, i2 := range intervals {
			// Don't look at an interval pair if either of its members has been marked as replaced.
			if m <= n || replaced[n] || replaced[m] {
				continue
			}
			if min(i1.end, i2.end) >= max(i1.start, i2.start) {
				ret = append(ret, Interval{min(i1.start, i2.start), max(i1.end, i2.end)})
				replaced[n], replaced[m] = true, true
			}
		}
	}

	for n, interval := range intervals {
		if !replaced[n] {
			ret = append(ret, interval)
		}
	}

	return ret, len(replaced) > 0
}
