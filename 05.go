package main

import (
	"bufio"
	"os"
	"strconv"
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

func intPlease(s string) uint64 {
	ret, _ := strconv.ParseUint(s, 10, 0)
	return ret
}

func main() {
	file, _ := os.Open("inputs/05.in")
	scanner := bufio.NewScanner(file)

	intervals := make(map[int]Interval)
	var nums []uint64

	parsingRanges := true
	index := 0
	for scanner.Scan() {
		line := scanner.Text()
		if strings.TrimSpace(line) == "" {
			parsingRanges = false
			continue
		}
		if parsingRanges {
			parts := strings.Split(line, "-")
			intervals[index] = Interval{intPlease(parts[0]), intPlease(parts[1])}
			index++
		} else {
			nums = append(nums, intPlease(line))
		}
	}

	p1 := 0
	for _, num := range nums {
		for _, ival := range intervals {
			if ival.contains(num) {
				p1++
				break
			}
		}
	}
	println("p1:", p1)

	consolidating := true
	for consolidating {
		intervals, consolidating = consolidate(intervals)
	}
	p2 := uint64(0)
	for _, ival := range intervals {
		p2 += ival.length()
	}
	println("p2:", p2)
}

func consolidate(intervals map[int]Interval) (map[int]Interval, bool) {
	replaced := make(map[int]bool)
	ret := make(map[int]Interval)
	index := 0

	for n, i1 := range intervals {
		for m, i2 := range intervals {
			if m <= n || replaced[n] || replaced[m] {
				continue
			}
			if min(i1.end, i2.end) >= max(i1.start, i2.start) {
				ret[index] = Interval{min(i1.start, i2.start), max(i1.end, i2.end)}
				index++
				replaced[n], replaced[m] = true, true
			}
		}
	}
	for n, ival := range intervals {
		if !replaced[n] {
			ret[index] = ival
			index++
		}
	}
	return ret, len(replaced) > 0
}
